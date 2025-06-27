package localfs

import "github.com/ZXSQ1/rviv/filesystem"

type LocalFs struct {
	prefix string
}

func Init(prefix string) filesystem.Filesystem {
	return &LocalFs{prefix: prefix}
}
