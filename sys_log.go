package golog

import(
	"fmt"
)

type SysLog struct{
	catalog string
}


func GetCatalog(l SysLog) string {
	if l.catalog=="" {
		l.catalog = "default"
	}
	return l.catalog
}


func (l SysLog)Error(v ...interface{}){
	/*fmt.Println("[Error]",": ",v)*/
	fmt.Println("[Error]",GetCatalog(l), ": ",v)
}
func (l SysLog)Warn(v ...interface{}){
	fmt.Println("[Warn]",v)
}
func (l SysLog)Info(v ...interface{}){
	fmt.Println("[Info]",v)
}
func (l SysLog)Debug(v ...interface{}){
	fmt.Println("[Debug]",v)
}
func (l SysLog)Trace(v ...interface{}){
	fmt.Println("[Trace]",v)
}
