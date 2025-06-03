package processes

import (
	"archive/tar"
	"io"
)

func ArchiveTar(archive *Path, entries *Paths, ignoreBadFiles bool) error {
	archiveObj, err := archive.Filesys.Open(archive.Filename)

	if err != nil {
		return err
	}

	defer archiveObj.Close()
	tarWriter := tar.NewWriter(archiveObj)
	defer tarWriter.Close()

	for _, entry := range entries.Filenames {
		entryWriter, err := entries.Filesys.Open(entry)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		info, err := entries.Filesys.Stat(entry)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		_, err = io.Copy(tarWriter, entryWriter)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		entryWriter.Close()
	}

	return nil
}
