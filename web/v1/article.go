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

// 获取文章列表
func GetArticles(ctx *gin.Context) {
	req := new(dto.GetArticlesRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err: %v", err)
		e.ParamsError(ctx)
		return
	}
	resp, err := service.GetArticles(ctx, req)
	if err != nil {
		logging.Error("services#GetArticles fail,%v", err)
		e.ServerError(ctx, "获取文章列表失败")
		return
	}
	e.Success(ctx, resp)
}

// 获取指定文章
func GetArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if id <= 0 || err != nil {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	article, err := service.GetArticle(ctx, uint(id))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			e.OtherError(ctx, e.ERROR_NOT_EXIST_ARTICLE)
			return
		}
		logging.Error("services#GetArticle fail,%v", err)
		e.ServerError(ctx, "获取文章失败")
		return
	}
	e.Success(ctx, article)
}

// 新增一篇文章
func AddArticle(ctx *gin.Context) {
	req := new(dto.AddArticleRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	if err := service.AddArticle(ctx, req); err != nil {
		logging.Error("services#AddArticle fail,%v", err)
		e.ServerError(ctx, "新增文章失败")
		return
	}
	e.Success(ctx, "新增文章成功")
}

// 编辑文章
func EditArticle(ctx *gin.Context) {
	req := new(dto.EditArticleRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	if err := service.EditArticle(ctx, req); err != nil {
		logging.Error("services#EditArticle fail,%v", err)
		e.ServerError(ctx, "编辑文章失败")
		return
	}
	e.Success(ctx, "编辑文章成功")
}

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	if err := service.DeleteArticle(ctx, uint(id)); err != nil {
		logging.Error("services#DeleteArticle fail,%v", err)
		e.ServerError(ctx, "删除文章失败")
	}
	e.Success(ctx, "删除文章成功")
}

// 生成海报
func GenerateArticlePoster(ctx *gin.Context) {
	req := new(dto.GenArticlePosterReq)
	if err := ctx.Bind(req); err != nil {
		logging.Error("GenerateArticlePoster | bind params fail, err:%v", err)
		e.ParamsError(ctx)
		return
	}

	if req.Width < 33 || req.Height < 33 {
		logging.Error("GenerateArticlePoster | 参数错误，req:%v", req)
		e.ParamsError(ctx)
		return
	}
	data, err := service.GenPoster(req)
	if err != nil {
		logging.Error("GenerateArticlePoster#service.GenPoster() | 生成海报失败,err:%v", err)
		e.ServerError(ctx, data)
		return
	}

	e.Success(ctx, data)
}
