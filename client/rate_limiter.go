package client

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/time/rate"
)

type RateLimitManager interface {
	Register(key string, rate float64, burst int) error
	Wait(ctx context.Context, key string) error
}

type RateLimiter struct {
	Limiter *rate.Limiter
}

func (r *RateLimiter) Wait(ctx context.Context) error {
	return r.Limiter.Wait(ctx)
}

type DefaultRateLimitManager struct {
	mu       sync.Mutex
	limiters map[string]*RateLimiter
}

func (m *DefaultRateLimitManager) Register(key string, rateVal float64, burst int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.limiters[key] = &RateLimiter{
		Limiter: rate.NewLimiter(rate.Limit(rateVal), burst),
	}
	return nil
}

func (m *DefaultRateLimitManager) Wait(ctx context.Context, key string) error {
	m.mu.Lock()
	limiter, ok := m.limiters[key]
	m.mu.Unlock()

	if !ok {
		return fmt.Errorf("rate limiter not registered for key: %s", key)
	}
	return limiter.Wait(ctx)
}

func NewRateLimitManager() *DefaultRateLimitManager {
	return &DefaultRateLimitManager{
		limiters: make(map[string]*RateLimiter),
	}
}
