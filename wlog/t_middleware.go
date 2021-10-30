package wlog

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
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
	// 开始时间
	startTime := time.Now()
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw
	ctx.Next()
	statusCode := ctx.Writer.Status()
	response := blw.body.String()
	endTime := time.Now()
	latencyTime := endTime.Sub(startTime)
	reqMethod := ctx.Request.Method
	reqUri := ctx.Request.RequestURI
	clientIP := ctx.ClientIP()
	WithField("status", statusCode)
	WithFields(Fields{
		"cost":     latencyTime,
		"method":   reqMethod,
		"uri":      reqUri,
		"clientIP": clientIP,
		"response": response,
	})
	Infof(ctx, "status")
}

func LoggerToFile() gin.HandlerFunc {
	return ginAccessLog
}
