package processes

func Move(src, dest *Path, progress chan int) error {
	err := Copy(src, dest, progress)

	if err != nil {
		return err
	}

	return src.Filesys.Remove(src.Filename)
}
