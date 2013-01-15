// run a golog demo
package main

import(
	"github.com/mabetle/golog"
)

func main() {

	golog.DemoLogger(golog.GetLogger(""))
	golog.DemoLogger(golog.GetSimpleLogger(""))
	golog.DemoLogger(golog.GetSeeLogger(""))
	golog.DemoLogger(golog.GetTimberLogger(""))

}

