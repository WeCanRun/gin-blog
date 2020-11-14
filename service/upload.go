package service

import (
	"fmt"
	"github.com/WeCanRun/gin-blog/dto"
	"github.com/WeCanRun/gin-blog/pkg/e"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func UploadImage(ctx *gin.Context, file multipart.File, image *multipart.FileHeader) (dto.UploadImageResponse, int, error) {
	data := dto.UploadImageResponse{}
	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	src := fullPath + "/" + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		logging.Error("ext or size of image is error, image:%v", image)
		return data, e.INVALID_PARAMS, fmt.Errorf(e.GetMsg(e.INVALID_PARAMS))
	}

	err := upload.CheckImage(fullPath)
	if err != nil {
		logging.Error("upload.CheckImage fail, err:%v", err)
		return data, e.SERVER_ERROR, err
	}

	if err := ctx.SaveUploadedFile(image, src); err != nil {
		logging.Error("ctx.SaveUploadedFile fail, err:%v", err)
		return data, e.SERVER_ERROR, err
	}

	data.ImageUrl = upload.GetImageFullUrl(imageName)
	data.ImageSaveUrl = savePath + imageName

	return data, 0, nil
}
