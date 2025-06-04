package ftpfs

func (client *FtpFs) Close() error {
	return client.conn.Quit()
}
