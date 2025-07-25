package ftpfs

func (client *FtpFs) List(filename string) ([]string, error) {
	entries, err := client.conn.NameList(filename)
	err = stderr(err)

	return entries, err
}
