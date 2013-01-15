package timberb

import(
	"testing"
)

func TestTimberLog(t *testing.T){

	logger := GetTimberLogger("golog")

	logger.Error("Hello")
	logger.Warn("Hello")
	logger.Info("Hello")
	logger.Debug("Hello")
	logger.Trace("Hello")

}

