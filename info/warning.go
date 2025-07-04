package info

import (
	"fmt"

	"github.com/ZXSQ1/rviv/env"
	"github.com/fatih/color"
)

func Warning(format string, objs ...any) {
	if !env.Colored {
		fmt.Printf(
			"warning: %s\n", fmt.Sprintf(format, objs...),
		)
	} else {
		warningPrefix := color.New(color.Bold, color.FgYellow).Sprint("W:")

		fmt.Printf(
			"%s %s\n", warningPrefix, fmt.Sprintf(
				format, objs...,
			),
		)
	}
}
