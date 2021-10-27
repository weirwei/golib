package wlog

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

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
	writers := []io.Writer{
		src,
		os.Stdout}
	//同时写文件和控制台打印
	fileAndStdoutWriter := io.MultiWriter(writers...)
	//实例化
	logger := logrus.New()

	//设置输出
	logger.SetOutput(fileAndStdoutWriter)

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
