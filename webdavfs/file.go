package webdavfs

import (
	"fmt"
	"io"

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
		if file.mode == filesystem.ModeClosed {
			return -1, filesystem.ErrClosed
		}

		return -1, filesystem.ErrModeRead
	}

	if file.reader == nil {
		return -1, fmt.Errorf("reader not initialized")
	}

	return file.reader.Read(p)
}

func (file *File) Write(p []byte) (int, error) {
	if file.mode != filesystem.ModeWrite {
		if file.mode == filesystem.ModeClosed {
			return -1, filesystem.ErrClosed
		}

		return -1, filesystem.ErrModeWrite
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
	if file.mode == filesystem.ModeClosed {
		return filesystem.ErrClosed
	}

	if file.mode == filesystem.ModeWrite {
		if file.writer != nil {
			if err := file.writer.Close(); err != nil {
				return err
			}
		}

		if file.done != nil {
			if err := <-file.done; err != nil {
				return err
			}
		}
	}

	if file.mode == filesystem.ModeRead && file.reader != nil {
		if err := file.reader.Close(); err != nil {
			return err
		}
	}

	file.done = nil
	file.mode = filesystem.ModeClosed
	file.reader = nil
	file.writer = nil

	return nil
}
