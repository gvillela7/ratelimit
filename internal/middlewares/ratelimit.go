package middlewares

import (
	"github.com/gvillela7/ratelimit/configs"
	"github.com/gvillela7/ratelimit/internal/data/response"
	redis2 "github.com/gvillela7/ratelimit/internal/database/redis"
	"net"
	"time"

	"net/http"
)

func Limit(c redis2.IRedis, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := configs.GetAPIConfig()
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		token := r.Header.Get("API_KEY")
		rateLimiter, _ := c.NewRateLimit(cfg.RateLimitRequest, time.Duration(cfg.RateLimitTimeSecond)*time.Second)

		if token != "" {
			result, _ := c.Allow(token, rateLimiter)
			if !result {
				rateLimiter, _ = c.NewRateLimit(cfg.RateLimitRequest, time.Duration(cfg.RateLimitTimeBlock)*time.Second)
				_, _ = c.Allow(token, rateLimiter)
				response.HttpResponse(
					w, http.StatusTooManyRequests,
					"To Many Request", map[string]interface{}{
						"message": "you have reached the maximum number of requests or actions allowed within a certain time frame.",
					})
				return
			}

		} else {
			result, _ := c.Allow(ip, rateLimiter)
			if !result {
				rateLimiter, _ = c.NewRateLimit(cfg.RateLimitRequest, time.Duration(cfg.RateLimitTimeBlock)*time.Second)
				_, _ = c.Allow(ip, rateLimiter)
				response.HttpResponse(
					w, http.StatusTooManyRequests,
					"To Many Request", map[string]interface{}{
						"message": "you have reached the maximum number of requests or actions allowed within a certain time frame.",
					})
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
