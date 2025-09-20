package localfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestLocalFs_RemoveFile(t *testing.T) {
	client, err := Init(testPrefix)
	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
	})

	if os.Mkdir(testPrefix+"/"+testFilename, filesystem.PermDir) != nil {
		t.FailNow()
	}

	if client.RemoveFile(testFilename) == nil {
		t.FailNow()
	}

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}

	fileObj, err := os.Create(testPrefix + "/" + testFilename)

	if err != nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	if client.RemoveFile(testFilename) != nil {
		t.FailNow()
	}
}

func TestLocalFs_RemoveDir(t *testing.T) {
	client, err := Init(testPrefix)
	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
	})

	if os.Mkdir(testPrefix+"/"+testFilename, filesystem.PermDir) != nil {
		t.FailNow()
	}

	if client.RemoveDir(testFilename) != nil {
		t.FailNow()
	}

	fileObj, err := os.Create(testPrefix + "/" + testFilename)

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
