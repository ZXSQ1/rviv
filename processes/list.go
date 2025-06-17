package processes

import "github.com/ZXSQ1/rviv/logging"

func ListDir(src *Path, recursive, ignoreBadFiles bool) (*Paths, error) {
	if !recursive {
		entries, err := src.Filesys.ListDir(src.Filename)
		logging.ReportErr(err)

		return &Paths{
			Filenames: entries,
			Filesys:   src.Filesys,
		}, err
	}

	entries, err := ListDir(src, false, ignoreBadFiles)
	logging.ReportErr(err)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries.Filenames {
		entrystat, err := src.Filesys.Stat(entry)
		logging.ReportErr(err)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return nil, err
		}

		if !entrystat.IsDir() {
			entries.Filenames = append(entries.Filenames, entry)
		} else {
			newEntries, err := ListDir(
				&Path{
					Filename: entry,
					Filesys:  src.Filesys,
				}, true, ignoreBadFiles,
			)

			logging.ReportErr(err)

			if err != nil {
				if ignoreBadFiles {
					continue
				}

				return nil, err
			}

			entries.Filenames = append(
				entries.Filenames, newEntries.Filenames...,
			)
		}
	}

	return entries, nil
}
