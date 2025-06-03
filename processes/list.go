package processes

// continues when paths have permission errors or the like
func ListDir(src *Path, recursive bool) (*Paths, error) {
	if !recursive {
		entries, err := src.Filesys.ListDir(src.Filename)

		return &Paths{
			Filenames: entries,
			Filesys:   src.Filesys,
		}, err
	}

	entries, err := ListDir(src, false)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries.Filenames {
		entrystat, err := src.Filesys.Stat(entry)

		if err != nil {
			continue
		}

		if !entrystat.IsDir() {
			entries.Filenames = append(entries.Filenames, entry)
		} else {
			newEntries, err := ListDir(
				&Path{
					Filename: entry,
					Filesys:  src.Filesys,
				}, true,
			)

			if err != nil {
				continue
			}

			entries.Filenames = append(
				entries.Filenames, newEntries.Filenames...,
			)
		}
	}

	return entries, nil
}
