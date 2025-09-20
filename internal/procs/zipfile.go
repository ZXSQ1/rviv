package procs

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/info"
	"github.com/ZXSQ1/rviv/localfs"
)

func ZipFile(archiveBasename string, parents, entries []config.Path, level int,
	verbose bool) error {

	localFs, err := localfs.Init(os.TempDir())

	if err != nil {
		return fmt.Errorf("unable to create local client for temporary files")
	}

	tempArchive := config.Path{
		Filename: archiveBasename,
		Active:   true,
		Fsys:     localFs,
	}

	if err = IsRegular(tempArchive); err != nil {
		return err
	}

	zipFile, err := Open(tempArchive, filesystem.ModeWrite)

	if err != nil {
		return err
	}

	defer zipFile.Close()
	zipObj := zip.NewWriter(zipFile)
	defer zipObj.Close()

	zipObj.RegisterCompressor(
		zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
			return flate.NewWriter(w, level)
		},
	)

	for _, entry := range entries {
		if err := IsExist(entry, nil); err != nil {
			return err
		}

		zipEntry, err := Open(entry, filesystem.ModeRead)

		if err != nil {
			return err
		}

		defer zipEntry.Close()

		stat, err := Stat(entry)

		if err != nil {
			return err
		}

		if stat.IsDir() {
			header := &zip.FileHeader{
				Name:   filepath.Clean(entry.Filename) + "/",
				Method: zip.Store,
			}

			header.SetMode(filesystem.PermDir)
			_, err := zipObj.CreateHeader(header)

			if err != nil {
				return info.Error(
					"unable to create header for zip entry '%s'", ShowPath(entry),
				)
			}

			continue
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

	for _, parent := range parents {
		if err := IsDir(parent); err != nil {
			return err
		}

		err := CopyFiles([]config.Path{tempArchive}, parent, verbose)

		if err != nil {
			return err
		}
	}

	return nil
}
