package localfs

import (
	"os"
	"testing"
)

func TestLocalFs_Create(t *testing.T) {
	client := Init("/")
	testFilename := os.TempDir() + "/test"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if client.Create(testFilename) != nil {
		t.FailNow()
	}

	if client.Create(testFilename) == nil {
		t.FailNow()
	}

	if os.Remove(testFilename) != nil {
		t.FailNow()
	}
}

func TestLocalFs_CreateDir(t *testing.T) {
	client := Init("/")
	testFilename := os.TempDir() + "/test"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if client.CreateDir(testFilename) != nil {
		t.FailNow()
	}

	if client.CreateDir(testFilename) == nil {
		t.FailNow()
	}

	if os.Remove(testFilename) != nil {
		t.FailNow()
	}
}
