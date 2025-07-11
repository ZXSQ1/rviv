package config

type ArchiveOpts struct {
	Archivename  Path
	Archivefmt   Path
	Entries      []Path
	ExpiryInDays int
	Compression  string
	Level        int
	Safe         bool
	Verbose      bool
}
