package ftpfs

func (client *Client) OpenFile(filename string) (*File, error) {
	if !client.IsExist(filename) {
		client.CreateFile(filename)
	}

	return &File{
		conn:     client.conn,
		filename: filename,
		wpos:     0,
		rpos:     0,
	}, nil
}
