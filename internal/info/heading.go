package info

import (
	"fmt"

	"github.com/ZXSQ1/rviv/internal/env"
	"github.com/fatih/color"
)

func Heading(verbose bool, format string, objs ...any) {
	if !verbose {
		return
	}

	prompt := fmt.Sprintf(format+":\n", objs...)

	if env.Colored {
		prompt = color.New(color.Bold, color.FgGreen).Sprint(prompt)
	}

	fmt.Println(prompt)
}
