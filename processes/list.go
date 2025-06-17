package processes

func ListDir(src *Path, recursive, ignoreBadFiles bool) (*Paths, error) {
	if !recursive {
		entries, err := src.Filesys.ListDir(src.Filename)

		return &Paths{
			Filenames: entries,
			Filesys:   src.Filesys,
		}, err
	}

	entries, err := ListDir(src, false, ignoreBadFiles)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries.Filenames {
		entrystat, err := src.Filesys.Stat(entry)

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
