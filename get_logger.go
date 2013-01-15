package golog

import(
	"github.com/mabetle/golog/logb"
	"github.com/mabetle/golog/slog"
	"github.com/mabetle/golog/seelogb"
	"github.com/mabetle/golog/timberb"
)


func GetLogger(v ...interface{}) Logger{
	return logb.GetLogger()
}

func GetSimpleLogger(v ...interface{}) Logger{
	return slog.GetLogger("")
}

func GetSeeLogger(v ...interface{}) Logger{
	return seelogb.GetLogger("")
}

func GetGoLangLogger(v ...interface{}) Logger{
	return logb.GetLogger("")
}

func GetTimberLogger(v ...interface{}) Logger{
	return timberb.GetLogger("")
}



