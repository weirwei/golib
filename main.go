package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weirwei/golib/middleware"
	"github.com/weirwei/golib/wlog"
)

type Req struct {
	Msg string `json:"msg" form:"msg"`
	ID  int    `json:"id" form:"id"`
}

type Resp struct {
	ErrNo  int    `json:"errNo"`
	ErrMsg string `json:"errMsg"`
	Data   string `json:"data"`
}

func main() {
	router := gin.Default()
	wlog.InitLog(wlog.LogConfig{
		Level:     "info",
		Stdout:    true,
		FileOut:   true,
		Path:      "./logs",
		Formatter: "json",
	})
	router.Use(middleware.AccessLog())
	router.GET("/get", func(context *gin.Context) {
		var req Req
		err := context.ShouldBindQuery(&req)
		if err != nil {
			return
		}
		wlog.Infof(context, "get a info %s", req.Msg)
		context.JSON(200, Resp{
			ErrNo:  0,
			ErrMsg: "",
			Data:   req.Msg,
		})
	})

	router.POST("/post", func(context *gin.Context) {
		var req Req
		err := context.ShouldBind(&req)
		if err != nil {
			return
		}
		wlog.Infof(context, "get a info %s", req.Msg)
		context.JSON(200, Resp{
			ErrNo:  0,
			ErrMsg: "",
			Data:   req.Msg,
		})
	})

	_ = router.Run()
}
