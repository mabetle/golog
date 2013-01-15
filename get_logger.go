package golog




func GetLogger(v ...interface{}) Logger{
	return GetSysLogger(v ...)
}

func GetSysLogger(v ...interface{}) Logger{
	log := &SysLog(catalog:"root")
	return log
}



