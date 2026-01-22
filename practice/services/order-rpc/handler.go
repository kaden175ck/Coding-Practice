package main

import (
	"context"
	order "order-rpc/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// GetOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrders(ctx context.Context, req *order.GetOrdersRequest) (resp *order.GetOrdersResponse, err error) {
	// 模拟返回订单数据（和之前的 HTTP 版本一样）
	resp = &order.GetOrdersResponse{
		UserId: req.UserId,
		Orders: []*order.OrderItem{
			{
				OrderId: "o_1001",
				Amount:  99,
			},
			{
				OrderId: "o_1002",
				Amount:  199,
			},
		},
	}
	return resp, nil
}
