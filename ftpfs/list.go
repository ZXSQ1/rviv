package ftpfs

func (client *FtpFs) ListDir(filename string) ([]string, error) {
	entries, err := client.conn.NameList(filename)
	err = stderr(err)

	return entries, err
}
