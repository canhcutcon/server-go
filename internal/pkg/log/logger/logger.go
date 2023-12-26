package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type (
	// Logger is the interface for the logger
	// KV is the key value pair for the logger
	// this used to pass the data to the logger with function
	KV    map[string]interface{}
	Level int // level of logger

	ConfigLogger struct {
		Level         Level
		LogFile       string
		TimeForFormat string
		Caller        bool
		UseColor      bool
		UseJSON       bool
	}

	Logger interface {
		// Debug logs a message at level Debug on the standard logger.
		SetConfig(config *ConfigLogger) error      // this used to set the config for the logger
		SetLevel(level Level) error                // this used to set the level for the logger
		Debug(args ...interface{})                 // this used to log the debug message
		Debugf(format string, args ...interface{}) // this used to log the debug message with format
		Info(args ...interface{})                  // this used to log the info message
		Infof(format string, args ...interface{})  // this used to log the info message with format
		Infow(msg string, kv KV)                   // this used to log the info message with key value
		Warn(args ...interface{})                  // this used to log the warn message
		Warnf(format string, args ...interface{})
		Warnw(msg string, KV KV)
		Error(args ...interface{})
		Errorf(format string, args ...interface{})
		Errorw(msg string, KV KV)
		Fatal(args ...interface{}) // this used to log the fatal message
		Fatalf(format string, args ...interface{})
		Fatalw(msg string, KV KV)
	}
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	DebugLevel Level = iota // this is the debug level, it will log all the message, iota is the auto increment
	// InfoLevel is the default logging priority.
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

const (
	DebugLevelString = "debug"
	InfoLevelString  = "info"
	WarnLevelString  = "warn"
	ErrorLevelString = "error"
	FatalLevelString = "fatal"
)

const DefaultTimeFormat = time.RFC3339 // this is the default time format, time.RFC3339 is the format for the time

// StringToLevel is the function to convert the string to level
func StringToLevel(level string) Level {
	switch level {
	case DebugLevelString:
		return DebugLevel
	case InfoLevelString:
		return InfoLevel
	case WarnLevelString:
		return WarnLevel
	case ErrorLevelString:
		return ErrorLevel
	case FatalLevelString:
		return FatalLevel
	default:
		return InfoLevel
	}
}

func CreateLogFile(filename string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(filename), 0744) // this used to create the directory for the log file
	if err != nil {
		fmt.Println("error while creating the directory for the log file", err)
		return nil, err
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // this used to create the log file
	if err != nil {
		fmt.Println("error while creating the log file", err)
		return nil, err
	}
	return file, nil
}
