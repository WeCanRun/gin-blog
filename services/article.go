package services

import (
	"github.com/WeCanRun/gin-blog/dto"
	"github.com/WeCanRun/gin-blog/models"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetArticles(ctx *gin.Context, req *dto.GetArticlesRequest) (resp dto.GetArticlesResponse, err error) {
	pageNum, pageSize := req.PageNum, req.PageSize
	if pageNum <= 0 {
		pageNum = util.GetPage(ctx)
	}
	if pageSize <= 0 {
		pageSize = setting.App.PageSize
	}
	articles, err := models.GetArticles(pageNum, pageSize)
	if err != nil {
		return
	}
	for _, article := range articles {
		resp.IDs = append(resp.IDs, article.ID)
		resp.Titles = append(resp.Titles, article.Title)
		resp.TagIds = append(resp.TagIds, article.TagId)
	}
	tags, err := models.GetTagsByIds(resp.TagIds)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.TagIds = nil
		}
		logging.Error("models#GetTagsByIds fail, %v", err)
		return
	}
	for _, tag := range tags {
		resp.TagNames = append(resp.TagNames, tag.Name)
	}
	return
}

func GetArticle(ctx *gin.Context, id uint) (article models.Article, err error) {
	return models.GetArticleById(id)
}

func AddArticle(ctx *gin.Context, req *dto.AddArticleRequest) (err error) {
	// todo 各种逻辑校验
	return models.AddArticle(models.Article{
		TagId:     req.TagId,
		Title:     req.Title,
		Desc:      req.Desc,
		Content:   req.Desc,
		CreatedBy: req.CreatedBy,
		State:     req.State,
	})
}

func EditArticle(ctx *gin.Context, req *dto.EditArticleRequest) (err error) {
	// todo 各种逻辑校验
	return models.EditArticle(models.Article{
		Model: gorm.Model{
			ID: req.ID,
		},
		TagId:     req.TagId,
		Title:     req.Title,
		Desc:      req.Desc,
		Content:   req.Content,
		UpdatedBy: req.UpdatedBy,
		State:     req.State,
	})
}

func DeleteArticle(ctx *gin.Context, id uint) (err error) {
	// todo 各种逻辑校验
	return models.DeleteArticle(id)
}
