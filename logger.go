package golog

import(
	"io"
)

type Log struct{
	catalog string
	output io.Writer
	format string
}

type Logger interface{
	Error(v ...interface{})
	Warn(v ...interface{})
	Info(v ...interface{})
	Debug(v ...interface{})
	Trace(v ...interface{})
}

