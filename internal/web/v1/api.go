package v1

import (
	e "github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/internal/service"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/gin-gonic/gin"
)

// test
func Ping(ctx *gin.Context) {
	logging.Info(ctx.ClientIP())
	e.Success(ctx, "pong")
}

// 获取凭证
func GetToken(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	if len(username) > 50 || len(username) <= 0 || len(password) > 50 || len(password) < 6 {
		e.ParamsError(ctx)
		logging.Error("GetToken params err")
		return
	}
	code, token := service.GetTokenWithAuth(username, password)
	if len(token) <= 0 {
		e.OtherError(ctx, code)
		return
	}
	e.Success(ctx, map[string]string{"token": token})
}

// 上传图片
func UploadImage(ctx *gin.Context) {
	file, image, err := ctx.Request.FormFile("image")
	if err != nil || image == nil {
		logging.Error("param err: %v, image:%v", err, image)
		e.ParamsError(ctx)
		return
	}

	data, code, err := service.UploadImage(ctx, file, image)
	if err != nil || code != 0 {
		e.OtherError(ctx, code)
		return
	}
	e.Success(ctx, data)
}
