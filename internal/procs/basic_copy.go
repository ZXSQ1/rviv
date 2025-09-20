package procs

import "github.com/ZXSQ1/rviv/internal/config"

func CopyPath(basePath config.Path, newFilename string) config.Path {
	return config.Path{
		Filename: newFilename,
		Devname:  basePath.Devname,
		Devices:  basePath.Devices,
		Active:   basePath.Active,
		Fsys:     basePath.Fsys,
	}
}
