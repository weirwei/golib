package middleware

import (
	"bytes"
	"io"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/weirwei/ikit/iutil"

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
	requestParam := getReqBody(ctx)
	query := getQuery(ctx)
	ctx.Next()
	statusCode := ctx.Writer.Status()
	response := blw.body.String()
	endTime := time.Now()
	latencyTime := endTime.Sub(startTime).Milliseconds()
	reqMethod := ctx.Request.Method
	reqUri := ctx.Request.RequestURI
	clientIP := ctx.ClientIP()
	wlog.InitFields(wlog.Fields{
		"status":       statusCode,
		"cost":         latencyTime,
		"method":       reqMethod,
		"uri":          reqUri,
		"clientIP":     clientIP,
		"requestParam": requestParam,
		"query":        query,
		"response":     response,
	}).Infof("notice")
}

// AccessLog 每次请求信息汇总的中间件
func AccessLog() gin.HandlerFunc {
	return ginAccessLog
}

func getReqBody(ctx *gin.Context) string {
	var res string
	if ctx.Request.Body != nil {
		requestBody, err := ctx.GetRawData()
		if err != nil {
			return ""
		}
		res = iutil.BytesString(requestBody)
		// 写回Request.Body
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	}
	return res
}
func getQuery(ctx *gin.Context) string {
	var queryStr string
	if len(ctx.Request.URL.RawQuery) == 0 {
		return ""
	}
	queryStr = ctx.Request.URL.RawQuery
	queries := strings.Split(queryStr, "&")
	res, err := jsoniter.MarshalToString(queries)
	if err != nil {
		return ""
	}
	return res
}
