package config

type ZipOpts struct {
	Archivename Path
	Entries     []Path
	Safe        bool
	Necessary   bool
}
