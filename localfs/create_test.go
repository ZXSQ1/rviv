package localfs

import (
	"os"
	"testing"
)

func TestLocalFs_CreateFile(t *testing.T) {
	client, err := Init(testPrefix)
	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
	})

	if client.CreateFile(testFilename) != nil {
		t.FailNow()
	}

	if err := client.CreateFile(testFilename); err == nil {
		t.FailNow()
	}

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}
}

func TestLocalFs_CreateDir(t *testing.T) {
	client, err := Init(testPrefix)
	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
	})

	if client.CreateDir(testFilename) != nil {
		t.FailNow()
	}

	if client.CreateDir(testFilename) == nil {
		t.FailNow()
	}

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}
}
