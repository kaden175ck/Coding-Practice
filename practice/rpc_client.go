package main

import (
	"log"
	"os"

	"order-rpc/kitex_gen/order/orderservice"

	"github.com/cloudwego/kitex/client"
)

var orderClient orderservice.Client

func initOrderRPCClient() {
	addr := os.Getenv("ORDER_RPC_ADDR")
	if addr == "" {
		addr = "127.0.0.1:10004" // order-rpc 实际监听的地址
	}

	c, err := orderservice.NewClient(
		"order-rpc",
		client.WithHostPorts(addr),
	)
	if err != nil {
		log.Fatalf("init order rpc client failed: %v", err)
	}

	orderClient = c
}
