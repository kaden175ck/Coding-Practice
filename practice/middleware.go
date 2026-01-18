package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func AuthMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		auth := string(ctx.GetHeader("Authorization"))

		expected := "Bearer " + DemoToken
		if auth != expected {
			FailWithCode(ctx, 10002, "unauthorized")
			// ✅ 关键：中止后续 handler 执行
			ctx.Abort()
			return
		}

		// 放行，继续走后面的 handler
		ctx.Next(c)
	}
}

func RequestLogMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()

		// 继续执行后面的 handler
		ctx.Next(c)

		// handler 执行完之后再走到这里
		cost := time.Since(start)

		method := string(ctx.Method())
		path := string(ctx.Path())

		log.Printf("[REQ] %s %s cost=%v\n", method, path, cost)
	}
}

func BizErrorLogMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()

		// 先让后面的 handler 执行
		ctx.Next(c)

		cost := time.Since(start)

		// 读取响应 body（是 []byte）
		bodyBytes := ctx.Response.Body()
		if len(bodyBytes) == 0 {
			return
		}

		// 尝试解析成 Response
		var resp Response
		err := json.Unmarshal(bodyBytes, &resp)
		if err != nil {
			return
		}

		// 只要 code != 0，就认为是业务错误，打印日志
		if resp.Code != 0 {
			method := string(ctx.Method())
			path := string(ctx.Path())

			log.Printf("[BIZ_ERR] %s %s code=%d msg=%s cost=%v\n",
				method, path, resp.Code, resp.Msg, cost)
		}
	}
}
