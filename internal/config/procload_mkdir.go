package config

import "github.com/ZXSQ1/rviv/internal/env"

func LoadMkdirOpts(rawOpts map[string]any) any {
	files := []Path{}
	parent := false

	for _, filename := range rawOpts["paths"].([]any) {
		fileObj, _ := NewPath(filename.(string))
		files = append(files, fileObj)
	}

	if rawOpts["parent"] != nil {
		parent = true
	}

	return MkdirOpts{
		Filenames: files,
		Parent:    parent,
		Verbose:   env.Verbose,
	}
}
