package main

import (
	"SysZ/src/config"
	"SysZ/src/config/LoadConfigFactory"
	"SysZ/src/controller/interception"
	"SysZ/src/exception"
	"SysZ/src/image"
	"SysZ/src/pool"
	"SysZ/src/router"
	"SysZ/src/service"
	"SysZ/src/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			util.Loglevel(util.Error, "main", util.Strval(err))
		}
	}()
	gin.SetMode(gin.DebugMode)

	// 初始化日志模块
	util.LoggerInit(func(r any) {
		util.Loglevel(util.Error, "main", util.Strval(r))
		exception.HandleException(exception.NewSystemError("日志模块", util.Strval(r)))
	}, util.Debug)
	util.Loglevel(util.Debug, "main", "初始化日志模块")

	//初始化配置模块
	if err := LoadConfigFactory.InitConfig(); err != nil {
		exception.HandleException(err)
	}
	util.Loglevel(util.Debug, "main", "初始化配置文件")

	//初始化协程池
	if err := pool.InitRoutinePool(); err != nil {
		exception.HandleException(err)
	}
	util.Loglevel(util.Debug, "ThemisInitFactory", "初始化协程池")

	//画图
	time.Sleep(time.Second * 1)
	image.Factory()
	time.Sleep(time.Second * 1)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(interception.Interception())

	router.API(r)

	// 初始化服务模块
	if err := service.InitServer(); err != nil {
		exception.HandleException(err)
	}

	go func() {
		util.Loglevel(util.Info, "main", "start service"+config.Port)
		err := r.Run(":" + config.Port)
		if err != nil {
			util.Loglevel(util.Error, "main", util.Strval(err))
			util.Loglevel(util.Error, "main", "server start error")
			os.Exit(0)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	util.Loglevel(util.Info, "main", "SysZ is exiting...")

	// 关闭服务模块
	if err := service.Close(); err != nil {
		exception.HandleException(err)
	}

	runtime.GC()
	util.Loglevel(util.Info, "main", "SysZ is exited")
	time.Sleep(time.Second * 3)
}
