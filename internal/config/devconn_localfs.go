package config

import (
	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/ZXSQ1/rviv/internal/info"
	"github.com/ZXSQ1/rviv/internal/localfs"
)

func ConnectLocalFs(infoRaw any, verbose bool) (filesystem.Filesystem, error) {
	localInfo, ok := infoRaw.(LocalFsConfig)

	if !ok {
		return nil, info.Error("unable to load local filesystem info")
	}

	client, err := localfs.Init(localInfo.Prefix)

	if err != nil {
		return nil, info.Error("unable to connect to local filesystem")
	}

	return client, nil
}
