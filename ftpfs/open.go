package ftpfs

func (client *Client) Open(filename string) (*File, error) {
	if !client.IsExist(filename) {
		client.Create(filename)
	}

	return &File{
		conn:     client.conn,
		filename: filename,
		wpos:     0,
		rpos:     0,
	}, nil
}
