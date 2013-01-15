package seelogb

import(
	"github.com/cihub/seelog"
)


type SeeLog struct{}

func GetLogger(v ...interface{}) *SeeLog{
	return &SeeLog{}
}

func (l SeeLog)Error(v ...interface{}){
	seelog.Error(v)
}


func (l SeeLog)Warn(v ...interface{}){
	seelog.Warn(v)
}

func (l SeeLog)Info(v ...interface{}){
	seelog.Info(v)
}

func (l SeeLog)Debug(v ...interface{}){
	seelog.Debug(v)
}

func (l SeeLog)Trace(v ...interface{}){
	seelog.Trace(v)
}


