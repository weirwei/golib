package wlog

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// LogConfig
// Level 日志级别，default: INFO
// Stdout 是否在控制台打印，default: false
// FileOut 是否输出到文件，default: false
// Path 日志文件路径，default: ./log
// Formatter 日志格式，可选json 和 text，default: json
type LogConfig struct {
	Level     string `json:"level"`
	Stdout    bool   `json:"stdout"`
	FileOut   bool   `json:"fileOut"`
	Path      string `json:"path"`
	Formatter string `json:"formatter"`
}

func (l LogConfig) getLevel() {
	switch strings.ToTitle(l.Level) {
	case "DEBUG":
		config.Level = LevelDebug
	case "INFO":
		config.Level = LevelInfo
	case "WARN":
		config.Level = LevelWarn
	case "ERROR":
		config.Level = LevelError
	case "FATAL":
		config.Level = LevelFatal
	case "PANIC":
		config.Level = LevelPanic
	}
}

func (l LogConfig) getPath() {
	if len(l.Path) > 0 {
		config.Path = l.Path
	}
}

func (l LogConfig) getFormatter() {
	switch strings.ToTitle(l.Formatter) {
	case "JSON":
		config.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	case "TEXT":
		config.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	}
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
	config.Stdout = c.Stdout
	config.FileOut = c.FileOut
	c.getLevel()
	c.getPath()
	c.getFormatter()
	LogrusLogger = logrus.NewEntry(newLogrus())
	return LogrusLogger
}
