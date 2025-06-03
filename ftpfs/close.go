package ftpfs

func (client *FTPFS) Close() error {
	return client.conn.Quit()
}
