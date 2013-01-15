// package github.com/mabetle/golog/slog. Simple Logger implement, Using fmt.Println
package slog

import(
	"fmt"
)

type SimpleLog struct{
	catalog string
	level string
}

func GetLogger(v ...interface{}) *SimpleLog{
	return &SimpleLog{}
}


func GetCatalog(l SimpleLog) string {
	if l.catalog=="" {
		l.catalog = "default"
	}
	return l.catalog
}


func (l SimpleLog)Error(v ...interface{}){
	/*fmt.Println("[Error]",": ",v)*/
	l.writeLog("ERROR",v)
}

func (l SimpleLog)Warn(v ...interface{}){
	l.writeLog("WARN",v)
}

func (l SimpleLog)Info(v ...interface{}){
	l.writeLog("INFO",v)
}

func (l SimpleLog)Debug(v ...interface{}){
	l.writeLog("DEBUG",v)
}

func (l SimpleLog)Trace(v ...interface{}){
	l.writeLog("TRACE",v)
}

func (l SimpleLog) writeLog(level string,v ...interface{}){
	//catalogLevel
	//logLevel
	//TODO: determine write or not by level.
	fmt.Println("[",level,"]",GetCatalog(l),": ",v)
}

