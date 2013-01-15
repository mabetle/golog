package slog

import(
	"fmt"
)

type SimpleLog struct{
	catalog string
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
	fmt.Println("[Error]",GetCatalog(l), ": ",v)
}
func (l SimpleLog)Warn(v ...interface{}){
	fmt.Println("[Warn]",v)
}
func (l SimpleLog)Info(v ...interface{}){
	fmt.Println("[Info]",v)
}
func (l SimpleLog)Debug(v ...interface{}){
	fmt.Println("[Debug]",v)
}
func (l SimpleLog)Trace(v ...interface{}){
	fmt.Println("[Trace]",v)
}
