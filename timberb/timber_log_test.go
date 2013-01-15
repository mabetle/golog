package timberb

import(
	"testing"
)

func TestTimberLog(t *testing.T){

	logger := GetLogger("golog")

	logger.Error("Hello")
	logger.Warn("Hello")
	logger.Info("Hello")
	logger.Debug("Hello")
	logger.Trace("Hello")

}

