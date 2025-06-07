package main

import (
	"fmt"

	"github.com/ZXSQ1/rviv/ftpfs"
)

func main() {
	const filename = "A1-Omar"

	ftpfilesys, _ := ftpfs.Connect("192.168.1.7", 2121, "omar", "dlpob")
	defer ftpfilesys.Close()
}
