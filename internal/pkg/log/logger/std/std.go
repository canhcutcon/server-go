package std

import (
	"fmt"
	"io"
	"os"
	"server-go/internal/pkg/log/logger"
)

var _ logger.Logger = (*logger.Logger)(nil) // this used to check the logger is the pointer to the logger

var Logger struct {
	logger *logger.Logger
	config *logger.ConfigLogger
}

var levelFormat = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
} // this used to map the level with the string

// SetConfig is used to set the config for the logger
func SetConfig(config *logger.ConfigLogger) error {
	Logger.config = config
	return nil
}

func formatFields(fields logger.KV) string {
	var str string
	for k, v := range fields {
		str += fmt.Sprintf("%s=%s ", k, v)
		if str[len(str)-1] == ' ' {
			str = str[:len(str)-1]
		}
	}
	return str
}

func New(config *logger.ConfigLogger) (*Logger, error) {
	stdLogger, err := newLogger(config)
	if err != nil {
		return nil, err
	}
	return &Logger{
		logger: stdLogger,
		config: config,
	}, nil
}

func newLogger(config *logger.ConfigLogger) (*logger.Logger, error) { // this used to create the new logger
	stdLogger := &logger.Logger{} // this used to create the new logger, &logger.NewLogger is the pointer to the logger

	if config == nil { // this used to check the config is nil or not
		// if the config is nil, then set the default config
		config = &logger.ConfigLogger{
			Level:      logger.InfoLevel,
			TimeFormat: logger.DefaultTimeFormat,
		}
		stdLogger.SetOutput(os.Stderr) // this used to set the output for the logger, os.Stderr is the standard error
	}

	if config.TimeForFormat == "" {
		config.TimeForFormat = logger.DefaultTimeFormat
	}

	if config.LogFile != "" {
		file, err := logger.CreateLogFile(config.LogFile)
		if err != nil {
			return nil, err
		}

		writers := io.MultiWriter(os.Stderr, file)
		stdLogger.SetOutput(writers)
	}

	return stdLogger, nil
}

// SetConfig is used to set the config for the logger
func (l *Logger) SetConfig(config *logger.ConfigLogger) error {
	l.config = config
	return nil
}

// SetLevel is used to set the level for the logger
func (l *Logger) SetLevel(level logger.Level) error {
	l.config.Level = level
	return nil
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Print(logger.DebugLevel, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Printf(logger.DebugLevel, format, args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.logger.Println(logger.DebugLevel, args)
}

func (l *Logger) Debugw(msg string, kv logger.KV) {
	l.logger.Printw(logger.DebugLevel, msg, formatFields(kv))
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
	fmt.Print(logger.InfoLevel, args)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
	fmt.Print(logger.InfoLevel, format, args)
}

func (l *Logger) Infow(msg string, kv logger.KV) {
	l.logger.Infow(msg, kv)
	fmt.Print(logger.InfoLevel, msg, kv)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
	fmt.Print(logger.WarnLevel, args)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
	fmt.Print(logger.WarnLevel, format, args)
}

func (l *Logger) Warnw(msg string, kv logger.KV) {
	l.logger.Warnw(msg, kv)
	fmt.Print(logger.WarnLevel, msg, kv)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
	fmt.Print(logger.ErrorLevel, args)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
	fmt.Print(logger.ErrorLevel, format, args)
}

func (l *Logger) Errorw(msg string, kv logger.KV) {
	l.logger.Errorw(msg, kv)
	fmt.Print(logger.ErrorLevel, msg, kv)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
	fmt.Print(logger.FatalLevel, args)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
	fmt.Print(logger.FatalLevel, format, args)
}

func (l *Logger) Fatalw(msg string, kv logger.KV) {
	l.logger.Fatalw(msg, kv)
	fmt.Print(logger.FatalLevel, msg, kv)
}

// Print is used to print the message
func (l *Logger) Print(level logger.Level, args ...interface{}) {
	if level < l.config.Level {
		return
	}
	l.logger.Print(levelFormat[level], args...)
}

func (l *Logger) Printf(level logger.Level, format string, args ...interface{}) {
	if level < l.config.Level {
		return
	}
	format = levelFormat[level] + " " + format
	l.logger.Printf(format, args...)
}

func (l *Logger) Println(level logger.Level, args ...interface{}) {
	if level < l.config.Level {
		return
	}
	l.logger.Printw(levelFormat[level], args)
}

func (l *Logger) Printw(level logger.Level, msg string, kv logger.KV) { // this used to print the message with key value
	if level < l.config.Level {
		return
	}
	l.logger.Printw(levelFormat[level], msg, formatFields(kv))
}
