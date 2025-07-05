package config

type ArchiveOpts struct {
	Archivename  Path
	Entries      []Path
	ExpiryInDays float64
	Compression  string
	Safe         bool
	Necessary    bool
}
