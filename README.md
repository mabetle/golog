Mabetle GoLog
=============

GoLog is a logger for Go Language. The ideas are from SLF4J.

GoLog have bridges to:

	* SeeLog [github.com/cihub/seelog]
	* Timber [github.com/ngmoco/timber]
	* log (from go language package)

and GoLog has a very simple logger implement using fmt.Println()


Install
------

	go get -u github.com/mabetle/golog


Using In Go Code
----------------


	import "github.com/mabetle/golog"

	logger:=golog.GetLogger("")

	//logger:=golog.GetSeeLogger("")

	//logger:=golog.GetGoLangLogger("")
		
	//logger:=golog.GetSimpleLogger("")

	//logger:=golog.GetTimberLogger("")

	logger.Info(args ...interface{})

	logger.Error(args ...interface{})

	logger.Warn(args ...interface{})

	logger.Debug(args ...interface{})

	logger.Trace(args ...interface{})


Run Demo
-------

	cd MABETLE_GOLOG_PATH

	go run cmd/demo/main.go


BUGS
----
Not implement Catalog and Level 

TODO
----
Any ideas and suggestions are Welcome.


