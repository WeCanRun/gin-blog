package v1

import (
	"github.com/WeCanRun/gin-blog/dto"
	"github.com/WeCanRun/gin-blog/pkg/e"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

// 获取标签列表
func GetTags(ctx *gin.Context) {
	req := new(dto.GetTagsRequest)
	if err := ctx.Bind(&req); err != nil {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	data, err := service.GetTags(ctx, req)
	if err != nil {
		logging.Error("services#GetTags fail,%v", err)
		e.ServerError(ctx, "获取标签列表失败")
		return
	}

	e.Success(ctx, data)
}

// 获取指定标签
func GetTag(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if id <= 0 || err != nil {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	tag, err := service.GetTag(ctx, uint(id))
	if err != nil {
		logging.Error("services#GetTag err, %v", err)
		if gorm.IsRecordNotFoundError(err) {
			e.OtherError(ctx, e.ERROR_NOT_EXIST_TAG)
			return
		}
		e.ServerError(ctx, "获取标签失败")
		return
	}
	e.Success(ctx, tag)
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tag [post]cd cd
func AddTag(ctx *gin.Context) {
	req := new(dto.AddTagRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	var err = service.AddTag(ctx, req)
	if err != nil {
		logging.Error("services#AddTags fail,%v", err)
		e.ServerError(ctx, "增加标签失败")
		return
	}

	e.Success(ctx, nil)
}

// 删除指定标签
func DeleteTag(ctx *gin.Context) {
	idStr := ctx.Param("id")
	var id int
	var err error
	if id, err = strconv.Atoi(idStr); id <= 0 || err != nil {
		logging.Error("param err, id: %v, err: %v", idStr, err)
		e.ParamsError(ctx)
		return
	}
	if err = service.DeleteTag(ctx, uint(id)); err != nil {
		logging.Error("services#DeleteTag fail,%v", err)
		e.ServerError(ctx, "删除标签失败")
		return
	}

	e.Success(ctx, "删除标签成功")
}

// 编辑标签
func EditTag(ctx *gin.Context) {
	req := new(dto.EditRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	if err := service.EditTag(ctx, req); err != nil {
		logging.Error("services#EditTag fail,%v", err)
		e.ServerError(ctx, "编辑标签失败")
		return
	}
	e.Success(ctx, "编辑标签成功")
}

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
