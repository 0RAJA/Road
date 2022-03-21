package main

import (
	"context"
	"flag"
	"github.com/0RAJA/Bank/settings"
	"github.com/0RAJA/Road/internal/dao/mysql"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/logger"
	"github.com/0RAJA/Road/internal/pkg/setting"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {

}

//优雅关机
func gracefulExit(s *http.Server) {
	//退出通知
	quit := make(chan os.Signal)
	//等待退出通知
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	settings.Logger.Info("ShutDown Server...")
	//给几秒完成剩余任务
	ctx, cancel := context.WithTimeout(context.Background(), settings.ServerSetting.DefaultContextTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil { //优雅退出
		settings.Logger.Info("Server forced to ShutDown,Err:" + err.Error())
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

func SetupSetting() (err error) {
	setupFlag()
	newSetting, err := setting.NewSetting(configName, configType, strings.Split(configPaths, ",")...)
	if err != nil {
		return err
	}
	readSetting := func(k string, v interface{}) error {
		if err != nil {
			return err
		}
		return newSetting.ReadSection(k, v)
	}
	err = readSetting("Server", &global.AllSetting.Server)
	err = readSetting("Log", &global.AllSetting.Log)
	err = readSetting("App", &global.AllSetting.App)
	err = readSetting("Mysql", &global.AllSetting.Mysql)
	err = readSetting("Redis", &global.AllSetting.Redis)
	err = readSetting("Email", &global.AllSetting.Email)
	err = readSetting("Token", &global.AllSetting.Token)
	err = readSetting("Pagelines", &global.AllSetting.Pagelines)
	err = readSetting("Upload", &global.AllSetting.Upload)
	if err != nil {
		panic(err)
	}
	initLog()
	mysql.Init()
	redis.Init()
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
	global.Logger = logger.NewLogger(settings.LogSetting.Level)
}
