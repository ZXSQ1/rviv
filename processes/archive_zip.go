package processes

import (
	"archive/zip"
	"io"
)

// continues when entries have permission errors or the like
func ArchiveZip(archive *Path, entries *Paths) error {
	archiveObj, err := archive.Filesys.Open(archive.Filename)

	if err != nil {
		return err
	}

	defer archiveObj.Close()
	zipWriter := zip.NewWriter(archiveObj)
	defer zipWriter.Close()

	for _, entry := range entries.Filenames {
		entryWriter, err := entries.Filesys.Open(entry)

		if err != nil {
			continue
		}

		info, err := entries.Filesys.Stat(entry)

		if err != nil {
			continue
		}

		header, err := zip.FileInfoHeader(info)

		if err != nil {
			continue
		}

		header.Name = entry
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)

		if err != nil {
			continue
		}

		_, err = io.Copy(writer, entryWriter)
		entryWriter.Close()

		if err != nil {
			continue
		}
	}

	return nil
}
