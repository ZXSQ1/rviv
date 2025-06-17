package sftpfs

func (client *SFtpFs) Create(filename string) error {
	fileObj, err := client.conn.Create(filename)

	if err != nil {
		return err
	}

	return fileObj.Close()
}

func (client *SFtpFs) CreateDir(filename string) error {
	return client.conn.Mkdir(filename)
}
