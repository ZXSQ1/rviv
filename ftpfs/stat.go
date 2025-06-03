package ftpfs

import (
	"io/fs"
	"os"
)

func (client *Client) Stat(filename string) (fs.FileInfo, error) {
	if !client.IsExist(filename) {
		return nil, os.ErrNotExist
	}

	return FileInfo{
		filename: filename,
		conn:     client.conn,
	}, nil
}
