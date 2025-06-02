package ftpfs

func (client *Client) ListDir(filename string) ([]string, error) {
	entries, err := client.conn.NameList(filename)

	if err != nil {
		return nil, err
	}

	return entries, err
}
