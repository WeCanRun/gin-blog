package main

import (
	"context"
	"fmt"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := routers.InitRouters()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeOut,
		WriteTimeout:   setting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
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

	if err := s.Shutdown(timeout); err != nil {
		logging.Error("server shutdown err", err)
	}
	logging.Info("Server exiting")
}
