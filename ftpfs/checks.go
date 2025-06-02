package ftpfs

import (
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/jlaffaye/ftp"
)

func (client *Client) IsExist(filename string) (bool, error) {
	_, err := client.conn.FileSize(filename)
	return status(err) != ftp.StatusFileUnavailable, nil
}

func (client *Client) Type(filename string) (int8, error) {
	stat, err := client.conn.GetEntry(filename)

	if err != nil {
		return -1, err
	}

	switch stat.Type {
	case ftp.EntryTypeFile:
		return filesystem.TypeFile, nil
	case ftp.EntryTypeFolder:
		return filesystem.TypeDir, nil
	default:
		return filesystem.TypeLink, nil
	}
}

func (client *Client) Size(filename string) (uint64, error) {
	size, err := client.conn.FileSize(filename)

	if err != nil {
		return 0, err
	}

	return uint64(size), nil
}
