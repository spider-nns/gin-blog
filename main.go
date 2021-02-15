package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
	"http-gin/global"
	"http-gin/internal/model"
	"http-gin/internal/routers"
	"http-gin/pkg/logger"
	"http-gin/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}
}

//初始化配置读取
func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
func setupLogger() error {
	global.Logger =
		logger.NewLogger(
			&lumberjack.Logger{
				Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
				MaxSize:   500,
				MaxAge:    30,
				LocalTime: true,
			}, "spider-GIN-#>", log.LstdFlags).WithCaller(2)
	return nil
}

// @title gin-blog
// @version v1
// @description Go 语言编程之旅:一起用Go写项目
// @termsOfService https://github.com/spider-nns/gin-blog
func main() {
	router := routers.NewRouter()
	server := &http.Server{
		Addr:              global.ServerSetting.HttpPort,
		Handler:           router,
		ReadTimeout:       global.ServerSetting.ReadTimeout,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}
	global.Logger.Infof("%s: http-gin/%s", "spider", "gin-blog")
	err := server.ListenAndServe()
	//
	//url := ginSwagger.URL("xxx")
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if err != nil {
		global.Logger.Errorf("%s: server start err:%v", err)
	}
}
