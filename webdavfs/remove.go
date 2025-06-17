package webdavfs

func (client *WebDavFs) Remove(filename string) error {
	return client.conn.Remove(filename)
}

func (client *WebDavFs) RemoveDir(filename string) error {
	return client.conn.RemoveAll(filename)
}
