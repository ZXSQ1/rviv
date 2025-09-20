package ftpfs

func (client *FtpFs) Close() error {
	return stderr(client.conn.Quit())
}
