package config

import "time"

type ArchiveOpts struct {
	Archivename Path
	Archivefmt  Path
	Entries     []Path
	ExpiryTime  *time.Time
	Compression string
	Level       int
	Verbose     bool
}
