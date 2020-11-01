package routers

import (
	_ "github.com/WeCanRun/gin-blog/docs"
	"github.com/WeCanRun/gin-blog/middleware"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	v1 "github.com/WeCanRun/gin-blog/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	gin.SetMode(setting.RunMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", v1.Ping)
	// auth
	router.GET("/auth", v1.GetToken)
	router.Use(middleware.JWT())

	apiV1 := router.Group("/api/v1")
	{
		// tag
		apiV1.GET("/tags", v1.GetTags)
		apiV1.GET("/tag/:id", v1.GetTag)
		apiV1.POST("/tag", v1.AddTag)
		apiV1.PUT("/tag", v1.EditTag)
		apiV1.DELETE("/tag/:id", v1.DeleteTag)

		// article
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/article/:id", v1.GetArticle)
		apiV1.POST("/article", v1.AddArticle)
		apiV1.PUT("/article", v1.EditArticle)
		apiV1.DELETE("/article/:id", v1.DeleteArticle)

	}
	return router
}
