package wlog

import "github.com/sirupsen/logrus"

type LogrusHook struct {
}

//设置所有的日志等级都走这个钩子
func (hook *LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//修改其中的数据，或者进行其他操作
func (hook *LogrusHook) Fire(entry *logrus.Entry) error {
	entry.Data["requestID"] = getRequestID()
	return nil
}
