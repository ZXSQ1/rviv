package main

import (
	"fmt"

	"github.com/ZXSQ1/rviv/config"
)

func main() {
	config.LoadConfig("./test_config.json")
	config.LoadDevices()
	config.LoadProcesses()

	fmt.Println(config.Devices)
	fmt.Println(config.Processes)
}
