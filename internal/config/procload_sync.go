package config

import "github.com/ZXSQ1/rviv/env"

func LoadSyncOpts(optsRaw map[string]any) any {
	src, _ := NewPath(optsRaw["src"].(string))
	dest, _ := NewPath(optsRaw["dest"].(string))
	oneway := false
	method := "add"

	if optsRaw["oneway"] != nil {
		oneway = optsRaw["oneway"].(bool)
	}

	if optsRaw["method"] != nil {
		method = optsRaw["method"].(string)
	}

	return SyncOpts{
		Src:     src,
		Dest:    dest,
		Oneway:  oneway,
		Method:  method,
		Verbose: env.Verbose,
	}
}
