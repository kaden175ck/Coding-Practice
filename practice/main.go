package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type PingResponse struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

type BizError struct {
	Code int
	Msg  string
}

func (e BizError) Error() string {
	return e.Msg
}

var store = NewSessionStore(5 * time.Minute)

func getServerAddr() string {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		// 默认值（本地开发用）
		return "127.0.0.1:8888"
	}
	return addr
}

func main() {
	// h := server.Default(
	// 	server.WithHostPorts("127.0.0.1:8888"),
	// )
	addr := getServerAddr()

	h := server.Default(
		server.WithHostPorts(addr),
	)

	h.Use(RequestLogMiddleware())
	h.Use(BizErrorLogMiddleware())

	// 1) 最简单路由handler：GET /ping
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		resp := PingResponse{
			Message: "pong",
			Time:    time.Now().Format(time.RFC3339),
		}
		// ⚠️ 一定要用 ctx
		ctx.JSON(200, resp)
	})

	// 2) 带路径参数：GET /hello/:name
	h.GET("/hello/:name", func(c context.Context, ctx *app.RequestContext) {
		name := ctx.Param("name")
		msg := fmt.Sprintf("hello, %s", name)
		ctx.String(200, msg)
	})

	// 3) POST /sum 练 JSON 绑定
	h.POST("/sum", func(c context.Context, ctx *app.RequestContext) {
		var req struct {
			A int `json:"a"`
			B int `json:"b"`
		}

		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(400, map[string]any{
				"error": "invalid json body",
			})
			return
		}

		ctx.JSON(200, map[string]any{
			"a":   req.A,
			"b":   req.B,
			"sum": req.A + req.B,
		})
	})

	h.GET("/square/:n", func(c context.Context, ctx *app.RequestContext) {
		nStr := ctx.Param("n")

		result, err := SquareService(nStr)
		if err != nil {
			if bizErr, ok := err.(BizError); ok {
				FailWithCode(ctx, bizErr.Code, bizErr.Msg)
				return
			}
			FailWithCode(ctx, 10000, "internal error")
			return
		}

		Success(ctx, map[string]any{
			"square": result,
		})
	})

	h.GET("/add", func(c context.Context, ctx *app.RequestContext) {
		aStr := string(ctx.Query("a"))
		bStr := string(ctx.Query("b"))

		a, err := strconv.Atoi(aStr)
		if err != nil {
			ctx.JSON(400, map[string]any{
				"error": "a must be an integer",
			})
			return
		}

		b, err := strconv.Atoi(bStr)
		if err != nil {
			ctx.JSON(400, map[string]any{
				"error": "b must be an integer",
			})
			return
		}

		ctx.JSON(200, map[string]any{
			"a":   a,
			"b":   b,
			"sum": a + b,
		})
	})

	authGroup := h.Group("/")

	authGroup.Use(AuthMiddleware())

	authGroup.GET("/secret", func(c context.Context, ctx *app.RequestContext) {
		Success(ctx, map[string]any{
			"message": "you are authorized",
		})
	})

	h.POST("/login", func(c context.Context, ctx *app.RequestContext) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		err := ctx.Bind(&req)
		if err != nil {
			FailWithCode(ctx, 10001, "invalid json body")
			return
		}

		if req.Username != "admin" || req.Password != "123" {
			FailWithCode(ctx, 10003, "invalid username or password")
			return
		}

		token := store.NewToken()
		if token == "" {
			FailWithCode(ctx, 10000, "internal error")
			return
		}

		Success(ctx, map[string]any{
			"token": token,
			"ttl":   "5m",
		})
	})

	h.GET("/me/orders/:userId", func(c context.Context, ctx *app.RequestContext) {
		userId := ctx.Param("userId")

		data, err := FetchOrdersFromOrderService(c, userId)
		if err != nil {
			FailWithCode(ctx, 20001, "failed to fetch orders")
			return
		}

		Success(ctx, map[string]any{
			"from":   "order-service",
			"userId": data.UserId,
			"orders": data.Orders,
		})
	})

	// log.Println("Hertz server listening on http://127.0.0.1:8888")
	log.Println("=================================")
	log.Println("Service starting...")
	log.Println("Listen addr:", addr)
	log.Println("=================================")
	h.Spin() // ⚠️ 必须大写
}
