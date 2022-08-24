package v1

import (
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/internal/service"
	"github.com/WeCanRun/gin-blog/pkg/logging"
)

// test
func Ping(ctx *server.Context) error {
	logging.Infof("from: %s", ctx.Request.RemoteAddr)
	panic("PONG")
	return ctx.Success("pong")
}

// 获取凭证
func GetToken(ctx *server.Context) error {
	username := ctx.Query("username")
	password := ctx.Query("password")
	if len(username) > 50 || len(username) <= 0 || len(password) > 50 || len(password) < 6 {
		logging.Error("GetToken params err")
		return ctx.ParamsError("")
	}
	err, token := service.GetTokenWithAuth(ctx, username, password)
	if len(token) <= 0 {
		return err
	}
	return ctx.Success(map[string]string{"token": token})
}

// 上传图片
func UploadImage(ctx *server.Context) error {
	file, image, err := ctx.Request.FormFile("image")
	if err != nil || image == nil {
		logging.Error("param err: %v, image:%v", err, image)
		return ctx.ParamsError(err)
	}

	return service.UploadImage(ctx, file, image)
}
