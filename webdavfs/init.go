package webdavfs

import (
	"fmt"

	"github.com/studio-b12/gowebdav"
)

type WebDavFs struct {
	conn *gowebdav.Client
}

func Connect(addr, user, pass string) (*WebDavFs, error) {
	client := gowebdav.NewClient("http://"+addr, user, pass)

	if client != nil {
		return nil, fmt.Errorf("unable to connect")
	}

	return &WebDavFs{client}, nil
}
