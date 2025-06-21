package ftpfs

import (
	"fmt"

	"github.com/jlaffaye/ftp"
)

func (client *FtpFs) Remove(filename string) error {
	entry, err := client.conn.GetEntry(filename)

	if err != nil {
		return stderr(err)
	}

	if entry.Type != ftp.EntryTypeFile {
		return fmt.Errorf("file not a regular file")
	}

	return stderr(client.conn.Delete(filename))
}

func (client *FtpFs) RemoveDir(filename string) error {
	entry, err := client.conn.GetEntry(filename)

	if err != nil {
		return stderr(err)
	}

	if entry.Type != ftp.EntryTypeFolder {
		return fmt.Errorf("file not a directory")
	}

	return stderr(client.conn.RemoveDirRecur(filename))
}
