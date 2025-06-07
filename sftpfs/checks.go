package sftpfs

import "os"

func (client *SFtpFs) IsExist(filename string) bool {
	_, err := client.conn.Stat(filename)
	return !os.IsNotExist(err)
}
