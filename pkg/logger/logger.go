package logger

import "go.uber.org/zap"

var (
	logger *zap.Logger
)

// NewPrototype new logger prototype
func NewPrototype(opts ...LoggerOption) LogI {

	l := new(Log)

	for _, opt := range opts {
		opt(l)
	}

	if logger == nil {
		logger = getZapLogger(l.debug, l.path)
	}
	l.logger = logger

	return l
}
