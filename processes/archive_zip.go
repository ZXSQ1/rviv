package processes

import (
	"archive/zip"
	"io"

	"github.com/ZXSQ1/rviv/filesystem"
)

func ArchiveZip(archive *Path, entries *Paths, ignoreBadFiles bool) error {
	archiveObj, err := archive.Filesys.Open(
		archive.Filename, filesystem.ModeWrite,
	)

	if err != nil {
		return err
	}

	defer archiveObj.Close()
	zipWriter := zip.NewWriter(archiveObj)
	defer zipWriter.Close()

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

		header, err := zip.FileInfoHeader(info)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		header.Name = entry
		header.Method = zip.Deflate
		header.Flags = 0x8

		writer, err := zipWriter.CreateHeader(header)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		_, err = io.Copy(writer, entryWriter)

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
