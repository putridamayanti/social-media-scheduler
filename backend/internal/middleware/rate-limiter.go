package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

type ClientLimiter struct {
	Limiter  *rate.Limiter
	LastSeen time.Time
}

type RateLimiter struct {
	mu      sync.Mutex
	clients map[string]*ClientLimiter
	rate    rate.Limit
	burst   int
	ttl     time.Duration
}

func NewRateLimiter(r rate.Limit, burst int, ttl time.Duration) *RateLimiter {
	return &RateLimiter{
		clients: make(map[string]*ClientLimiter),
		rate:    r,
		burst:   burst,
		ttl:     ttl,
	}
}

func (rateLimiter *RateLimiter) GetLimiter(key string) *rate.Limiter {
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()

	client, exists := rateLimiter.clients[key]
	if exists {
		client.LastSeen = time.Now()
		return client.Limiter
	}

	limiter := rate.NewLimiter(rate.Limit(rateLimiter.rate), rateLimiter.burst)
	rateLimiter.clients[key] = &ClientLimiter{
		Limiter:  limiter,
		LastSeen: time.Now(),
	}

	return limiter
}

func (rateLimiter *RateLimiter) Cleanup() {
	for {
		time.Sleep(time.Minute)
		rateLimiter.mu.Lock()
		for key, client := range rateLimiter.clients {
			if time.Since(client.LastSeen) > rateLimiter.ttl {
				delete(rateLimiter.clients, key)
			}
		}
		rateLimiter.mu.Unlock()
	}
}

func (rateLimiter *RateLimiter) Middleware() gin.HandlerFunc {
	go rateLimiter.Cleanup()

	return func(c *gin.Context) {
		key := c.GetString("user_id")
		if key == "" {
			key = c.ClientIP()
		}

		limiter := rateLimiter.GetLimiter(key)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}
