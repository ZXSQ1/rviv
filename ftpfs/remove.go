package ftpfs

func (client *FTPFS) Remove(filename string) error {
	return client.conn.Delete(filename)
}

func (client *FTPFS) RemoveDir(filename string) error {
	return client.conn.RemoveDirRecur(filename)
}
