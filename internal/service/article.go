package service

import (
	"github.com/WeCanRun/gin-blog/internal/dto"
	"github.com/WeCanRun/gin-blog/internal/model"
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/internal/service/cache_service"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/share"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"github.com/boombuler/barcode/qr"
	"github.com/jinzhu/gorm"
)

const QRCODE_URL = "https://github.com/WeCanRun/gin-blog%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"

func GetArticles(ctx *server.Context, req *dto.GetArticlesRequest) (resp dto.GetArticlesResponse, err error) {
	pageNum := util.GetPage(ctx)
	pageSize := util.GetPageSize(ctx)
	articles, err := model.GetArticles(ctx.Request.Context(), pageNum, pageSize)
	if err != nil {
		return
	}
	for _, article := range articles {
		resp.IDs = append(resp.IDs, article.ID)
		resp.Titles = append(resp.Titles, article.Title)
		resp.TagIds = append(resp.TagIds, article.TagId)
	}
	tags, err := model.GetTagsByIds(ctx.Request.Context(), resp.TagIds)
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

	resp.Pager = dto.Pager{
		PageNum:   pageNum,
		PageSize:  pageSize,
		TotalRows: uint(len(articles)),
	}
	return
}

func GetArticle(ctx *server.Context, id uint) (model.Article, error) {
	article, _ := cache_service.GetArticleCacheById(id)
	if article.ID > 0 {
		return article, nil
	}
	article, err := model.GetArticleById(ctx.Request.Context(), id)
	if err != nil {
		logging.Error("GetArticle | model.GetArticleById fail, err:%v", err)
		return article, err
	}
	err = cache_service.SetArticleCacheById(article)
	if err != nil {
		logging.Error("cache_service.SetArticleCacheById fail, err: %v", err)
	}

	return article, nil
}

func AddArticle(ctx *server.Context, req *dto.AddArticleRequest) error {
	// todo 各种逻辑校验
	return model.AddArticle(ctx.Request.Context(), model.Article{
		TagId:         req.TagId,
		Title:         req.Title,
		Desc:          req.Desc,
		Content:       req.Content,
		CoverImageUrl: req.CoverImageUrl,
		CreatedBy:     req.CreatedBy,
		State:         req.State,
	})
}

func EditArticle(ctx *server.Context, req *dto.EditArticleRequest) error {
	// todo 各种逻辑校验
	updateArticle := model.Article{
		Model: gorm.Model{
			ID: req.ID,
		},
		TagId:         req.TagId,
		Title:         req.Title,
		Desc:          req.Desc,
		Content:       req.Content,
		CoverImageUrl: req.CoverImageUrl,
		UpdatedBy:     req.UpdatedBy,
		State:         req.State,
	}

	err := model.EditArticle(ctx.Request.Context(), updateArticle)
	if err != nil {
		return err
	}

	article, err := model.GetArticleById(ctx.Request.Context(), req.ID)
	if err != nil {
		logging.Error("EditArticle |  model.GetArticleById fail, err%v", err)
	}

	err = cache_service.SetArticleCacheById(article)
	if err != nil {
		logging.Error("cache_service.SetArticleCacheById fail, err%v", err)
	}

	return err
}

func DeleteArticle(ctx *server.Context, id uint) error {
	// todo 各种逻辑校验
	err := model.DeleteArticle(ctx.Request.Context(), id)
	if err != nil {
		logging.Error("DeleteArticle | model.DeleteArticle fail, err:%v", err)
		return err
	}

	err = cache_service.DeleteArticleCacheById(id)
	if err != nil {
		logging.Error("DeleteArticle | cache_service.DeleteArticleCacheById fail, err:%v", err)
	}

	return err
}

func GenPoster(req *dto.GenArticlePosterReq) (dto.GenArticlePosterResp, error) {
	resp := dto.GenArticlePosterResp{}
	qrc := share.NewQrCode(QRCODE_URL, req.Height, req.Width, qr.M, qr.Auto)
	article := model.Article{}
	posterName := share.GetPosterFlag() + "-" + share.GetQrCodeFileName(QRCODE_URL) + qrc.GetQrCodeExt()
	articlePoster := share.NewArticlePoster(posterName, &article, qrc)
	articlePosterBg := share.NewArticlePosterBg("bg.jpg", articlePoster,
		&share.Rect{
			X0: 0,
			Y0: 0,
			X1: 550,
			Y1: 750,
		}, &share.Pt{
			X: 125,
			Y: 298,
		})

	path, _, err := articlePosterBg.Generate()
	if err != nil {
		logging.Error("GenPoster | articlePosterBg.Generate fail, err:%v", err)
		return resp, err
	}
	logging.Info("GenPoster | src:%v", path+posterName)

	resp.PosterSaveUrl = share.GetQrCodeSavePath(posterName)
	resp.PosterUrl = share.GetQrCodeFullUrl(posterName)

	return resp, nil
}
