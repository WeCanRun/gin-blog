package v1

import (
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/internal/service"
	"github.com/WeCanRun/gin-blog/pkg/logging"
)

// test
func Ping(ctx *server.Context) error {
	logging.Info(ctx.ClientIP())
	return ctx.Success("pong")
}

// 获取凭证
func GetToken(ctx *server.Context) error {
	username := ctx.Query("username")
	password := ctx.Query("password")
	if len(username) > 50 || len(username) <= 0 || len(password) > 50 || len(password) < 6 {
		logging.Error("GetToken params err")
		return ctx.ParamsError()
	}
	code, token := service.GetTokenWithAuth(username, password)
	if len(token) <= 0 {
		return ctx.OtherError(code)
	}
	return ctx.Success(map[string]string{"token": token})
}

// 上传图片
func UploadImage(ctx *server.Context) error {
	file, image, err := ctx.Request.FormFile("image")
	if err != nil || image == nil {
		logging.Error("param err: %v, image:%v", err, image)
		return ctx.ParamsError()
	}

	return service.UploadImage(ctx, file, image)
}
