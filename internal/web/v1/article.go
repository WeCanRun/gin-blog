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

// 获取文章列表
func GetArticles(ctx *server.Context) error {
	req := new(dto.GetArticlesRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err: %v", err)
		return ctx.ParamsError()
	}
	resp, err := service.GetArticles(ctx, req)
	if err != nil {
		logging.Error("services#GetArticles fail,%v", err)
		return ctx.ServerError("获取文章列表失败")
	}

	ctx.Logger().Debug(resp)

	return ctx.Success(resp)
}

// 获取指定文章
func GetArticle(ctx *server.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if id <= 0 || err != nil {
		logging.Error("bind param err, %v", err)
		return ctx.ParamsError()
	}
	article, err := service.GetArticle(ctx, uint(id))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return e.ErrorNotExistArticle
		}
		logging.Error("services#GetArticle fail,%v", err)
		return ctx.ServerError("获取文章失败")
	}
	return ctx.Success(article)
}

// 新增一篇文章
func AddArticle(ctx *server.Context) error {
	req := new(dto.AddArticleRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err, %v", err)
		return ctx.ParamsError()
	}
	if err := service.AddArticle(ctx, req); err != nil {
		logging.Error("services#AddArticle fail,%v", err)
		return ctx.ServerError("新增文章失败")
	}
	return ctx.Success("新增文章成功")
}

// 编辑文章
func EditArticle(ctx *server.Context) error {
	req := new(dto.EditArticleRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err, %v", err)
		return ctx.ParamsError()
	}
	if err := service.EditArticle(ctx, req); err != nil {
		logging.Error("services#EditArticle fail,%v", err)
		return ctx.ServerError("编辑文章失败")
	}
	return ctx.Success("编辑文章成功")
}

// 删除文章
func DeleteArticle(ctx *server.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		logging.Error("bind param err, %v", err)
		return ctx.ParamsError()
	}
	if err := service.DeleteArticle(ctx, uint(id)); err != nil {
		logging.Error("services#DeleteArticle fail,%v", err)
		return ctx.ServerError("删除文章失败")
	}
	return ctx.Success("删除文章成功")
}

// 生成海报
func GenerateArticlePoster(ctx *server.Context) error {
	req := new(dto.GenArticlePosterReq)
	if err := ctx.Bind(req); err != nil {
		logging.Error("GenerateArticlePoster | bind params fail, err:%v", err)
		return ctx.ParamsError()
	}

	if req.Width < 33 || req.Height < 33 {
		logging.Error("GenerateArticlePoster | 参数错误，req:%v", req)
		return ctx.ParamsError()
	}
	data, err := service.GenPoster(req)
	if err != nil {
		logging.Error("GenerateArticlePoster#service.GenPoster() | 生成海报失败,err:%v", err)
		return ctx.ServerError(data)
	}

	return ctx.Success(data)
}
