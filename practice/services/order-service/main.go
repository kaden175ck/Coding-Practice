package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func getServerAddr() string {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		return "127.0.0.1:10001"
	}
	return addr
}

func main() {
	addr := getServerAddr()

	h := server.Default(
		server.WithHostPorts(addr),
	)

	// 健康检查
	h.GET("/health", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, map[string]any{
			"service": "order-service",
			"status":  "ok",
		})
	})

	// 模拟订单查询：GET /orders/:userId
	h.GET("/orders/:userId", func(c context.Context, ctx *app.RequestContext) {
		userId := ctx.Param("userId")

		// 这里先写死两条订单，后面我们再讲“数据怎么存”
		ctx.JSON(200, map[string]any{
			"userId": userId,
			"orders": []map[string]any{
				{
					"orderId": "o_1001",
					"amount":  99,
				},
				{
					"orderId": "o_1002",
					"amount":  199,
				},
			},
		})
	})

	log.Println("[order-service] starting...")
	log.Println("[order-service] listen:", addr)
	h.Spin()
}
