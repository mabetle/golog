package seelogb

import(
	"testing"
)

func TestSeeLog(t *testing.T){
	logger :=GetSeeLogger()

	logger.Error("Hello")
	logger.Warn("Hello")
	logger.Info("Hello")
	logger.Debug("Hello")
	logger.Trace("Hello")

}