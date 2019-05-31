package main

import (
	"fmt"
	"github.com/hequan2017/go-admin/middleware/inject"
	"github.com/hequan2017/go-admin/models"
	"github.com/hequan2017/go-admin/pkg/logging"
	"github.com/hequan2017/go-admin/pkg/setting"
	"github.com/hequan2017/go-admin/routers"
	"log"
	"net/http"
)

// @title go-admin
// @version 1.2.1
// @description  go-admin
// @termsOfService https://github.com/hequan2017/go-admin

// @contact.name hequan
// @contact.url https://github.com/hequan2017
// @contact.email hequan2011@sina.com

// @license.name MIT
// @license.url  https://github.com/hequan2017/go-admin/blob/master/LICENSE

// @host   127.0.0.1:8000
// @BasePath
func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	inject.Init()
	err := inject.LoadCasbinPolicyData()
	if err != nil {
		panic("加载casbin策略数据发生错误: " + err.Error())
	}

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
