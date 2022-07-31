package main

import (
	"context"
	"github.com/WeCanRun/gin-blog/internal/model"
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/internal/service/cache_service"
	"github.com/WeCanRun/gin-blog/internal/web"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
)

func main() {
	//加载配置文件
	setting.Setup("")
	// 加载日志配置
	logging.Setup()
	// 加载数据库
	model.Setup()
	// 加载 redis
	cache_service.Setup()

	routers := web.InitRouters()
	server.Init(routers)

	ctx := context.Background()
	server.Run(ctx)

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

}
