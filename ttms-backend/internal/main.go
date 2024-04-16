/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:07
 */

package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mognolia/internal/global"
	"mognolia/internal/routing"
	"mognolia/internal/setting"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func initSettings() {
	setting.AllInit()
}

// @title        ttms
// @version      1.0
// @description  ttms影院管理系统

// @license.name
// @license.url

// @host      127.0.0.1
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	initSettings()
	if global.Settings.Serve.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := routing.NewRouter() // 注册路由
	s := &http.Server{
		Addr:           global.Settings.Serve.Address,
		Handler:        r,
		ReadTimeout:    global.Settings.Serve.ReadTimeout,
		WriteTimeout:   global.Settings.Serve.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Info("Server started!")
	fmt.Println("AppName:", global.Settings.App.Name, "Version:", global.Settings.App.Version, "Address:", global.Settings.Serve.Address, "RunMode:", global.Settings.Serve.RunMode)
	setting.Group.Auto.Init()

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			global.Logger.Info(err.Error())
		}
	}()

	gracefulExit(s) // 优雅退出
	global.Logger.Info("Server exited!")
}

// 优雅退出
func gracefulExit(s *http.Server) {
	// 退出通知
	quit := make(chan os.Signal, 1)
	// 等待退出通知
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("ShutDown Server...")
	// 给几秒完成剩余任务
	ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Serve.DefaultContextTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil { // 优雅退出
		global.Logger.Info("Server forced to ShutDown,Err:" + err.Error())
	}
}
