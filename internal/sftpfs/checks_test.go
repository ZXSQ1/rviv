package sftpfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func TestSFtpFs_IsExist(t *testing.T) {
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
		os.Remove(testPrefix + "/" + testFilename)
		client.Close()
		server.Close()
	})

	if os.Mkdir(testPrefix+"/"+testFilename, filesystem.PermDir) != nil {
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
