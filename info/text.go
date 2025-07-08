package info

import "fmt"

var Padding = "  "

func Text(verbose bool, format string, objs ...any) {
	if !verbose {
		return
	}

	fmt.Printf(
		Padding+format, objs...,
	)
}
