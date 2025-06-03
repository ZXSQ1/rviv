package localfs

import "os"

func (local *LocalFS) Open(filename string) (*os.File, error) {
	return os.Open(filename)
}
