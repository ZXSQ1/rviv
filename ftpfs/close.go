package ftpfs

func (client *Client) Close() error {
	return client.conn.Quit()
}
