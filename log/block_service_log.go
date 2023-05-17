package log

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// Get zap log instance.
func Logger() *zap.Logger {
	once.Do(func() {
		InitLogger("info")
	})
	return logger
}
