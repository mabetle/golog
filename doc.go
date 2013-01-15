/*
 Mabetle Go Log Package
 
 golog has a simple logger implement, using fmt.Println().

 golog make a bridge to 
  
	* seelog (github.com/cihub/seelog)
	* timber (github.com/ngmoco/timber)
	* log (in go lang package)

 Install
 =======

 go get -u github.com/mabetle/golog

 Usage In Code
 =============

 import "github.com/mabetle/golog"

 logger := golog.GetLogger()
 //logger :=golog.GetSeeLogger()
 //logger :=golog.GetGoLangLogger()
 //logger :=golog.GetTimberLogger()

 logger.Error("Error")
 logger.Warn("Warn")
 logger.Info("Info")
 logger.Debug("Debug")
 logger.Trace("Trace")

 Run Demo
 ========
 cd MABETLE_GOLOG_PATH
 go run cmd/demo/main.go



*/
package golog

