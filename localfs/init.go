package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

type LocalFs struct {
	prefix  string
	currdir string
}

func Init(prefix string) (filesystem.Filesystem, error) {
	currdir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	if err := os.Chdir(prefix); err != nil {
		return nil, err
	}

	return &LocalFs{
		prefix:  prefix,
		currdir: currdir,
	}, nil
}
