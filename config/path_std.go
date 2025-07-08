package config

import (
	"os"
	"strings"
	"time"

	"github.com/itchyny/timefmt-go"
)

// replaces the '%x' format with the format like unix date (from the timefmt
// library); to stop the replacement, you can do '%%' instead of '%'; expands
// environment variables
func StdPath(filename string) string {
	filename = os.ExpandEnv(filename)
	filename = strings.ReplaceAll(filename, "%%", "%\\")
	filename = timefmt.Format(time.Now(), filename)
	filename = strings.ReplaceAll(filename, "%\\", "%")

	return filename
}
