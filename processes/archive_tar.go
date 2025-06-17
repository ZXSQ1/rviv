package processes

import (
	"archive/tar"
	"io"

	"github.com/ZXSQ1/rviv/filesystem"
)

func ArchiveTar(archive *Path, entries *Paths, ignoreBadFiles bool) error {
	archiveObj, err := archive.Filesys.Open(
		archive.Filename, filesystem.ModeWrite,
	)

	if err != nil {
		return err
	}

	defer archiveObj.Close()
	tarWriter := tar.NewWriter(archiveObj)
	defer tarWriter.Close()

	for _, entry := range entries.Filenames {
		entryWriter, err := entries.Filesys.Open(entry, filesystem.ModeRead)

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

		err = tarWriter.WriteHeader(header)

		if err != nil {
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

		err = entryWriter.Close()

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}
	}

	return nil
}
