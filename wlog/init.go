package wlog

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// LogConfig
// Level 日志级别
// Stdout 是否在控制台打印
// FileOut 是否输出到文件
// Path 日志文件路径，默认根目录
// Formatter 日志格式，可选json 和 text，默认json
type LogConfig struct {
	Level     string `json:"level"`
	Stdout    bool   `json:"stdout"`
	FileOut   bool   `json:"fileOut"`
	Path      string `json:"path"`
	Formatter string `json:"formatter"`
}

var config logConfig

type logConfig struct {
	Level     Level            `json:"level"`
	Stdout    bool             `json:"stdout"`
	FileOut   bool             `json:"fileOut"`
	Path      string           `json:"path"`
	Formatter logrus.Formatter `json:"formatter"`
}

type Level uint32

func init() {
	config.Level = LevelInfo
	config.Stdout = true
	config.FileOut = true
	config.Path = "./log"
	config.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

// InitLog 初始化日志
func InitLog(c LogConfig) *logrus.Entry {
	config.Level = getLevel(c.Level)
	config.Stdout = c.Stdout
	config.FileOut = c.FileOut
	config.Path = c.Path
	config.Formatter = getFormatter(c.Formatter)
	LogrusLogger = logrus.NewEntry(newLogrus())
	return LogrusLogger
}

func getLevel(level string) Level {
	switch strings.ToTitle(level) {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	case "FATAL":
		return LevelFatal
	case "PANIC":
		return LevelPanic
	}
	return LevelInfo
}

func getFormatter(formatter string) logrus.Formatter {
	switch strings.ToTitle(formatter) {
	case "JSON":
		return &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	case "TEXT":
		return &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	}
	return config.Formatter
}
