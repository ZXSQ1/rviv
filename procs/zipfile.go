package procs

import (
	"archive/zip"
	"compress/flate"
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/info"
)

func ZipFile(archivename, archivefmt config.Path, entries []config.Path,
	expiryInDays, level int, safe, verbose bool) error {

	if err := CheckExistsCreate(archivename); err != nil {
		return err
	} else if err := CheckIsRegular(archivename); err != nil {
		return err
	}

	outObj, err := archivename.Fsys.Open(
		archivename.Filename, filesystem.ModeWrite,
	)

	if err != nil {
		return err
	}

	defer outObj.Close()

	resultEntries := []config.Path{}
	zipObj := zip.NewWriter(outObj)
	defer zipObj.Close()

	zipObj.RegisterCompressor(
		zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
			return flate.NewWriter(w, level)
		},
	)

	for _, entry := range entries {
		if err := CheckExists(entry); err != nil {
			return err
		} else if err := CheckIsRegular(entry); err != nil {
			entryEntries, err := ListDir(config.ListOpts{
				Filename:  entry,
				Recursive: true,
				Verbose:   verbose,
			})

			if err != nil {
				resultEntries = append(entries, entryEntries...)
			}
		} else {
			resultEntries = append(resultEntries, entry)
		}
	}

	for _, entry := range entries {
		zipEntry, err := CheckOpen(entry, filesystem.ModeRead)

		if err != nil {
			return err
		}

		defer zipEntry.Close()

		stat, err := CheckStat(entry)

		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(stat)

		if err != nil {
			return info.Error(
				"unable to get zip header for file '%s'", ShowPath(entry),
			)
		}

		header.Method = zip.Deflate
		writer, err := zipObj.CreateHeader(header)

		if err != nil {
			return info.Error(
				"unable to create header for zip entry '%s'", ShowPath(entry),
			)
		}

		_, err = io.Copy(writer, zipEntry)

		if err != nil {
			return err
		}
	}

	return nil
}
