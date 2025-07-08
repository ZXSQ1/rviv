package compressor

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/localfs"
)

var (
	testContent    = slices.Repeat([]byte(testRawContent), 1_000)
	testRawContent = "Sed ut perspiciatis unde omnis iste natus error sit."
)

func TestXz(t *testing.T) {
	client, err := localfs.Init("/")

	if err != nil {
		t.FailNow()
	}

	testFilename := config.Path{
		Filename: os.TempDir() + "/test",
		Fsys:     client,
	}

	testArchivename := config.Path{
		Filename: os.TempDir() + "/test.xz",
		Fsys:     client,
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

	compressor, err := NewXzCompressor(testFilename)

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
