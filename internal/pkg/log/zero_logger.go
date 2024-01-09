package log

import (
	"fmt"
	"os"
	"server-go/internal/configs"
	"sync"

	"github.com/rs/zerolog"
)

var once sync.Once
var zeroSinLogger *zerolog.Logger

// ZeroLogger returns a singleton instance of zerolog.Logger
type zeroLogger struct {
	cfg    *configs.Config
	logger *zerolog.Logger
}

// Zero level logger map
var zeroLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
}

func (l *zeroLogger) Init() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = func(err error) interface{} {
			return map[string]interface{}{
				"error": err.Error(),
				"stack": err,
			}
		}
		fileName := fmt.Sprintf("%s/%s", l.cfg.Log.Path, l.cfg.Log.FileName)

		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("service", "server-go").
			Str("LoggerName", "ZeroLogger").
			Logger()

		zerolog.SetGlobalLevel(zeroLevelMap[l.cfg.Log.Level])
		zeroSinLogger = &logger
	})
	l.logger = zeroSinLogger
}

func newZeroLogger(cfg *configs.Config) *zeroLogger {
	l := &zeroLogger{
		cfg: cfg,
	}
	l.Init()
	return l
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	return zeroLevelMap[l.cfg.Log.Level]
}

func (l *zeroLogger) Debug(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Debug().Str("category", string(c)).Str("sub_category", string(sub)).Msg(msg)
}

func (l *zeroLogger) Info(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Info().Str("category", string(c)).Str("sub_category", string(sub)).Msg(msg)
}

func (l *zeroLogger) Warn(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Warn().Str("category", string(c)).Str("sub_category", string(sub)).Msg(msg)
}

func (l *zeroLogger) Error(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Error().Str("category", string(c)).Str("sub_category", string(sub)).Msg(msg)
}

func (l *zeroLogger) Fatal(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Fatal().Str("category", string(c)).Str("sub_category", string(sub)).Msg(msg)
}

func (l *zeroLogger) Panic(c Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Panic().Str("category", string(c)).Str("sub_category", string(sub)).Msg(msg)
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}

func (l *zeroLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panic().Msgf(template, args...)
}
