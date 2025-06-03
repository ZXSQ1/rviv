package ftpfs

func (client *FTPFS) Open(filename string) (*File, error) {
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
