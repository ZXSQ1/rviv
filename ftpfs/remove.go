package ftpfs

func (client *Client) Remove(filename string) error {
	return client.conn.Delete(filename)
}

func (client *Client) RemoveDir(filename string) error {
	return client.conn.RemoveDirRecur(filename)
}
