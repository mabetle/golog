Mabetle Go Log
================

Usage:
------

go get -u github.com/mabetle/golog

import "github.com/mabetle/golog"

logger:=golog.GetLogger("")


logger.Info(args ...interface{})

logger.Error(args ...interface{})

logger.Warn(args ...interface{})

logger.Debug(args ...interface{})

logger.Trace(args ...interface{})


