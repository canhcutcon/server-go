package std

import (
	"fmt"
	"io"
	"log"
	"os"

	"server-go/internal/pkg/log/logger"
)

var _ logger.Logger = (*Logger)(nil)

type Logger struct {
	logger *log.Logger
	config *logger.ConfigLogger
}

// SetConfig implements logger.Logger.
func (*Logger) SetConfig(config *logger.ConfigLogger) error {
	panic("unimplemented")
}

var levelFormat = []string{
	"[DEBUG]",
	"[INFO]",
	"[WARN]",
	"[ERROR]",
	"[FATAL]",
} // this used to map the level with the string

func formatFields(fields logger.KV) string {
	str := ""
	for k, v := range fields {
		str += fmt.Sprintf("%s=%s ", k, v)
		if str[len(str)-1] == ' ' {
			str = str[:len(str)-1]
		}
	}
	return str
}

func NewInstance(config *logger.ConfigLogger) (*Logger, error) {
	return nil, nil
}

func New(config *logger.ConfigLogger) (*log.Logger, error) {
	stdLogger := &log.Logger{}
	if config == nil {
		config = &logger.ConfigLogger{
			Level:         logger.DebugLevel,
			LogFile:       "",
			TimeForFormat: logger.DefaultTimeFormat,
			Caller:        false,
			UseColor:      false,
			UseJSON:       false,
		}
		stdLogger.SetOutput(os.Stderr)
	}

	if config.LogFile != "" {
		file, err := logger.CreateLogFile(config.LogFile)
		if err != nil {
			return nil, err
		}
		writes := io.MultiWriter(os.Stderr, file)
		stdLogger.SetOutput(writes)
	}

	return stdLogger, nil
}

// SetConfig is used to reset the config for the logger
func (l *Logger) ResetLogger(config *logger.ConfigLogger) error {
	return nil
}

// SetLevel is used to reset log level for the logger
func (l *Logger) SetLevel(level logger.Level) error {
	return nil
}

// Print is used to print the message
func (l *Logger) Print(level logger.Level, args ...interface{}) {
	if level < l.config.Level {
		return
	}
	l.logger.Print(levelFormat[level], args)
}

func (l *Logger) Printf(level logger.Level, format string, args ...interface{}) {
	if level < l.config.Level {
		return
	}
	format = levelFormat[level] + " " + format
	l.logger.Printf(format, args)
}

func (l *Logger) Println(level logger.Level, args ...interface{}) {
	if level < l.config.Level {
		return
	}
	l.logger.Print(levelFormat[level], args)
}

func (l *Logger) Printw(level logger.Level, msg string, kv logger.KV) { // this used to print the message with key value
	if level < l.config.Level {
		return
	}
	l.logger.Println(levelFormat[level], msg, formatFields(kv))
}

func (l *Logger) Debug(args ...interface{}) {
	l.Print(logger.DebugLevel, args)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Printf(logger.DebugLevel, format, args)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.Println(logger.DebugLevel, args)
}

func (l *Logger) Debugw(msg string, kv logger.KV) {
	l.Printw(logger.DebugLevel, msg, kv)
}

func (l *Logger) Info(args ...interface{}) {
	l.Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l *Logger) Infow(msg string, kv logger.KV) {
	l.Infow(msg, kv)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Warn(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func (l *Logger) Warnw(msg string, kv logger.KV) {
	l.Warnw(msg, kv)
}

func (l *Logger) Error(args ...interface{}) {
	l.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func (l *Logger) Errorw(msg string, kv logger.KV) {
	l.Errorw(msg, kv)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
	fmt.Print(logger.FatalLevel, args)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func (l *Logger) Fatalw(msg string, kv logger.KV) {
	l.Fatalw(msg, kv)
}
