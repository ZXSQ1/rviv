package config

import "github.com/ZXSQ1/rviv/internal/env"

func LoadCopyOpts(optsRaw map[string]any) any {
	srcs := []Path{}
	dest, _ := NewPath(optsRaw["dest"].(string))

	for _, srcRaw := range optsRaw["srcs"].([]any) {
		src, _ := NewPath(srcRaw.(string))
		srcs = append(srcs, src)
	}

	return CopyOpts{
		Srcs:    srcs,
		Dest:    dest,
		Method:  optsRaw["method"].(string),
		Verbose: env.Verbose,
	}
}
