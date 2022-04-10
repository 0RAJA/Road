package main

import (
	"context"
	"flag"
	"github.com/0RAJA/Road/internal/dao/mysql"
	"github.com/0RAJA/Road/internal/dao/redis"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/logger"
	"github.com/0RAJA/Road/internal/pkg/setting"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/0RAJA/Road/internal/pkg/token"
	"github.com/0RAJA/Road/internal/pkg/upload"
	"github.com/0RAJA/Road/internal/routing"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// @title           Road
// @version         1.0
// @description     The Road Of Code

// @license.name  Raja
// @license.url   https://github.com/0RAJA

// @host      humraja.com
// @BasePath  /road/

// @securityDefinitions.basic  BasicAuth
func main() {
	SetupSetting()
	if global.AllSetting.Server.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := routing.NewRouting()
	s := &http.Server{
		Addr:           global.AllSetting.Server.Address,
		Handler:        r,
		ReadTimeout:    global.AllSetting.Server.ReadTimeout,
		WriteTimeout:   global.AllSetting.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			global.Logger.Info(err.Error())
		}
	}()
	gracefulExit(s)
	global.Logger.Info("Exit!")
}

//优雅关机
func gracefulExit(s *http.Server) {
	//退出通知
	quit := make(chan os.Signal)
	//等待退出通知
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("ShutDown Server...")
	//给几秒完成剩余任务
	ctx, cancel := context.WithTimeout(context.Background(), global.AllSetting.Server.DefaultContextTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil { //优雅退出
		global.Logger.Info("Server forced to ShutDown,Err:" + err.Error())
	}
}

var (
	configPaths string
	configName  string
	configType  string
)

func setupFlag() {
	//命令行参数绑定
	flag.StringVar(&configName, "name", "app", "配置文件名")
	flag.StringVar(&configType, "type", "yml", "配置文件类型")
	flag.StringVar(&configPaths, "path", global.RootDir+"/configs/app", "指定要使用的配置文件路径")
	flag.Parse()
}

func SetupSetting() {
	setupFlag()
	newSetting, err := setting.NewSetting(configName, configType, strings.Split(configPaths, ",")...)
	myPanic(err)
	readSetting := func(k string, v interface{}) error {
		if err != nil {
			return err
		}
		return newSetting.ReadSection(k, v)
	}
	err = readSetting("Server", &(global.AllSetting.Server))
	err = readSetting("Log", &(global.AllSetting.Log))
	err = readSetting("App", &(global.AllSetting.App))
	err = readSetting("Mysql", &(global.AllSetting.Mysql))
	err = readSetting("Redis", &(global.AllSetting.Redis))
	err = readSetting("Email", &(global.AllSetting.Email))
	err = readSetting("Token", &(global.AllSetting.Token))
	err = readSetting("Pagelines", &(global.AllSetting.Pagelines))
	err = readSetting("Upload", &(global.AllSetting.Upload))
	err = readSetting("Github", &(global.AllSetting.Github))
	err = readSetting("Rule", &(global.AllSetting.Rule))
	myPanic(err)
	err = snowflake.Init(global.AllSetting.App.StartTime, global.AllSetting.App.Format, 1)
	myPanic(err)
	global.Maker, err = token.NewPasetoMaker([]byte(global.AllSetting.Token.Key))
	myPanic(err)
	initLog()
	initUpload()
	mysql.QueryInit(global.AllSetting.Mysql.DriverName, global.AllSetting.Mysql.SourceName)
	redis.QueryInit(global.AllSetting.Redis.Address, global.AllSetting.Redis.Password, global.AllSetting.Redis.PoolSize, global.AllSetting.Redis.DB)
	app.Init(global.AllSetting.Pagelines.DefaultPageSize, global.AllSetting.Pagelines.MaxPageSize, global.AllSetting.Pagelines.PageKey, global.AllSetting.Pagelines.PageSizeKey)
}
func myPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func initLog() {
	logger.Init(&logger.InitStruct{
		LogSavePath:   global.AllSetting.Log.LogSavePath,
		LogFileExt:    global.AllSetting.Log.LogFileExt,
		MaxSize:       global.AllSetting.Log.MaxSize,
		MaxBackups:    global.AllSetting.Log.MaxBackups,
		MaxAge:        global.AllSetting.Log.MaxAge,
		Compress:      global.AllSetting.Log.Compress,
		LowLevelFile:  global.AllSetting.Log.LowLevelFile,
		HighLevelFile: global.AllSetting.Log.HighLevelFile,
	})
	global.Logger = logger.NewLogger(global.AllSetting.Log.Level)
}

func initUpload() {
	image := upload.NewType(upload.FileType(global.AllSetting.Upload.Image.Type), global.AllSetting.Upload.Image.Suffix, global.AllSetting.Upload.Image.MaxSize, global.AllSetting.Upload.Image.UrlPrefix, global.AllSetting.Upload.Image.LocalPath)
	file := upload.NewType(upload.FileType(global.AllSetting.Upload.File.Type), global.AllSetting.Upload.File.Suffix, global.AllSetting.Upload.File.MaxSize, global.AllSetting.Upload.File.UrlPrefix, global.AllSetting.Upload.File.LocalPath)
	global.Upload = upload.Init(image, file)
}
