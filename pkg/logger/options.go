package logger

import "os"

func WithPath(path string) LoggerOption {
	return func(l *Log) {
		l.path = path
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		}

	}
}

func WithDebug(debug bool) LoggerOption {
	return func(l *Log) {
		l.debug = debug
	}
}
