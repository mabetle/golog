package golog

import(
	"github.com/mabetle/golog/logb"
	"github.com/mabetle/golog/slog"
	"github.com/mabetle/golog/seelogb"
	"github.com/mabetle/golog/timberb"
)

// Get a usable Logger
// TODO: make return Logger by config file or config Code
func GetLogger(v ...interface{}) Logger{
	return logb.GetLogger()
}

// Simple Logger, using fmt.Println()
func GetSimpleLogger(v ...interface{}) Logger{
	return slog.GetLogger("")
}

// Use seelog
func GetSeeLogger(v ...interface{}) Logger{
	return seelogb.GetLogger("")
}

// use go lang log
func GetGoLangLogger(v ...interface{}) Logger{
	return logb.GetLogger("")
}

// use timber log
func GetTimberLogger(v ...interface{}) Logger{
	return timberb.GetLogger("")
}



