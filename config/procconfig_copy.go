package config

type CopyOpts struct {
	Srcs      []Path
	Dest      Path
	Method    string
	Necessary bool
}
