package localfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestLocalFs_Remove(t *testing.T) {
	client := Init("/")
	testFilename := os.TempDir() + "/test"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if os.Mkdir(testFilename, filesystem.PermDir) != nil {
		t.FailNow()
	}

	if client.Remove(testFilename) == nil {
		t.FailNow()
	}

	if os.Remove(testFilename) != nil {
		t.FailNow()
	}

	fileObj, err := os.Create(testFilename)

	if err != nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	if client.Remove(testFilename) != nil {
		t.FailNow()
	}
}

func TestLocalFs_RemoveDir(t *testing.T) {
	client := Init("/")
	testFilename := os.TempDir() + "/test"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if os.Mkdir(testFilename, filesystem.PermDir) != nil {
		t.FailNow()
	}

	if client.RemoveDir(testFilename) != nil {
		t.FailNow()
	}

	fileObj, err := os.Create(testFilename)

	if err != nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	if client.RemoveDir(testFilename) == nil {
		t.FailNow()
	}
}
