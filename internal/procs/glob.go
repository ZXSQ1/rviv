package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/info"
)

func Glob(opts config.GlobOpts) ([]config.Path, error) {
	results := []config.Path{}

	for _, entry := range opts.Entries {
		if !strings.Contains(entry.Filename, "*") {
			results = append(results, entry)
			continue
		}

		if !entry.Active {
			if err := entry.Connect(opts.Verbose); err != nil {
				return nil, err
			}
		}

		pathPrefix := strings.Split(entry.Filename, "*")[0]

		if !strings.HasSuffix(pathPrefix, "/") {
			pathPrefix, _ = filepath.Split(pathPrefix)
		}

		pathPrefix = strings.TrimRight(pathPrefix, "/")
		prefixEntries, err := ListDir(config.ListOpts{
			Filename: config.Path{
				Filename: pathPrefix,
				Devname:  entry.Devname,
				Active:   true,
				Fsys:     entry.Fsys,
				Devices:  entry.Devices,
			},

			Recursive: false,
		})

		if err != nil {
			return nil, err
		}

		info.Text(opts.Verbose, "globbing entry '%s'", ShowPath(entry))

		for _, prefixEntry := range prefixEntries {
			ok, err := filepath.Match(entry.Filename, prefixEntry.Filename)

			if err != nil {
				return nil, info.Error(
					"unable to match path '%s' to glob path", prefixEntry.Filename,
				)
			}

			if ok {
				results = append(results, prefixEntry)
			}
		}
	}

	return results, nil
}
