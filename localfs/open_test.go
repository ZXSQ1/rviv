package localfs

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestLocalFs_Open(t *testing.T) {
	client := Init("/")
	testFilename := os.TempDir() + "/test"
	testContent := "abcdefghijklmnopqrstuvwxyz"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if os.WriteFile(testFilename, []byte(""), filesystem.PermRegular) != nil {
		t.FailNow()
	}

	fileObj, err := client.Open(testFilename, filesystem.ModeWrite)

	if err != nil {
		t.FailNow()
	}

	_, err = fileObj.Write([]byte(testContent))

	if err != nil {
		t.FailNow()
	}

	_, err = fileObj.Write([]byte(testContent))

	if err != nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	fileObj, err = client.Open(testFilename, filesystem.ModeRead)

	if err != nil {
		t.FailNow()
	}

	buffer := make([]byte, len(testContent)*2)
	_, err = fileObj.Read(buffer)

	if err != nil {
		t.FailNow()
	}

	if !slices.Equal(buffer, []byte(testContent+testContent)) {
		t.FailNow()
	}
}
