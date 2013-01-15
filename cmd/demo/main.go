// run a golog demo
package main

import(
	"github.com/mabetle/golog"
)

func main() {

	logger := golog.GetLogger("")


	logger.Error("Hello")
	logger.Warn("Hello")
	logger.Info("Hello")
	logger.Debug("Hello")
	logger.Trace("Hello")

}

