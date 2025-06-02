package ftpfs

import (
	"strconv"
	"strings"
)

func status(err error) int {
	if err == nil {
		return -1
	}

	statusString := strings.Split(err.Error(), " ")[0]
	status, _ := strconv.Atoi(statusString)

	return status
}
