package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weirwei/golib/wlog"
)

func main() {
	router := gin.Default()
	wlog.InitLog(&wlog.LogConfig{
		Level: "info",
		Path:  "./logs",
	})
	router.Use(wlog.LoggerToFile())
	router.GET("/", func(context *gin.Context) {
		//Info级别的日志
		wlog.Infof(context, "get a info %d", 1)
		//Error级别的日志
		wlog.Error(context, "get a error")
		//Warn级别的日志
		wlog.Warn(context, "get a warn")
		//Debug级别的日志
		wlog.Debug(context, "get a debug")
	})
	router.Run()
}
