package localfs

import "os"

func (local *LocalFs) Close() error {
	if err := os.Chdir(local.prefix); err != nil {
		return err
	}

	local.prefix = ""
	local.currdir = ""

	return nil
}
