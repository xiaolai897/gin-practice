package main

import (
	"context"
	"gin-practice/config"
	"gin-practice/initialize"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	config.SELF_VIPER = initialize.GetConfig()

	config.SELF_LOG = initialize.GetZap()
	zap.ReplaceGlobals(config.SELF_LOG)
	defer config.SELF_LOG.Sync()

	if config.SELF_DB = initialize.GetGorm(); config.SELF_DB != nil {
		db, _ := config.SELF_DB.DB()
		defer db.Close()
	}

	router := initialize.Routers()
	if router == nil {
		return
	}

	// 优雅关闭gin
	srv := &http.Server{
		Addr:    ":7900",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			config.SELF_LOG.Error("监听失败端口失败:")
			config.SELF_LOG.Error(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	config.SELF_LOG.Info("正在关闭Server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		config.SELF_LOG.Error("服务关闭失败:")
		config.SELF_LOG.Error(err.Error())
	}

	config.SELF_LOG.Info("服务已退出!")
	// config.SELF_LOG.Info("测试日志", zap.String("user", "nnmwqs"))
}
