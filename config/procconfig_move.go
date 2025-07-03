package config

type MoveOpts struct {
	Srcs      []Path
	Dest      Path
	Method    string
	Necessary bool
}
