package logging

import (
	"fmt"
	"log"
	"runtime"
)

func Log(logmsg string) {
	if !Debug {
		return
	}

	callerName := "unknown caller"
	pc, _, _, ok := runtime.Caller(1)

	if ok {
		fn := runtime.FuncForPC(pc)
		callerName = fn.Name()
	}

	log.Printf("%s: %s", callerName, logmsg)
}

func Logln(logmsg string) {
	if !Debug {
		return
	}

	callerName := "unknown caller"
	pc, _, _, ok := runtime.Caller(1)

	if ok {
		fn := runtime.FuncForPC(pc)
		callerName = fn.Name()
	}

	log.Printf("%s: %s\n", callerName, logmsg)
}

func Logf(logformat string, v ...any) {
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
		"%s: %s", callerName,
		fmt.Sprintf(logformat, v...),
	)
}
