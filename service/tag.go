package service

import (
	"errors"
	"github.com/WeCanRun/gin-blog/dto"
	"github.com/WeCanRun/gin-blog/model"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"github.com/WeCanRun/gin-blog/service/cache_service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTags(c *gin.Context, req *dto.GetTagsRequest) (resp *dto.GetTagsResponse, err error) {
	var pageNum = req.PageNum
	var pageSize = req.PageSize
	if pageNum <= 0 {
		pageNum = util.GetPage(c)
	}
	if pageSize <= 0 {
		pageSize = setting.App.PageSize
	}

	tags, err := model.GetTags(pageNum, pageSize)
	if err != nil {
		logging.Error("services#GetTags fail %v", err)
		return
	}
	var ids []uint
	var names []string
	for _, tag := range tags {
		ids = append(ids, tag.ID)
		names = append(names, tag.Name)
	}
	resp = &dto.GetTagsResponse{
		IDs:   ids,
		Names: names,
	}
	return
}

func GetTag(ctx *gin.Context, id uint) (resp dto.GetTagResponse, err error) {
	tag, err := cache_service.GetTagCacheById(id)
	if tag.ID > 0 {
		resp.ID = id
		resp.Name = tag.Name
		logging.Info("GetTag | get tag from cache, id:%d", id)
		return
	}

	tag, err = model.GetTagById(id)
	if err != nil {
		logging.Error("models#GetTagById fail, %v", err)
		return
	}

	if err := cache_service.SetTagCacheById(tag); err != nil {
		logging.Error("GetTag | cache_service.SetTagCacheById fail, err:%v", err)
	}

	resp.ID = id
	resp.Name = tag.Name
	return
}

func AddTag(ctx *gin.Context, req *dto.AddTagRequest) (err error) {
	if req.State != 0 && req.State != 1 {
		err = errors.New("services#AddTag state 只能为 0 或 1")
		logging.Error("err: %v", err)
		return
	}
	// todo 其他敏感词校验
	err = model.AddTag(model.Tag{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		State:     req.State,
	})
	return
}

func DeleteTag(ctx *gin.Context, id uint) (err error) {
	// todo 各种校验
	err = model.DeleteTag(id)
	if err != nil {
		logging.Error("DeleteTag | model.DeleteTag fail, err:%v", err)
		return
	}

	err = cache_service.DeleteTagCacheById(id)
	if err != nil {
		logging.Error("DeleteTag | cache_service.DeleteTagCacheById fail, err:%v", err)

	}

	return nil
}

func EditTag(ctx *gin.Context, request *dto.EditRequest) (err error) {
	// todo 各种校验
	err = model.EditTag(model.Tag{
		Model:     gorm.Model{ID: request.ID},
		Name:      request.Name,
		UpdatedBy: request.UpdatedBy,
		State:     request.State,
	})
	if err != nil {
		logging.Error("EditTag | model.EditTag fail, err:%v", err)
		return
	}

	tag, err := model.GetTagById(request.ID)
	if err != nil {
		logging.Error("EditTag | model.GetTagById fail, err:%v", err)
	}

	err = cache_service.SetTagCacheById(tag)
	if err != nil {
		logging.Error("EditTag | cache_service.SetTagCacheById fail, err%v", err)
	}

	return nil
}
