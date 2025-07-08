package info

import (
	"fmt"
	"os"

	"github.com/ZXSQ1/rviv/env"
	"github.com/fatih/color"
)

func Fatal(format string, objs ...any) {
	err := Error(format, objs...)

	if !env.Colored {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		errorPrefix := color.New(color.Bold, color.FgRed).Sprint("error:")
		fmt.Printf("%s %s\n", errorPrefix, err.Error())
	}

	os.Exit(1)
}
