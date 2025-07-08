package config

type SyncOpts struct {
	Src     Path
	Dest    Path
	Oneway  bool
	Method  string
	Verbose bool
}
