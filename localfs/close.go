package localfs

// implementation of the Close method for the local filesystem; does not do
// anything in the case of a local filesystem
func (local *LocalFs) Close() error {
	return nil
}
