package middleware

import (
	e "github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"time"
)

func JWT() server.Handler {
	return func(ctx *server.Context) error {
		//token := ctx.Query("token")
		token := ctx.GetHeader("token")
		if token == "" || len(token) <= 0 {
			logging.Info("ctx.GetHeader | token 不存在 ")
			return e.ErrorAuthCheckTokenFail
		}
		claims, err := util.ParseToken(token)
		if err != nil || claims == nil {
			logging.Info("JWT#ParseToken fail，%v", err)
			return e.ErrorAuthCheckTokenFail
		} else if claims.ExpiresAt < time.Now().Unix() {
			return e.ErrorAuthCheckTokenTimeout
		}
		ctx.Next()
		return e.Success
	}
}
