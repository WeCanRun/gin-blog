package web

import (
	"github.com/WeCanRun/gin-blog/internal/middleware"
	v1 "github.com/WeCanRun/gin-blog/internal/web/v1"
	"github.com/WeCanRun/gin-blog/pkg/export"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/share"
	"github.com/WeCanRun/gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	gin.SetMode(setting.Server.RunMode)

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", v1.Ping)

	router.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	router.StaticFS("/export", http.Dir(export.GetExcelRealDir()))
	router.StaticFS("/qr_code", http.Dir(share.GetQrCodeSaveDir()))

	// auth
	router.GET("/auth", v1.GetToken)

	// upload
	router.POST("/image", v1.UploadImage)

	router.Use(middleware.JWT())
	apiV1 := router.Group("/api/v1")
	{
		// tag
		apiV1.GET("/tags", v1.GetTags)
		apiV1.GET("/tag/:id", v1.GetTag)
		apiV1.POST("/tag", v1.AddTag)
		apiV1.PUT("/tag", v1.EditTag)
		apiV1.DELETE("/tag/:id", v1.DeleteTag)
		apiV1.POST("/tag/export", v1.ExportTags)
		apiV1.POST("/tag/import", v1.ImportTag)

		// article todo 文章导入导出功能
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/article/:id", v1.GetArticle)
		apiV1.POST("/article", v1.AddArticle)
		apiV1.PUT("/article", v1.EditArticle)
		apiV1.DELETE("/article/:id", v1.DeleteArticle)
		apiV1.POST("/article/share", v1.GenerateArticlePoster)

	}
	return router
}
