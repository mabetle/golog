package logb


import(
	"log"
)

type GoLangLog struct{

}


func GetGoLangLogger() *GoLangLog{
	return new(GoLangLog)
}

func (sl GoLangLog)Error(v ...interface{}){
	log.Println(v)
}


func (sl GoLangLog)Warn(v ...interface{}){
	log.Println(v)
}

func (sl GoLangLog)Info(v ...interface{}){
	log.Println(v)
}

func (sl GoLangLog)Debug(v ...interface{}){
	log.Println(v)
}

func (sl GoLangLog)Trace(v ...interface{}){
	log.Println(v)
}

