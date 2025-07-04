package info

import (
	"fmt"
	"os"

	"github.com/ZXSQ1/rviv/env"
	"github.com/fatih/color"
)

var ExitOnError = true

func Error(format string, objs ...any) error {
	err := fmt.Errorf(format, objs...)

	if !env.Colored {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		errorPrefix := color.New(color.Bold, color.FgRed).Sprint("error:")
		fmt.Printf("%s %s\n", errorPrefix, err.Error())
	}

	if ExitOnError {
		os.Exit(1)
	}

	return err
}
