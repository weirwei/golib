package wlog

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	LogrusLogger *logrus.Logger
)

func newLogrus() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/" + config.Path
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	switch config.Level {
	case LevelPanic:
		logger.SetLevel(logrus.PanicLevel)
	case LevelFatal:
		logger.SetLevel(logrus.FatalLevel)
	case LevelError:
		logger.SetLevel(logrus.ErrorLevel)
	case LevelWarn:
		logger.SetLevel(logrus.WarnLevel)
	case LevelInfo:
		logger.SetLevel(logrus.InfoLevel)
	case LevelDebug:
		logger.SetLevel(logrus.DebugLevel)
	}

	//设置日志格式
	logger.SetFormatter(config.Formatter)
	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := newLogrus()
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
