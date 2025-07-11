package config

type OrganizeOpts struct {
	Srcs        []Path
	OrganizeDir Path
	Method      string
	Datefmt     string
	Verbose     bool
}
