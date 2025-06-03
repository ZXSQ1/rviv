package localfs

import "os"

func (local *LocalFS) Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}
