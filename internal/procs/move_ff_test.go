package procs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestMoveFile(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testContent := []byte("The brown fox jumps over the lazy dog.")
	src := config.Path{
		Filename: "src",
		Active:   true,
		Fsys:     client,
	}

	dest := config.Path{
		Filename: "dest",
		Active:   true,
		Fsys:     client,
	}

	t.Cleanup(func() {
		RemoveFile(src)
		RemoveFile(dest)
		dest.Fsys.Close()
	})

	if CreateFile(src) != nil {
		t.FailNow()
	}

	srcObj, err := Open(src, filesystem.ModeWrite)

	if err != nil {
		t.FailNow()
	}

	for i := 0; i < 1_000_000; i++ {
		if _, err = srcObj.Write(testContent); err != nil {
			t.FailNow()
		}
	}

	if srcObj.Close() != nil {
		t.FailNow()
	}

	statSrc, err := Stat(src)

	if err != nil {
		t.FailNow()
	}

	if MoveFile(src, dest, false, nil) != nil {
		t.FailNow()
	}

	if IsExist(src, nil) == nil {
		t.FailNow()
	}

	statDest, err := Stat(dest)

	if err != nil {
		t.FailNow()
	}

	if statSrc.Size() != statDest.Size() {
		t.FailNow()
	}
}
