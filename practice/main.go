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

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(ctx *app.RequestContext, data interface{}) {
	ctx.JSON(200, Response{
		Code: 0,
		Msg:  "ok",
		Data: data,
	})
}
func Fail(ctx *app.RequestContext, msg string) {
	ctx.JSON(400, Response{
		Code: -1,
		Msg:  msg,
		Data: nil,
	})
}

func main() {
	h := server.Default(
		server.WithHostPorts("127.0.0.1:8888"),
	)

	// 1) 最简单：GET /ping
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

		n, err := strconv.Atoi(nStr)
		if err != nil {
			Fail(ctx, "n must be an integer")
			return
		}

		Success(ctx, map[string]any{
			"n":      n,
			"square": n * n,
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

	log.Println("Hertz server listening on http://127.0.0.1:8888")
	h.Spin() // ⚠️ 必须大写
}
