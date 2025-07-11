package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/config"
)

type Differences struct {
	UniqueSrcEntries  []config.Path
	UniqueDestEntries []config.Path
	SrcCommonEntries  []config.Path
}

func Differ(opts config.DifferOpts) (Differences, error) {
	diffs := Differences{}

	if err := AssertExists(opts.Src); err != nil {
		return diffs, err
	}

	if err := AssertIsDir(opts.Src); err != nil {
		return diffs, err
	}

	if err := AssertExists(opts.Dest); err != nil {
		return diffs, err
	}

	if err := AssertIsDir(opts.Dest); err != nil {
		return diffs, err
	}

	srcEntries, err := ListDir(config.ListOpts{
		Filename:  opts.Src,
		Recursive: true,
		Verbose:   opts.Verbose,
	})

	if err != nil {
		return Differences{}, err
	}

	destEntries, err := ListDir(config.ListOpts{
		Filename:  opts.Dest,
		Recursive: true,
		Verbose:   opts.Verbose,
	})

	if err != nil {
		return Differences{}, err
	}

	const (
		BOTH     = 0
		SRCONLY  = 1
		DESTONLY = 2
	)

	frequencyMap := map[string]int{}

	for _, fullSrcEntry := range srcEntries {
		srcEntry := strings.TrimLeft(strings.Replace(
			fullSrcEntry.Filename, opts.Src.Filename, "", 1,
		), "/")

		frequencyMap[srcEntry] = SRCONLY
	}

	for _, fullDestEntry := range destEntries {
		destEntry := strings.TrimLeft(strings.Replace(
			fullDestEntry.Filename, opts.Dest.Filename, "", 1,
		), "/")

		freq, ok := frequencyMap[destEntry]

		if !ok {
			frequencyMap[destEntry] = DESTONLY
		} else if freq == SRCONLY {
			frequencyMap[destEntry] = BOTH
		}
	}

	for entry, freq := range frequencyMap {
		if freq == SRCONLY {
			diffs.UniqueSrcEntries = append(diffs.UniqueSrcEntries, config.Path{
				Filename: filepath.Join(opts.Src.Filename, entry),
				Devname:  opts.Src.Devname,
				Devices:  opts.Src.Devices,
				Active:   opts.Src.Active,
				Fsys:     opts.Src.Fsys,
			})
		} else if freq == DESTONLY {
			diffs.UniqueDestEntries = append(diffs.UniqueDestEntries, config.Path{
				Filename: filepath.Join(opts.Dest.Filename, entry),
				Devname:  opts.Dest.Devname,
				Devices:  opts.Dest.Devices,
				Active:   opts.Dest.Active,
				Fsys:     opts.Dest.Fsys,
			})
		} else {
			diffs.SrcCommonEntries = append(diffs.SrcCommonEntries, config.Path{
				Filename: filepath.Join(opts.Src.Filename, entry),
				Devname:  opts.Src.Devname,
				Devices:  opts.Src.Devices,
				Active:   opts.Src.Active,
				Fsys:     opts.Src.Fsys,
			})
		}
	}

	return diffs, nil
}
