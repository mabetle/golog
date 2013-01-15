package timberb

import(
	tl "github.com/ngmoco/timber"
)

type TimberLog struct{
}

func GetLogger(v ...interface{}) *TimberLog {
	return new(TimberLog)
}


func (log TimberLog) Error(v ...interface{}){
	tl.Error("",v)
}


func (log TimberLog) Warn(v ...interface{}){
	tl.Error("",v)
}

func (log TimberLog) Info(v ...interface{}){
	tl.Error("",v)
}


func (log TimberLog) Debug(v ...interface{}){
	tl.Debug("",v)
}


func (log TimberLog) Trace(v ...interface{}){
	tl.Trace("",v)
}

