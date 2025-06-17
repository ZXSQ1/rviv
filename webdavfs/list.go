package webdavfs

func (client *WebDavFs) ListDir(filename string) ([]string, error) {
	results := []string{}
	entries, err := client.conn.ReadDir(filename)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		results = append(results, filename+"/"+entry.Name())
	}

	return results, nil
}
