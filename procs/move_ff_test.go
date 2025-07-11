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
	testSrcfile := config.Path{
		Filename: "src",
		Active:   true,
		Fsys:     client,
	}

	testDestfile := config.Path{
		Filename: "dest",
		Active:   true,
		Fsys:     client,
	}

	t.Cleanup(func() {
		testSrcfile.Fsys.Remove(testSrcfile.Filename)
		testDestfile.Fsys.Remove(testDestfile.Filename)
		testDestfile.Fsys.Close()
	})

	if testSrcfile.Fsys.Create(testSrcfile.Filename) != nil {
		t.FailNow()
	}

	srcObj, err := testSrcfile.Fsys.Open(testSrcfile.Filename,
		filesystem.ModeWrite)

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

	statSrc, err := testSrcfile.Fsys.Stat(testSrcfile.Filename)

	if err != nil {
		t.FailNow()
	}

	if MoveFile(testSrcfile, testDestfile, false, nil) != nil {
		t.FailNow()
	}

	if testSrcfile.Fsys.IsExist(testSrcfile.Filename) {
		t.FailNow()
	}

	statDest, err := testDestfile.Fsys.Stat(testDestfile.Filename)

	if err != nil {
		t.FailNow()
	}

	if statSrc.Size() != statDest.Size() {
		t.FailNow()
	}
}
