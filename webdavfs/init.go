package webdavfs

import (
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/studio-b12/gowebdav"
)

type WebDavFs struct {
	conn     *gowebdav.Client
	connInfo *filesystem.ConnInfo
}

func Connect(connInfo *filesystem.ConnInfo) (*WebDavFs, error) {
	conn := gowebdav.NewClient(
		"http://"+connInfo.Addr, connInfo.User, connInfo.Pass,
	)

	err := conn.Connect()

	if err != nil {
		return nil, err
	}

	return &WebDavFs{
		conn:     conn,
		connInfo: connInfo,
	}, nil
}
