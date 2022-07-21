package middleware

import (
	e "github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"

	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//token := ctx.Query("token")
		token := ctx.GetHeader("token")
		if token == "" || len(token) <= 0 {
			logging.Info("ctx.GetHeader | token 不存在 ")
			e.AuthError(ctx, e.ERROR_AUTH_CHECK_TOKEN_FAIL)
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil || claims == nil {
			logging.Info("JWT#ParseToken fail，%v", err)
			e.AuthError(ctx, e.ERROR_AUTH_CHECK_TOKEN_FAIL)
			return
		} else if claims.ExpiresAt < time.Now().Unix() {
			e.AuthError(ctx, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			return
		}
		ctx.Next()
	}
}
