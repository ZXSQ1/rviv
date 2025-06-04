package ftpfs

func (client *FtpFs) Remove(filename string) error {
	return client.conn.Delete(filename)
}

func (client *FtpFs) RemoveDir(filename string) error {
	return client.conn.RemoveDirRecur(filename)
}
