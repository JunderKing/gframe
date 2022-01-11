package main

import (
	"gframe/config"
	"gframe/model"
	"gframe/pkg/logger"
	"gframe/pkg/redis"
	"gframe/router"
	"gframe/schedule"
	"os"
	"os/signal"
	"syscall"
)

// var db *gorm.DB

func main() {
	// 初始化
	config.Setup()
	// 数据库初始化
	model.Setup()
	// Redis初始化
	redis.Setup()
	// 定时任务初始化
	schedule.Setup()
	// 日志初始化
	logger.Setup()

	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()
}
