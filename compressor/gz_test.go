package compressor

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestGz(t *testing.T) {
	testFilename := config.Path{
		Filename: os.TempDir() + "/test",
		Fsys:     localfs.Init("/"),
	}

	testArchivename := config.Path{
		Filename: os.TempDir() + "/test.gz",
		Fsys:     localfs.Init("/"),
	}

	t.Cleanup(func() {
		testFilename.Fsys.Remove(testFilename.Filename)
		testArchivename.Fsys.Remove(testArchivename.Filename)
	})

	if testFilename.Fsys.Create(testFilename.Filename) != nil {
		t.FailNow()
	}

	fileObj, err := testFilename.Fsys.Open(
		testFilename.Filename, filesystem.ModeWrite)

	if err != nil {
		t.FailNow()
	}

	if _, err = fileObj.Write(testContent); err != nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	compressor, err := NewGzCompressor(testFilename)

	if err != nil {
		t.FailNow()
	}

	if compressor.Compress(testArchivename, 6) != nil {
		t.FailNow()
	}

	if !testArchivename.Fsys.IsExist(testArchivename.Filename) {
		t.FailNow()
	}
}
