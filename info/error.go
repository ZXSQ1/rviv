package info

import (
	"fmt"
	"os"

	"github.com/ZXSQ1/rviv/env"
	"github.com/fatih/color"
)

func Error(format string, objs ...any) {
	if !env.Colored {
		fmt.Printf(
			"error: %s\n", fmt.Sprintf(format, objs...),
		)
	} else {
		errorPrefix := color.New(color.Bold, color.FgRed).Sprint("E:")

		fmt.Printf(
			"%s %s\n", errorPrefix, fmt.Sprintf(
				format, objs...,
			),
		)
	}

	os.Exit(1)
}
