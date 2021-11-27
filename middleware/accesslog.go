package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weirwei/golib/wlog"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginAccessLog(ctx *gin.Context) {
	// 默认钩子，添加requestID
	wlog.AddHook(&wlog.LogrusHook{})
	startTime := time.Now()
	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: ctx.Writer,
	}
	ctx.Writer = blw
	ctx.Next()
	statusCode := ctx.Writer.Status()
	response := blw.body.String()
	endTime := time.Now()
	latencyTime := endTime.Sub(startTime).Milliseconds()
	reqMethod := ctx.Request.Method
	reqUri := ctx.Request.RequestURI
	clientIP := ctx.ClientIP()
	wlog.InitFields(wlog.Fields{
		"status":   statusCode,
		"cost":     latencyTime,
		"method":   reqMethod,
		"uri":      reqUri,
		"clientIP": clientIP,
		"response": response,
	}).Infof("notice")
}

// AccessLog 每次请求信息汇总的中间件
func AccessLog() gin.HandlerFunc {
	return ginAccessLog
}
