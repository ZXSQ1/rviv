package config

import "github.com/ZXSQ1/rviv/env"

func LoadRemoveOpts(optsRaw map[string]any) any {
	files := []Path{}
	recursive := false

	for _, filename := range optsRaw["paths"].([]any) {
		fileObj, _ := NewPath(filename.(string))
		files = append(files, fileObj)
	}

	if optsRaw["recursive"] != nil {
		recursive = optsRaw["recursive"].(bool)
	}

	return RemoveOpts{
		Filenames: files,
		Recursive: recursive,
		Verbose:   env.Verbose,
	}
}
