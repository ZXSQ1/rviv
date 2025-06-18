package localfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestLocalFs_Stat(t *testing.T) {
	client := Init()
	testFilename := os.TempDir() + "/test"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if os.Mkdir(testFilename, filesystem.PermDir) != nil {
		t.FailNow()
	}

	actualStat, err := os.Stat(testFilename)

	if err != nil {
		t.FailNow()
	}

	resultStat, err := client.Stat(testFilename)

	if err != nil {
		t.FailNow()
	}

	if actualStat.IsDir() != resultStat.IsDir() {
		t.FailNow()
	}

	if actualStat.ModTime() != resultStat.ModTime() {
		t.FailNow()
	}

	if actualStat.Mode() != resultStat.Mode() {
		t.FailNow()
	}

	if actualStat.Name() != resultStat.Name() {
		t.FailNow()
	}

	if actualStat.Size() != resultStat.Size() {
		t.FailNow()
	}
}
