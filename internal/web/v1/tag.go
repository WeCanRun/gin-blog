package v1

import (
	e "github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/internal/dto"
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/internal/service"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/jinzhu/gorm"
	"strconv"
)

// 获取标签列表
func GetTags(ctx *server.Context) error {
	req := new(dto.GetTagsRequest)
	if err := ctx.Bind(&req); err != nil {
		logging.Error("bind param err, %v", err)
		return ctx.ParamsError()
	}
	data, err := service.GetTags(ctx, req)
	if err != nil {
		logging.Error("services#GetTags fail,%v", err)
		return ctx.ServerError("获取标签列表失败")
	}

	return ctx.Success(data)
}

// 获取指定标签
func GetTag(ctx *server.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if id <= 0 || err != nil {
		logging.Error("bind param err, %v", err)
		return ctx.ParamsError()
	}
	tag, err := service.GetTag(ctx, uint(id))
	if err != nil {
		logging.Error("services#GetTag err, %v", err)
		if gorm.IsRecordNotFoundError(err) {
			return e.ErrorNotExistTag
		}
		return ctx.ServerError("获取标签失败")
	}

	return ctx.Success(tag)
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tag [post]cd cd
func AddTag(ctx *server.Context) error {
	req := new(dto.AddTagRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err, %v", err)
		return ctx.ParamsError()
	}

	if err := service.AddTag(ctx, req); err != nil {
		logging.Error("services#AddTags fail,%v", err)
		return ctx.ServerError("增加标签失败")
	}

	return ctx.Success(nil)
}

// 删除指定标签
func DeleteTag(ctx *server.Context) error {
	idStr := ctx.Param("id")
	var id int
	var err error
	if id, err = strconv.Atoi(idStr); id <= 0 || err != nil {
		logging.Error("param err, id: %v, err: %v", idStr, err)
		return ctx.ParamsError()
	}

	if err = service.DeleteTag(ctx, uint(id)); err != nil {
		logging.Error("services#DeleteTag fail,%v", err)
		return ctx.ServerError("删除标签失败")
	}

	return ctx.Success("删除标签成功")
}

// 编辑标签
func EditTag(ctx *server.Context) error {
	req := new(dto.EditRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err, %v", err)
		return ctx.ParamsError()
	}
	if err := service.EditTag(ctx, req); err != nil {
		logging.Error("services#EditTag fail,%v", err)
		return ctx.ServerError("编辑标签失败")
	}
	return ctx.Success("编辑标签成功")
}

func ExportTags(ctx *server.Context) error {
	// todo reqDto
	req := new(dto.ExportTagsRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("ExportTags | params bind fail,err:%v", err)
		return ctx.ParamsError()
	}
	if req.Name == "" || len(req.Name) <= 0 {
		logging.Error("ExportTags | 参数错误, req:%v", req)
		return ctx.ParamsError()
	}
	resp, err := service.ExportTags(req.Name, req.State)
	if err != nil {
		logging.Error("ExportTags | ExportTags fail, err:%v", err)
		return ctx.ServerError("服务错误")
	}

	return ctx.Success(resp)
}

func ImportTag(ctx *server.Context) error {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		logging.Error("%v", err)
		return ctx.ParamsError()
	}

	if err := service.ImportTags(file); err != nil {
		logging.Error("ImportTag | servicreturn ctx.ImportTags fail, err:%v", err)
		return ctx.ServerError("导入标签失败")
	}

	return ctx.Success("导入标签成功")
}
