package web

import (
	"github.com/WeCanRun/gin-blog/internal/middleware"
	"github.com/WeCanRun/gin-blog/internal/server"
	v1 "github.com/WeCanRun/gin-blog/internal/web/v1"
	"github.com/WeCanRun/gin-blog/pkg/export"
	"github.com/WeCanRun/gin-blog/pkg/limiter"
	"github.com/WeCanRun/gin-blog/pkg/share"
	"github.com/WeCanRun/gin-blog/pkg/upload"
	"net/http"
	"time"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1
func InitRouters(router *server.RouterWarp) {

	// use middleware
	router.Use(middleware.Recovery())
	router.Use(middleware.AccessLog())
	router.Use(middleware.Translations())
	router.Use(middleware.Tracer())
	router.Use(middleware.Limiter(limiter.NewMethodLimiter().AddBuckets(limiter.BucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     100,
		Quantum:      10,
	})))

	router.Use(middleware.TimeOut())

	router.GET("/ping", v1.Ping)

	router.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	router.StaticFS("/export", http.Dir(export.GetExcelRealDir()))
	router.StaticFS("/qr_code", http.Dir(share.GetQrCodeSaveDir()))

	// auth
	router.GET("/auth", v1.GetToken)

	// upload
	router.POST("/image", v1.UploadImage)

	//router.Use(middleware.JWT())
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
}
