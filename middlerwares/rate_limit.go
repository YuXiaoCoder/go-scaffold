package middlerwares

import (
	"go-scaffold/controllers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// RateLimit 限速中间件（令牌桶）
// Param fillInterval 填充令牌的时间间隔
// Param capacity 令牌桶的容量
func RateLimit(fillInterval time.Duration, capacity int64) func(ctx *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	return func(ctx *gin.Context) {
		// 若获取不到令牌，则返回响应
		if bucket.TakeAvailable(1) == 0 {
			controllers.ResponseError(ctx, controllers.CodeRateLimit)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
