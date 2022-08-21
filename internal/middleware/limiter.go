package middleware

import (
	"github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func Limiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		if bucket, ok := l.GetBucket(l.Key(c)); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				res := errcode.TooManyRequests
				c.AbortWithStatusJSON(res.StatusCode(), res)
				return
			}
		}
		c.Next()
	}
}
