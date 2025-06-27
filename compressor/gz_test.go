package compressor

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestGz(t *testing.T) {
	testFilename := filesystem.Path{
		Filename: os.TempDir() + "/test",
		Filesys:  localfs.Init("/"),
	}

	testArchivename := filesystem.Path{
		Filename: os.TempDir() + "/test.gz",
		Filesys:  localfs.Init("/"),
	}

	t.Cleanup(func() {
		testFilename.Filesys.Remove(testFilename.Filename)
		testArchivename.Filesys.Remove(testArchivename.Filename)
	})

	if testFilename.Filesys.Create(testFilename.Filename) != nil {
		t.FailNow()
	}

	fileObj, err := testFilename.Filesys.Open(
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

	if !testArchivename.Filesys.IsExist(testArchivename.Filename) {
		t.FailNow()
	}
}
