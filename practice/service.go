package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func SquareService(nStr string) (int, error) {
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return 0, BizError{Code: 10001, Msg: "invalid parameter: n"}
	}
	return n * n, nil
}

const DemoToken = "abc123"

type SessionStore struct {
	mu    sync.Mutex
	items map[string]time.Time // token -> expireTime
	ttl   time.Duration
}

func NewSessionStore(ttl time.Duration) *SessionStore {
	return &SessionStore{
		items: make(map[string]time.Time),
		ttl:   ttl,
	}
}

func (s *SessionStore) NewToken() string {
	// 生成 16 字节随机数 => 32 位 hex 字符串
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// demo 里简化处理：极少发生
		return ""
	}

	token := hex.EncodeToString(b)

	s.mu.Lock()
	s.items[token] = time.Now().Add(s.ttl)
	s.mu.Unlock()

	return token
}

func (s *SessionStore) Validate(token string) bool {
	now := time.Now()

	s.mu.Lock()
	expireTime, ok := s.items[token]
	if !ok {
		s.mu.Unlock()
		return false
	}

	if now.After(expireTime) {
		// 过期就删掉
		delete(s.items, token)
		s.mu.Unlock()
		return false
	}

	s.mu.Unlock()
	return true
}

func getOrderServiceBaseURL() string {
	base := os.Getenv("ORDER_SERVICE_URL")
	if base == "" {
		return "http://127.0.0.1:10001"
	}
	return base
}

type OrderItem struct {
	OrderId string `json:"orderId"`
	Amount  int    `json:"amount"`
}

type OrdersResponse struct {
	UserId string      `json:"userId"`
	Orders []OrderItem `json:"orders"`
}

func FetchOrdersFromOrderService(ctx context.Context, userId string) (OrdersResponse, error) {
	baseURL := getOrderServiceBaseURL()
	url := fmt.Sprintf("%s/orders/%s", baseURL, userId)

	// 1) 给 HTTP 请求加超时（微服务必备）
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	// 2) 创建请求（把 ctx 传进去，以后可以支持取消/超时）
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return OrdersResponse{}, err
	}

	// 3) 发请求
	resp, err := client.Do(req)
	if err != nil {
		return OrdersResponse{}, fmt.Errorf("call order-service failed: %w", err)
	}
	defer resp.Body.Close()

	// 4) 非 200 当成错误（demo 简化）
	if resp.StatusCode != 200 {
		return OrdersResponse{}, fmt.Errorf("order-service returned non-200: %d", resp.StatusCode)
	}

	// 5) 解析 JSON
	var data OrdersResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return OrdersResponse{}, err
	}

	return data, nil
}

func FetchOrdersWithRetry(ctx context.Context, userId string) (OrdersResponse, error) {
	var lastErr error

	// 最多 3 次：第 1 次 + 重试 2 次
	for attempt := 1; attempt <= 3; attempt++ {
		data, err := FetchOrdersFromOrderService(ctx, userId)
		if err == nil {
			// 成功了
			if attempt > 1 {
				// 如果重试后才成功，记录一下
				fmt.Printf("[RETRY] succeeded after %d attempts, userId=%s\n", attempt, userId)
			}
			return data, nil
		}

		lastErr = err
		fmt.Printf("[RETRY] attempt=%d failed, userId=%s, err=%v\n", attempt, userId, err)

		// 如果已经是最后一次，就不 sleep 了
		if attempt == 3 {
			fmt.Printf("[RETRY] all attempts failed, userId=%s\n", userId)
			break
		}

		// 退避：100ms、200ms
		var sleepDuration time.Duration
		if attempt == 1 {
			sleepDuration = 100 * time.Millisecond
		} else if attempt == 2 {
			sleepDuration = 200 * time.Millisecond
		}
		fmt.Printf("[RETRY] sleeping %v before next attempt\n", sleepDuration)
		time.Sleep(sleepDuration)
	}

	return OrdersResponse{}, lastErr
}
