package webdavfs

import "os"

func (client *WebDavFs) IsExist(filename string) bool {
	_, err := client.conn.Stat(filename)
	return !os.IsNotExist(err)
}
