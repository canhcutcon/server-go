package log

import "server-go/internal/configs"

type Logger interface {
	Init()

	Debug(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Warn(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})

	Panic(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Panicf(template string, args ...interface{})
}

func NewLogger(cfg *configs.Config) Logger {
	return newZeroLogger(cfg)
}
