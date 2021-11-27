package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weirwei/golib/middleware"
	"github.com/weirwei/golib/wlog"
)

func main() {
	router := gin.Default()
	wlog.InitLog(wlog.LogConfig{
		Level:   "info",
		Stdout:  true,
		FileOut: true,
		Path:    "./logs",
	})
	router.Use(middleware.AccessLog())
	router.GET("/", func(context *gin.Context) {
		m := context.Query("msg")
		//Info级别的日志
		wlog.Infof(context, "get a info %s", m)
		//Error级别的日志
		wlog.Error(context, "get a error")
		//Warn级别的日志
		wlog.Warn(context, "get a warn")
		//Debug级别的日志
		wlog.Debug(context, "get a debug")
	})
	_ = router.Run()
}
