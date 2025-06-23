package localfs

import (
	"os"
	"testing"
	"time"

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

	if !actualStat.ModTime().Truncate(time.Second).Equal(
		resultStat.ModTime().Truncate(time.Second)) {

		t.FailNow()
	}

	if actualStat.Mode() != resultStat.Mode() {
		t.FailNow()
	}

	if actualStat.Name() != resultStat.Name() {
		t.FailNow()
	}

	if resultStat.Size() != -1 {
		t.FailNow()
	}

	if os.Remove(testFilename) != nil {
		t.FailNow()
	}

	fileObj, err := os.Create(testFilename)

	if err != nil {
		t.FailNow()
	}

	fileObj.Close()
	actualStat, err = os.Stat(testFilename)

	if err != nil {
		t.FailNow()
	}

	resultStat, err = client.Stat(testFilename)

	if err != nil {
		t.FailNow()
	}

	if actualStat.IsDir() != resultStat.IsDir() {
		t.FailNow()
	}

	if !actualStat.ModTime().Truncate(time.Second).Equal(
		resultStat.ModTime().Truncate(time.Second)) {

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
