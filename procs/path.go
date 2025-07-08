package procs

import (
	"strings"

	"github.com/ZXSQ1/rviv/config"
)

func ShortenPath(filename string) string {
	processedFilename := ""
	pathParts := strings.Split(filename, "/")

	for _, filenamePart := range pathParts[:len(pathParts)-1] {
		processedFilename += "/" + string(filenamePart[0])
	}

	return processedFilename + "/" + pathParts[len(pathParts)-1]
}

func ShowPath(pathObj config.Path) string {
	return pathObj.Devname + config.DeviceSeparator + ShortenPath(
		pathObj.Filename)
}
