package processes

import (
	"archive/tar"
	"io"
)

// continues when paths have permission errors or the like
func ArchiveTar(archive *Path, entries *Paths) error {
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
			continue
		}

		info, err := entries.Filesys.Stat(entry)

		if err != nil {
			continue
		}

		header, err := tar.FileInfoHeader(info, info.Name())

		if err != nil {
			continue
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			continue
		}

		_, err = io.Copy(tarWriter, entryWriter)

		if err != nil {
			continue
		}

		entryWriter.Close()
	}

	return nil
}
