package main

import (
	"context"
	"fmt"
	"github.com/WeCanRun/gin-blog/model"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/service/cache_service"
	"github.com/WeCanRun/gin-blog/web"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//加载配置文件
	setting.Setup("./conf/app.ini")
	// 加载日志配置
	logging.Setup()
	// 加载数据库
	model.Setup()
	// 加载 redis
	cache_service.Setup()

	router := web.InitRouters()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.Server.ReadTimeout,
		WriteTimeout:   setting.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动服务
	go func() {
		if err := s.ListenAndServe(); err != nil {
			logging.Error("Listen:%s", err)
		}
	}()

	// 定时任务
	//go func() {
	//	logging.Info("Starting...Cron Job")
	//	c := cron.New()
	//	c.AddFunc("1-59/10 * * * * *", func() {
	//		logging.Info("begin exec job1...")
	//	})
	//	c.AddFunc("1,11,41,51 * * * * *", func() {
	//		logging.Info("begin exec job2...")
	//	})
	//	c.Start()
	//	select {}
	//}()

	quit := make(chan os.Signal)
	// 阻塞、等待终止信号
	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("shutdown server...")
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	if err := s.Shutdown(timeout); err != nil {
		logging.Error("server shutdown err", err)
	}
	logging.Info("server exiting")
}
