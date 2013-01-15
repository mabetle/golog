package golog

import(
	"testing"
)

func TestLog(t *testing.T) {
	// default getLogger()
	DemoLogger(GetLogger(""))
	DemoLogger(GetSimpleLogger(""))
	DemoLogger(GetSeeLogger(""))
	DemoLogger(GetTimberLogger(""))
}

