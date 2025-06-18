package localfs

import (
	"fmt"
	"os"
)

func (local *LocalFs) Remove(filename string) error {
	stat, err := os.Stat(filename)

	if err != nil {
		return os.ErrNotExist
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("file is not a regular file")
	}

	return os.Remove(filename)
}

func (local *LocalFs) RemoveDir(filename string) error {
	stat, err := os.Stat(filename)

	if err != nil {
		return os.ErrNotExist
	}

	if !stat.IsDir() {
		return fmt.Errorf("file is not a regular file")
	}

	return os.RemoveAll(filename)
}
