package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/spf13/afero"
)

type LocalFs struct {
	fsys afero.Fs
}

func Init(prefix string) (filesystem.Filesystem, error) {
	info, err := os.Stat(prefix)

	if err != nil {
		return nil, filesystem.ErrExist
	}

	if !info.IsDir() {
		return nil, filesystem.ErrFileNotDir
	}

	fsys := afero.NewBasePathFs(
		afero.NewOsFs(), prefix,
	)

	return &LocalFs{
		fsys: fsys,
	}, nil
}
