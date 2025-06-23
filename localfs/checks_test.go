package localfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestLocalFs_IsExist(t *testing.T) {
	client := Init()
	testFilename := os.TempDir() + "/test"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if os.MkdirAll(testFilename, filesystem.PermDir) != nil {
		println("error")
		t.FailNow()
	}

	if !client.IsExist(testFilename) {
		t.FailNow()
	}

	if os.Remove(testFilename) != nil {
		t.FailNow()
	}

	if client.IsExist(testFilename) {
		t.FailNow()
	}

	fileObj, err := os.Create(testFilename)

	if err != nil {
		t.FailNow()
	}

	fileObj.Close()

	if !client.IsExist(testFilename) {
		t.FailNow()
	}
}
