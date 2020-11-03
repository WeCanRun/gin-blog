package services

import (
	"errors"
	"github.com/WeCanRun/gin-blog/dto"
	"github.com/WeCanRun/gin-blog/models"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/util"
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

	tags, err := models.GetTags(pageNum, pageSize)
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
	tag, err := models.GetTagById(id)
	if err != nil {
		logging.Error("models#GetTagById fail, %v", err)
		return
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
	err = models.AddTag(models.Tag{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		State:     req.State,
	})
	return
}

func DeleteTag(ctx *gin.Context, id int) (err error) {
	// todo 各种校验
	return models.DeleteTag(id)
}

func EditTag(ctx *gin.Context, tag *dto.EditRequest) (err error) {
	// todo 各种校验
	return models.EditTag(models.Tag{
		Model:     gorm.Model{ID: tag.ID},
		Name:      tag.Name,
		UpdatedBy: tag.UpdatedBy,
		State:     tag.State,
	})
}
