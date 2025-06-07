package sftpfs

func (client *SFtpFs) Remove(filename string) error {
	return client.conn.Remove(filename)
}

func (client *SFtpFs) RemoveDir(filename string) error {
	return client.conn.RemoveAll(filename)
}
