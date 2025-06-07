package sftpfs

func (client *SFtpFs) Create(filename string) error {
	fileObj, err := client.conn.Create(filename)

	if err != nil {
		return err
	}

	defer fileObj.Close()
	fileObj.Write([]byte(""))

	return nil
}

func (client *SFtpFs) CreateDir(filename string) error {
	return client.conn.MkdirAll(filename)
}
