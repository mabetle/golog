package golog

import(
	"testing"
)

var(
	logger Logger = GetLogger("test")
)

func TestLog(t *testing.T) {
	logger.Error("Hello")
	logger.Warn("Hello")
	logger.Info("Hello")
	logger.Debug("Hello")
	logger.Trace("Hello")
}

