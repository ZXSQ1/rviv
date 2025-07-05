package config

import (
	"os"
	"time"

	"github.com/itchyny/timefmt-go"
)

// replaces the '%x' format with the format like unix date (from the timefmt
// library); expands environment variables
func StdPath(filename string) string {
	filename = os.ExpandEnv(filename)
	filename = timefmt.Format(time.Now(), filename)

	return filename
}
