namespace go order

struct OrderItem {
  1: string order_id
  2: i32 amount
}

struct GetOrdersRequest {
  1: string user_id
}

struct GetOrdersResponse {
  1: string user_id
  2: list<OrderItem> orders
}

service OrderService {
  GetOrdersResponse GetOrders(1: GetOrdersRequest req)
}
