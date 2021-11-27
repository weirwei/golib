package wlog

import "github.com/sirupsen/logrus"

type LogrusHook struct {
	requestID string
}

// Levels 设置所有的日志等级都走这个钩子
func (hook *LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 每次请求都添加唯一的requestID
func (hook *LogrusHook) Fire(entry *logrus.Entry) error {
	if len(hook.requestID) == 0 {
		hook.requestID = getRequestID()
	}
	entry.Data["requestID"] = hook.requestID
	return nil
}
