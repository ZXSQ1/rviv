package localfs

import "os"

func (local *LocalFS) IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
