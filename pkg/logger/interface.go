package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log interface
type LogI interface {
	Add(key string, field interface{}) *Log
	Append(fields ...zapcore.Field) *Log
	Level(level zapcore.Level) *Log
	Development() *Log
	Namespace(name string, fields ...zapcore.Field) *Log
	Commit(string)
	Get() *Log
}

type Log struct {
	logger *zap.Logger
	fields []zapcore.Field
	level  zapcore.Level
	path   string
	debug  bool
}

type LoggerOption func(l *Log)
