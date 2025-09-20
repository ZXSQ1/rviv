package localfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func TestLocalFs_IsExist(t *testing.T) {
	testFilename := "test"
	client, err := Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
	})

	if os.MkdirAll(testPrefix+"/"+testFilename, filesystem.PermDir) != nil {
		t.FailNow()
	}

	if !client.IsExist(testFilename) {
		t.FailNow()
	}

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}

	if client.IsExist(testFilename) {
		t.FailNow()
	}

	fileObj, err := os.Create(testPrefix + "/" + testFilename)

	if err != nil {
		t.FailNow()
	}

	fileObj.Close()

	if !client.IsExist(testFilename) {
		t.FailNow()
	}
}
