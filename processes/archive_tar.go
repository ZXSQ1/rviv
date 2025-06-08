package processes

import (
	"archive/tar"
	"io"

	"github.com/ZXSQ1/rviv/logging"
)

// makes a tar archive given a filesystem-independent path and some entries; has
// an option to ignore erroneous entries; not completed; has not been tested
func ArchiveTar(archive *Path, entries *Paths, ignoreBadFiles bool) error {
	archiveObj, err := archive.Filesys.Open(archive.Filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	defer archiveObj.Close()
	tarWriter := tar.NewWriter(archiveObj)
	defer tarWriter.Close()

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

		header, err := tar.FileInfoHeader(info, info.Name())
		logging.ReportErr(err)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		err = tarWriter.WriteHeader(header)
		logging.ReportErr(err)

		if err != nil {
			if ignoreBadFiles {
				continue
			}

			return err
		}

		_, err = io.Copy(tarWriter, entryWriter)
		logging.ReportErr(err)

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
