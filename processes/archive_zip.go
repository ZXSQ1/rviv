package processes

import (
	"archive/zip"
	"io"

	"github.com/ZXSQ1/rviv/logging"
)

// archives a zip archive given a path for the archive and a set of entries; has
// an option to ignore erroneous file; not complete; has not been tested
func ArchiveZip(archive *Path, entries *Paths, ignoreBadFiles bool) error {
	archiveObj, err := archive.Filesys.Open(archive.Filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	defer archiveObj.Close()
	zipWriter := zip.NewWriter(archiveObj)
	defer zipWriter.Close()

	for _, entry := range entries.Filenames {
		entryWriter, err := entries.Filesys.Open(entry)
		logging.ReportErr(err)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		info, err := entries.Filesys.Stat(entry)
		logging.ReportErr(err)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		header, err := zip.FileInfoHeader(info)
		logging.ReportErr(err)

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
		logging.ReportErr(err)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		_, err = io.Copy(writer, entryWriter)
		logging.ReportErr(err)

		err = entryWriter.Close()
		logging.ReportErr(err)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}
	}

	return nil
}
