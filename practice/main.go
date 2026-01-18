package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type PingResponse struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

// type Response struct {
// 	Code int         `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Data interface{} `json:"data"`
// }

type BizError struct {
	Code int
	Msg  string
}

func (e BizError) Error() string {
	return e.Msg
}

// func Success(ctx *app.RequestContext, data interface{}) {
// 	ctx.JSON(200, Response{
// 		Code: 0,
// 		Msg:  "ok",
// 		Data: data,
// 	})
// }

// func FailWithCode(ctx *app.RequestContext, code int, msg string) {
// 	ctx.JSON(400, Response{
// 		Code: code,
// 		Msg:  msg,
// 		Data: nil,
// 	})
// }

// func SquareService(nStr string) (int, error) {
// 	n, err := strconv.Atoi(nStr)
// 	if err != nil {
// 		return 0, BizError{Code: 10001, Msg: "invalid parameter: n"}
// 	}
// 	return n * n, nil
// }

// func AuthMiddleware() app.HandlerFunc {
// 	return func(c context.Context, ctx *app.RequestContext) {
// 		token := string(ctx.GetHeader("X-Token"))
// 		if token != "123" {
// 			FailWithCode(ctx, 10002, "unauthorized")
// 			return
// 		}

// 		// 放行，继续走后面的 handler
// 		ctx.Next(c)
// 	}
// }

// func RequestLogMiddleware() app.HandlerFunc {
// 	return func(c context.Context, ctx *app.RequestContext) {
// 		start := time.Now()

// 		// 继续执行后面的 handler
// 		ctx.Next(c)

// 		// handler 执行完之后再走到这里
// 		cost := time.Since(start)

// 		method := string(ctx.Method())
// 		path := string(ctx.Path())

// 		log.Printf("[REQ] %s %s cost=%v\n", method, path, cost)
// 	}
// }

// func BizErrorLogMiddleware() app.HandlerFunc {
// 	return func(c context.Context, ctx *app.RequestContext) {
// 		start := time.Now()

// 		// 先让后面的 handler 执行
// 		ctx.Next(c)

// 		cost := time.Since(start)

// 		// 读取响应 body（是 []byte）
// 		bodyBytes := ctx.Response.Body()
// 		if len(bodyBytes) == 0 {
// 			return
// 		}

// 		// 尝试解析成 Response
// 		var resp Response
// 		err := json.Unmarshal(bodyBytes, &resp)
// 		if err != nil {
// 			return
// 		}

// 		// 只要 code != 0，就认为是业务错误，打印日志
// 		if resp.Code != 0 {
// 			method := string(ctx.Method())
// 			path := string(ctx.Path())

// 			log.Printf("[BIZ_ERR] %s %s code=%d msg=%s cost=%v\n",
// 				method, path, resp.Code, resp.Msg, cost)
// 		}
// 	}
// }

func main() {
	h := server.Default(
		server.WithHostPorts("127.0.0.1:8888"),
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

	log.Println("Hertz server listening on http://127.0.0.1:8888")
	h.Spin() // ⚠️ 必须大写
}
