package wlog

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

type Fields = logrus.Fields

var (
	LogrusLogger *logrus.Entry
)

func WithFields(fields Fields) {
	LogrusLogger = LogrusLogger.WithFields(fields)
}

func WithField(key string, value interface{}) {
	LogrusLogger = LogrusLogger.WithField(key, value)
}

func newLogrus() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/" + config.Path
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("20060102") + ".log"
	// 写入日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writers := []io.Writer{
		src,
		os.Stdout}
	// 同时写文件和控制台打印
	fileAndStdoutWriter := io.MultiWriter(writers...)

	logger := logrus.New()

	logger.SetOutput(fileAndStdoutWriter)

	// 设置日志级别
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
