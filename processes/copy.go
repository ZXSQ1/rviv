package processes

import (
	"errors"
	"io"
)

func Copy(src, dest *Path, progress chan int) error {
	srcfs, destfs := src.Filesys, dest.Filesys
	srcobj, err := srcfs.Open(src.Filename)

	if err != nil {
		return err
	}

	defer srcobj.Close()
	destobj, err := destfs.Open(dest.Filename)

	if err != nil {
		return err
	}

	defer destobj.Close()

	for {
		buffer := make([]byte, BufferSize)
		n, err := srcobj.Read(buffer)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		_, err = destobj.Write(buffer[:n])

		if err != nil {
			return err
		}

		progress <- n
	}

	return nil
}
