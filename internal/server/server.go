package server

import (
	"context"
	"fmt"
	_ "github.com/WeCanRun/gin-blog/docs"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"net/http"
	"os"
	"os/signal"
	"time"
)

var svr *http.Server

func Init() *RouterWarp {
	router := NewRouter()

	svr = &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server.HttpPort),
		Handler:        router.Engine(),
		ReadTimeout:    setting.Server.ReadTimeout,
		WriteTimeout:   setting.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// use middleware

	// other

	return router
}

func Run(ctx context.Context) {
	// 启动服务
	go func() {
		if err := svr.ListenAndServe(); err != nil {
			logging.Error("Listen:%s", err)
		}
	}()

	quit := make(chan os.Signal)
	// 阻塞、等待终止信号
	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("shutdown server...")
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	if err := svr.Shutdown(timeout); err != nil {
		logging.Error("server shutdown err", err)
	}
	logging.Info("server exiting")
}

type RouterWarp struct {
	gh *gin.Engine
}

func NewRouter() *RouterWarp {
	router := gin.New()

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	return &RouterWarp{
		gh: router,
	}
}

func (w *RouterWarp) Engine() *gin.Engine {
	return w.gh
}

func (w *RouterWarp) POST(path string, handler Handler) {
	w.gh.POST(path, HandlerWarp(handler))
}

func (w *RouterWarp) GET(path string, handler Handler) {
	w.gh.GET(path, HandlerWarp(handler))
}

func (w *RouterWarp) Use(handlers ...Handler) {
	w.gh.Use(HandlerWarp(handlers...))
}

func (w *RouterWarp) Handle(method, path string, handler Handler) {
	w.gh.Handle(method, path, HandlerWarp(handler))
}

func (w *RouterWarp) Any(path string, handler Handler) {
	w.gh.Any(path, HandlerWarp(handler))
}

func (w *RouterWarp) DELETE(path string, handler Handler) {
	w.gh.DELETE(path, HandlerWarp(handler))
}

func (w *RouterWarp) PATCH(path string, handler Handler) {
	w.gh.PATCH(path, HandlerWarp(handler))
}

func (w *RouterWarp) PUT(path string, handler Handler) {
	w.gh.PUT(path, HandlerWarp(handler))
}

func (w *RouterWarp) OPTIONS(path string, handler Handler) {
	w.gh.OPTIONS(path, HandlerWarp(handler))
}

func (w *RouterWarp) HEAD(path string, handler Handler) {
	w.gh.HEAD(path, HandlerWarp(handler))
}

func (w *RouterWarp) StaticFile(path, filePath string) {
	w.gh.StaticFile(path, filePath)
}

func (w *RouterWarp) Static(path, filePath string) {
	w.gh.Static(path, filePath)
}

func (w *RouterWarp) StaticFS(path string, fs http.FileSystem) {
	w.gh.StaticFS(path, fs)
}

func (w *RouterWarp) Group(path string, handlers ...Handler) *RouterWarp {
	w.gh.RouterGroup = *w.gh.Group(path, HandlerWarp(handlers...))
	return &RouterWarp{
		gh: w.gh,
	}
}
