package wlog

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	Level     string
	Path      string
	Formatter string
}

var config logConfig

type logConfig struct {
	Level     Level
	Path      string
	Formatter logrus.Formatter
}

type Level uint32

func init() {
	config.Level = LevelInfo
	config.Path = "./logs"
	config.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

func InitLog(c *LogConfig) {
	if c != nil {
		config.Level = getLevel(c.Level)
		config.Path = c.Path
		config.Formatter = getFormatter(c.Formatter)
	}
	LogrusLogger = newLogrus()
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
	return new(logrus.JSONFormatter)
}
