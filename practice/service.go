package main

import (
	"crypto/rand"
	"encoding/hex"
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
