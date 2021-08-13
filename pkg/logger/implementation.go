package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// GetZapLogger func
// create new zap logger
func getZapLogger(debug bool, path string) *zap.Logger {

	if path == "" {
		path = "."
	}

	w := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s.log", path, time.Now().Local().Format("2006-01-02")),
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		},
	)

	prodEC := zap.NewProductionEncoderConfig()
	prodEC.EncodeTime = zapcore.RFC3339TimeEncoder

	if debug {
		core := zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(prodEC),
				w,
				zap.InfoLevel,
			),
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(prodEC),
				zapcore.AddSync(os.Stdout),
				zapcore.DebugLevel,
			),
		)
		logger = zap.New(core)

	} else {

		core := zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(prodEC),
				w,
				zap.InfoLevel,
			),
		)

		logger = zap.New(core)
	}

	return logger
}

// Add new log
func (l *Log) Add(key string, field interface{}) *Log {
	l.fields = append(l.fields, zap.Any(key, field))
	return l
}

// Append new log
func (l *Log) Append(fields ...zapcore.Field) *Log {
	l.fields = append(l.fields, fields...)
	return l
}

// Commit meth
func (l *Log) Commit(message string) {
	defer func() {
		l.logger.Sync()
		l.fields = nil
	}()

	if l.fields != nil {
		switch l.level {
		case zapcore.InfoLevel:
			l.logger.Info(message, l.fields...)
		case zapcore.WarnLevel:
			l.logger.Warn(message, l.fields...)
		case zapcore.DebugLevel:
			l.logger.Debug(message, l.fields...)
		case zapcore.FatalLevel:
			l.logger.Fatal(message, l.fields...)
		case zapcore.ErrorLevel:
			l.logger.Error(message, l.fields...)
		default:
			l.logger.Info(message, l.fields...)
		}
	}

}

// Level of log
func (l *Log) Level(level zapcore.Level) *Log {
	l.level = level
	return l
}

// Development method
func (l *Log) Development() *Log {
	var caller string = ""

	pc, _, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		caller = details.Name()
	}

	l.Add("line", line)
	l.Add("caller", caller)
	return l
}

func (l *Log) Get() *Log {
	return l
}

func (l *Log) Namespace(name string, fields ...zapcore.Field) *Log {
	l.logger.With(append(fields, zap.Namespace(name))...)
	return l
}
