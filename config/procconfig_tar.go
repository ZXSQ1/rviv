package config

type TarOpts struct {
	Archivename Path
	Entries     []Path
	Compression string
	Safe        bool
	Necessary   bool
}
