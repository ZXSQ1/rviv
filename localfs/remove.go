package localfs

import "os"

func (local *LocalFS) Remove(filename string) error {
	return os.Remove(filename)
}

func (local *LocalFS) RemoveDir(filename string) error {
	return os.RemoveAll(filename)
}
