package info

import (
	"fmt"
)

func Error(format string, objs ...any) error {
	return fmt.Errorf(format, objs...)
}
