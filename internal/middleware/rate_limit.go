package middleware

import (
	"net/http"
	"tikuAdapter/internal/registry/manager"

	"github.com/gin-gonic/gin"
)

// GlobalAPIRateLimit 全局API限流
func GlobalAPIRateLimit(c *gin.Context) {
	limiter := manager.GetManager().GetIPLimiter()
	congig := manager.GetManager().GetConfig()
	if congig.Limit.Enable && !limiter.GetLimiter(c.RemoteIP()).Allow() {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}
