package registry

import (
	"tikuAdapter/configs"
	"tikuAdapter/pkg/ratelimit"

	"golang.org/x/time/rate"
)

// Limit get ratelimit instance
func Limit(cfg configs.Config) *ratelimit.IPRateLimiter {
	limit := cfg.Limit
	r := rate.Limit(float64(limit.Duration) / float64(limit.Requests))
	return ratelimit.NewIPRateLimiter(r, 1)
}
