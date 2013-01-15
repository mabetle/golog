/*
 package github.com/mabetle/golog. Logger for Go language.
 
 GoLog ideas are form SNF4j.

 GoLog make bridges to 
  
	* seelog (github.com/cihub/seelog)
	* timber (github.com/ngmoco/timber)
	* log (in go lang package)

 and  GoLog has a very simple logger implement, using fmt.Println().


 Install
 =======

 go get -u github.com/mabetle/golog

 Usage In Code
 =============

 import "github.com/mabetle/golog"

 logger := golog.GetLogger("")
 //logger :=golog.GetSeeLogger("")
 //logger :=golog.GetGoLangLogger("")
 //logger :=golog.GetTimberLogger("")

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

