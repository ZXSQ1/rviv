package ftpfs

func (client *Client) GetFile(filename string) (*File, error) {
	exists, err := client.IsExist(filename)

	if err != nil {
		return nil, err
	}

	if !exists {
		client.CreateFile(filename)
	}

	return &File{
		conn:     client.conn,
		filename: filename,
		wpos:     0,
		rpos:     0,
	}, nil
}
