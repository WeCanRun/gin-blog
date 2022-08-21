package middleware

import (
	"context"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/gin-gonic/gin"
)

func TimeOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), global.Setting.Server.ReadTimeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
