package logb

import(
	"testing"
)

func TestGoLangLog(t *testing.T){

	logger := GetGoLangLogger()

	logger.Error("Hello")
	logger.Warn("Hello")
	logger.Info("Hello")
	logger.Debug("Hello")
	logger.Trace("Hello")
}
