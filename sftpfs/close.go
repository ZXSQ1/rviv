package sftpfs

func (client *SFtpFs) Close() error {
	return client.conn.Close()
}
