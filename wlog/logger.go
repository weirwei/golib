package wlog

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	// Fields 日志自定义kv
	Fields = logrus.Fields

	// Hook 钩子接口
	Hook = logrus.Hook
)

var (
	LogrusLogger *logrus.Entry
)

// AddFields 添加一组kv
func AddFields(fields Fields) {
	LogrusLogger = LogrusLogger.WithFields(fields)
}

// InitFields 返回一个新的Entry，并设置一个组的kv
func InitFields(fields Fields) *logrus.Entry {
	return LogrusLogger.WithFields(fields)
}

// AddField 添加kv
func AddField(key string, value interface{}) {
	LogrusLogger = LogrusLogger.WithField(key, value)
}

// InitField 返回一个新的Entry，并设置一个新的kv
func InitField(key string, value interface{}) *logrus.Entry {
	return LogrusLogger.WithField(key, value)
}

// AddHook 添加钩子
func AddHook(hook Hook) {
	LogrusLogger.Logger.Hooks.Add(hook)
}

func newLogrus() *logrus.Logger {
	now := time.Now()
	var logFilePath string
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
			panic(fmt.Errorf("log conf err: create log file '%s' error: %v", fileName, err))
		}
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(fmt.Errorf("log conf err: open log file '%s' error: %v", fileName, err))
	}
	var writers []io.Writer
	if config.Stdout {
		writers = append(writers, os.Stdout)
	}
	if config.FileOut {
		writers = append(writers, src)
	}
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
