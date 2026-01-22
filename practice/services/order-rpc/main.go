package main

import (
	"log"
	"net"
	order "order-rpc/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// 指定监听地址（用 10004 端口，避免和 HTTP 版本冲突）
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:10004")

	svr := order.NewServer(
		new(OrderServiceImpl),
		server.WithServiceAddr(addr),
	)

	log.Println("[order-rpc] starting on 127.0.0.1:10004...")
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
