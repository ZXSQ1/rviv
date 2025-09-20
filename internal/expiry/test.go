package expiry

import (
	"fmt"
	"os"
)

func println(objs ...any) {
	fmt.Println(objs...)
}

func exit() {
	os.Exit(1)
}
