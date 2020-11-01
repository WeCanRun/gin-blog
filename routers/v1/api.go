package v1

import (
	"github.com/WeCanRun/gin-blog/dto"
	"github.com/WeCanRun/gin-blog/pkg/e"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

// test
func Ping(ctx *gin.Context) {
	logging.Info(ctx.ClientIP())
	e.Success(ctx, "pong")
}

// 获取标签列表
func GetTags(ctx *gin.Context) {
	req := new(dto.GetTagsRequest)
	if err := ctx.Bind(&req); err != nil {
		logging.Error("bind param err, %v", err)
		e.ParamsError(ctx)
		return
	}
	data, err := services.GetTags(ctx, req)
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
	tag, err := services.GetTag(ctx, uint(id))
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
	var err = services.AddTag(ctx, req)
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
	if err = services.DeleteTag(ctx, id); err != nil {
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
	if err := services.EditTag(ctx, req); err != nil {
		logging.Error("services#EditTag fail,%v", err)
		e.ServerError(ctx, "编辑标签失败")
		return
	}
	e.Success(ctx, "编辑标签成功")
}

// 获取文章列表
func GetArticles(ctx *gin.Context) {
	req := new(dto.GetArticlesRequest)
	if err := ctx.Bind(req); err != nil {
		logging.Error("bind param err: %v", err)
		e.ParamsError(ctx)
		return
	}
	resp, err := services.GetArticles(ctx, req)
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
	article, err := services.GetArticle(ctx, uint(id))
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
	if err := services.AddArticle(ctx, req); err != nil {
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
	if err := services.EditArticle(ctx, req); err != nil {
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
	if err := services.DeleteArticle(ctx, uint(id)); err != nil {
		logging.Error("services#DeleteArticle fail,%v", err)
		e.ServerError(ctx, "删除文章失败")
	}
	e.Success(ctx, "删除文章成功")
}

func GetToken(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	if len(username) > 50 || len(username) <= 0 || len(password) > 50 || len(password) < 6 {
		e.ParamsError(ctx)
		logging.Error("GetToken params err")
		return
	}
	code, token := services.GetTokenWithAuth(username, password)
	if len(token) <= 0 {
		e.OtherError(ctx, code)
		return
	}
	e.Success(ctx, map[string]string{"token": token})
}
