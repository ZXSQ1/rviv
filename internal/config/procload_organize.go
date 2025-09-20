package config

import "github.com/ZXSQ1/rviv/internal/env"

func LoadOrganizeOpts(optsRaw map[string]any) any {
	srcs := []Path{}
	organizedir, _ := NewPath(optsRaw["organizedir"].(string))
	method := ""
	datefmt := ""

	for _, srcRaw := range optsRaw["srcs"].([]any) {
		src, _ := NewPath(srcRaw.(string))
		srcs = append(srcs, src)
	}

	if optsRaw["method"] != nil {
		method = optsRaw["method"].(string)
	}

	if optsRaw["date"] != nil {
		datefmt = optsRaw["date"].(string)
	}

	return OrganizeOpts{
		Srcs:        srcs,
		OrganizeDir: organizedir,
		Method:      method,
		Datefmt:     datefmt,
		Verbose:     env.Verbose,
	}
}
