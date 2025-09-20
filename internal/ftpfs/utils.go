package ftpfs

import (
	"fmt"
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

func stderr(err error) error {
	if err == nil {
		return nil
	}

	errorParts := strings.Split(err.Error(), ":")
	actualError := errorParts[len(errorParts)-1]
	actualError = strings.TrimSpace(actualError)

	return fmt.Errorf("%s", actualError)
}
