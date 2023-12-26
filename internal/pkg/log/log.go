package log

import (
	"errors"
	"server-go/internal/pkg/log/logger"
	"server-go/internal/pkg/log/logger/std"
)

type Config struct {
	Level         string
	LogFile       string
	TimeForFormat string
	Caller        bool
	UseColor      bool
	UseJSON       bool
}

var (
	// DefaultTimeFormat is the default time format for the logger
	_debugLogger logger.Logger
	_infoLogger  logger.Logger
	_warnLogger  logger.Logger
	_errorLogger logger.Logger
	_fatalLogger logger.Logger

	errInvalidLevel = errors.New("invalid level")
	errInvalidLog   = errors.New("invalid log")
)

func SetLogger(logger logger.Logger) error {
	if logger == nil {
		return errInvalidLog
	}
	_debugLogger = logger
	_infoLogger = logger
	_warnLogger = logger
	_errorLogger = logger
	_fatalLogger = logger
	return nil
}

func Init() {
	serverLog, err := std.New(nil)
	if err != nil {
		return
	}
	SetLogger(serverLog)
}

func SetConfig(config *logger.ConfigLogger) error {
	if config == nil {
		return errInvalidLog
	}

	if err := _debugLogger.SetConfig(config); err != nil {
		return err
	}

	if err := _infoLogger.SetConfig(config); err != nil {
		return err
	}

	if err := _warnLogger.SetConfig(config); err != nil {
		return err
	}

	if err := _errorLogger.SetConfig(config); err != nil {
		return err
	}

	if err := _fatalLogger.SetConfig(config); err != nil {
		return err
	}
	return nil
}

func SetLevel(level logger.Level) {
	_debugLogger.SetLevel(level)
	_infoLogger.SetLevel(level)
	_warnLogger.SetLevel(level)
	_errorLogger.SetLevel(level)
	_fatalLogger.SetLevel(level)
}

func SetLevelLogger(level logger.Level) {
	SetLevel(level)
}

func SetLevelLoggerString(level string) {
	SetLevel(logger.StringToLevel(level))
}

// Debug function
func Debug(args ...interface{}) {
	_debugLogger.Debug(args...)
}

// Debugf function
func Debugf(format string, v ...interface{}) {
	_debugLogger.Debugf(format, v...)
}

// Debugw function
func Debugw(msg string, keyValues logger.KV) {
	_debugLogger.Debugw(msg, keyValues)
}

// Print function
func Print(v ...interface{}) {
	_infoLogger.Info(v...)
}

// Println function
func Println(v ...interface{}) {
	_infoLogger.Info(v...)
}

// Printf function
func Printf(format string, v ...interface{}) {
	_infoLogger.Infof(format, v...)
}

// Info function
func Info(args ...interface{}) {
	_infoLogger.Info(args...)
}

// Infof function
func Infof(format string, v ...interface{}) {
	_infoLogger.Infof(format, v...)
}

// Infow function
func Infow(msg string, keyValues logger.KV) {
	_infoLogger.Infow(msg, keyValues)
}

// Warn function
func Warn(args ...interface{}) {
	_warnLogger.Warn(args...)
}

// Warnf function
func Warnf(format string, v ...interface{}) {
	_warnLogger.Warnf(format, v...)
}

// Warnw function
func Warnw(msg string, keyValues logger.KV) {
	_warnLogger.Warnw(msg, keyValues)
}

// Error function
func Error(args ...interface{}) {
	_errorLogger.Error(args...)
}

// Errorf function
func Errorf(format string, v ...interface{}) {
	_errorLogger.Errorf(format, v...)
}

// Errorw function
func Errorw(msg string, keyValues logger.KV) {
	_errorLogger.Errorw(msg, keyValues)
}

// Fatal function
func Fatal(args ...interface{}) {
	_fatalLogger.Fatal(args...)
}

// Fatalf function
func Fatalf(format string, v ...interface{}) {
	_fatalLogger.Fatalf(format, v...)
}

// Fatalw function
func Fatalw(msg string, keyValues logger.KV) {
	_fatalLogger.Fatalw(msg, keyValues)
}
