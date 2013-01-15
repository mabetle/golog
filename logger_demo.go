package golog


func DemoLogger(logger Logger){
	logger.Error("Hello")
	logger.Warn("Hello")
	logger.Info("Hello")
	logger.Debug("Hello")
	logger.Trace("Hello")
}


