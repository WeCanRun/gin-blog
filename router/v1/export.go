package v1

import (
	"github.com/WeCanRun/gin-blog/dto"
	"github.com/WeCanRun/gin-blog/pkg/e"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/service"
	"github.com/gin-gonic/gin"
)

func ExportTags(ctx *gin.Context) {
	// todo reqDto
	req := new(dto.ExportTagsRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("ExportTags | params bind fail,err:%v", err)
		e.ParamsError(ctx)
		return
	}
	if req.Name == "" || len(req.Name) <= 0 {
		logging.Error("ExportTags | 参数错误, req:%v", req)
		e.ParamsError(ctx)
		return
	}
	resp, err := service.ExportTags(req.Name, req.State)
	if err != nil {
		logging.Error("ExportTags | ExportTags fail, err:%v", err)
		e.ServerError(ctx, "服务错误")
		return
	}

	e.Success(ctx, resp)
}

func ImportTag(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		logging.Error("%v", err)
		e.ParamsError(ctx)
		return
	}

	err = service.ImportTags(file)
	if err != nil {
		logging.Error("ImportTag | service.ImportTags fail, err:%v", err)
		e.ServerError(ctx, "导入标签失败")
		return
	}

	e.Success(ctx, "导入标签成功")
}
