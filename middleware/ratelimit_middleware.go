package middleware

import (
	"net/http"
	apperr "porter/pkg/err"
	"sync"

	"golang.org/x/time/rate"
)

const rateLimit = rate.Limit(1) // 1 request per second

type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

func (rl *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if limiter, exists := rl.limiters[ip]; exists {
		return limiter
	}

	limiter := rate.NewLimiter(rateLimit, 10)
	rl.limiters[ip] = limiter
	return limiter
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIp := r.RemoteAddr
		limiter := rl.GetLimiter(userIp)
		if !limiter.Allow() {
			err := apperr.NewTooManyRequestsError("Too Many Requests")
			http.Error(w, err.Error(), err.StatusCode())
			return
		}

		next.ServeHTTP(w, r)

	})
}
