package config

import (
	"time"

	"github.com/ZXSQ1/rviv/env"
	"github.com/ZXSQ1/rviv/expiry"
)

func LoadArchiveOpts(optsRaw map[string]any) any {
	archivename, _ := NewPath(optsRaw["archivename"].(string))
	entries := []Path{}
	compression := ""
	level := 6

	var expiryTime *time.Time = nil

	for _, entry := range optsRaw["entries"].([]any) {
		entry, _ := NewPath(entry.(string))
		entries = append(entries, entry)
	}

	if optsRaw["expiry"] != nil {
		expiryTimeVal, _ := expiry.ParseExpiry(
			optsRaw["expiry"].(string),
		)

		expiryTime = &expiryTimeVal
	}

	if optsRaw["compression"] != nil {
		compression = optsRaw["compression"].(string)
	}

	if optsRaw["level"] != nil {
		level = int(optsRaw["level"].(float64))
	}

	return ArchiveOpts{
		Archivename: archivename,
		Archivefmt: Path{
			Filename: optsRaw["archivename"].(string),
			Devname:  archivename.Devname,
			Devices:  archivename.Devices,
			Active:   archivename.Active,
			Fsys:     archivename.Fsys,
		},

		Entries:     entries,
		Compression: compression,
		Level:       level,
		ExpiryTime:  expiryTime,
		Verbose:     env.Verbose,
	}
}
