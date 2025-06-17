package localfs

import (
	"os"
)

func (local *LocalFs) Remove(filename string) error {
	return os.Remove(filename)
}

func (local *LocalFs) RemoveDir(filename string) error {
	return os.RemoveAll(filename)
}
