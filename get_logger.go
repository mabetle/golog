package golog

import(
	"github.com/mabetle/golog/logb"
)


func GetLogger(v ...interface{}) Logger{
	return logb.GetGoLangLogger()
}





