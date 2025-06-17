package webdavfs

import (
	"fmt"
	"io"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

type File struct {
	mode   filesystem.OpenMode
	done   chan error
	writer io.WriteCloser
	reader io.ReadCloser
}

func (file *File) Read(p []byte) (int, error) {
	if file.mode != filesystem.ModeRead {
		return -1, os.ErrInvalid
	}

	if file.reader == nil {
		return -1, fmt.Errorf("reader not initialized")
	}

	return file.reader.Read(p)
}

func (file *File) Write(p []byte) (int, error) {
	if file.mode != filesystem.ModeWrite {
		return -1, os.ErrInvalid
	}

	if file.writer == nil {
		return -1, fmt.Errorf("writer not initialized")
	}

	if file.done == nil {
		return -1, fmt.Errorf("error channel not initialized")
	}

	return file.writer.Write(p)
}

func (file *File) Close() error {
	if file.mode == filesystem.ModeWrite {
		err := <-file.done

		if err != nil {
			return err
		}

		close(file.done)
	}

	file.done = nil
	file.mode = filesystem.OpenMode(2)
	file.reader = nil
	file.writer = nil

	return nil
}
