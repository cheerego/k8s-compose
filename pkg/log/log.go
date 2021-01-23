package log

import "go.uber.org/zap"

var logger = zap.NewExample()

func DefaultLogger() *zap.Logger {
	return logger
}
