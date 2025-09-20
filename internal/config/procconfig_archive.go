package config

import "time"

type ArchiveOpts struct {
	Archivename string
	Archivefmt  string
	Parents     []Path
	Entries     []Path
	ExpiryTime  *time.Time
	Compression string
	Level       int
	Verbose     bool
}
