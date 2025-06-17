package logging

import (
	"log"
	"runtime"
)

func ReportErr(err error) {
	if err == nil {
		return
	}

	if !Debug {
		return
	}

	callerName := "unknown caller"
	pc, _, _, ok := runtime.Caller(1)

	if ok {
		fn := runtime.FuncForPC(pc)
		callerName = fn.Name()
	}

	log.Printf(
		"%s: %s\n", callerName, err.Error(),
	)
}
