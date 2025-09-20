package webdavfs

import (
	"strconv"
	"strings"

	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/ZXSQ1/rviv/internal/logging"
)

func status(err error) int {
	if err == nil {
		return -1
	}

	parts := strings.Split(err.Error(), " ")
	code, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		logging.ReportErr(err)
	}

	return code
}

func stderr(err error) error {
	switch status(err) {
	case 404:
		return filesystem.ErrNotExist
	default:
		return err
	}
}
