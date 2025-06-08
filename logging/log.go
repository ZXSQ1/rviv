package logging

import (
	"fmt"
	"log"
	"runtime"
)

// logs the message (e.g. error messages from the direct caller function) to the
// standard output specifying the direct caller of the Log function
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

// does the same as Log, but adds a new line
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

// simply a formatted Log
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
