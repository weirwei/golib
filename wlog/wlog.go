package wlog

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func weirLog(ctx *gin.Context) *logrus.Entry {
	if ctx == nil {
		return LogrusLogger
	}
	return LogrusLogger
}

// Debug debug level log
func Debug(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Debug(args...)
}

// Debugf debug level log with format
func Debugf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Debugf(format, args...)
}

// Info  info level log
func Info(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Info(args...)
}

// Infof  info level log with format
func Infof(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Infof(format, args...)
}

// Warn warn level log
func Warn(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Warn(args...)
}

// Warnf warn level log with format
func Warnf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Warnf(format, args...)
}

// Error error level log
func Error(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Error(args...)
}

// Errorf error level log with format
func Errorf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Errorf(format, args...)
}

// Panic panic level log
func Panic(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Panic(args...)
}

// Panicf panic level log with format
func Panicf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Panicf(format, args...)
}

// Fatal fatal level log
func Fatal(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Fatal(args...)
}

// Fatalf fatal level log with format
func Fatalf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Fatalf(format, args...)
}

func check(ctx *gin.Context) bool {
	return ctx != nil
}
