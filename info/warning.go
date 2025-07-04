package info

import (
	"fmt"

	"github.com/ZXSQ1/rviv/env"
	"github.com/fatih/color"
)

func Warning(format string, objs ...any) error {
	err := fmt.Errorf(format, objs...)

	if !env.Colored {
		fmt.Printf("warning: %s\n", err.Error())
	} else {
		warningPrefix := color.New(color.Bold, color.FgYellow).Sprint("warning:")
		fmt.Printf("%s %s\n", warningPrefix, err.Error())
	}

	return err
}
