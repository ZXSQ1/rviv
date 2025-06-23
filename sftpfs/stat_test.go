package sftpfs

import (
	"os"
	"testing"
	"time"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestSFtpFs_Stat(t *testing.T) {
	server := OpenTestServer()
	client, err := Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	if err != nil {
		t.FailNow()
	}

	testFilename := "test"

	t.Cleanup(func() {
		client.Close()
		server.Close()
	})

	if os.Mkdir(testPrefix+"/"+testFilename, filesystem.PermDir) != nil {
		t.FailNow()
	}

	actualStat, err := os.Stat(testPrefix + "/" + testFilename)

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

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}

	fileObj, err := os.Create(testPrefix + "/" + testFilename)

	if err != nil {
		t.FailNow()
	}

	fileObj.Close()
	actualStat, err = os.Stat(testPrefix + "/" + testFilename)

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
