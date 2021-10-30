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

func Debug(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Debug(args...)
}

func Debugf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Debugf(format, args...)
}

func Info(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Info(args...)
}

func Infof(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Infof(format, args...)
}

func Warn(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Warn(args...)
}

func Warnf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Warnf(format, args...)
}

func Error(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Error(args...)
}

func Errorf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Errorf(format, args...)
}

func Panic(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Panic(args...)
}

func Panicf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Panicf(format, args...)
}

func Fatal(ctx *gin.Context, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Fatal(args...)
}

func Fatalf(ctx *gin.Context, format string, args ...interface{}) {
	if !check(ctx) {
		return
	}
	weirLog(ctx).Fatalf(format, args...)
}

func check(ctx *gin.Context) bool {
	return ctx != nil
}
