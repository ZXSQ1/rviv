package webdavfs

import (
	"net/http"

	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/studio-b12/gowebdav"
)

type WebDavFs struct {
	httpClient *http.Client
	conn       *gowebdav.Client
	connInfo   *filesystem.ConnInfo
	closed     bool
}

func Connect(connInfo *filesystem.ConnInfo) (filesystem.Filesystem, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
		},
	}

	conn := gowebdav.NewClient(
		"http://"+connInfo.Addr, connInfo.User, connInfo.Pass,
	)

	conn.SetTimeout(connInfo.Timeout)
	conn.SetTransport(httpClient.Transport)
	err := conn.Connect()

	if err != nil {
		return nil, err
	}

	return &WebDavFs{
		httpClient: httpClient,
		conn:       conn,
		connInfo:   connInfo,
		closed:     false,
	}, nil
}
