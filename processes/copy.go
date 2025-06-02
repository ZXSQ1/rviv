package processes

import (
	"errors"
	"io"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

func Copy(src, dest *Path, progress chan int) error {
	srcfs, destfs := src.Filesys, dest.Filesys
	srctype, err := srcfs.Type(src.Filename)

	if err != nil {
		return err
	}

	desttype, err := destfs.Type(dest.Filename)

	if err != nil {
		return err
	}

	if !srcfs.IsExist(src.Filename) {
		return os.ErrNotExist
	} else if srctype == filesystem.TypeDir {
		return os.ErrInvalid
	}

	if destfs.IsExist(dest.Filename) && desttype == filesystem.TypeDir {
		return os.ErrInvalid
	}

	srcobj, err := srcfs.OpenFile(src.Filename)

	if err != nil {
		return err
	}

	defer srcobj.Close()
	destobj, err := destfs.OpenFile(dest.Filename)

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
