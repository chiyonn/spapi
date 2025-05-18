package client

import "sync"

type RateLimiter struct {
	Rate  float64
	Burst int
}

type RateLimitManager interface {
	Register(key string, rate float64, burst int) error
}

type DefaultRateLimitManager struct {
	mu       sync.Mutex
	limiters map[string]*RateLimiter
}

func (m *DefaultRateLimitManager) Register(key string, rate float64, burst int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.limiters[key] = &RateLimiter{
		Rate:  rate,
		Burst: burst,
	}
	return nil
}

func NewRateLimitManager() *DefaultRateLimitManager {
	return &DefaultRateLimitManager{
		limiters: make(map[string]*RateLimiter),
	}
}
