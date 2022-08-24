package service

import (
	e "github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/internal/dto"
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/upload"
	"github.com/pkg/errors"
	"mime/multipart"
)

func UploadImage(ctx *server.Context, file multipart.File, image *multipart.FileHeader) *e.InternalError {
	data := dto.UploadImageResponse{}
	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	src := fullPath + "/" + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		err := errors.Errorf("ext or size of image is error, image:%v", image)
		logging.Error(err)
		return ctx.ParamsError(err.Error())
	}

	err := upload.CheckImage(fullPath)
	if err != nil {
		logging.Error("upload.CheckImage fail, err:%v", err)
		return ctx.ServerError(err)
	}

	if err := ctx.SaveUploadedFile(image, src); err != nil {
		logging.Error("ctx.SaveUploadedFile fail, err:%v", err)
		return ctx.ServerError(err)
	}

	data.ImageUrl = upload.GetImageFullUrl(imageName)
	data.ImageSaveUrl = savePath + imageName

	return ctx.Success(data)
}
