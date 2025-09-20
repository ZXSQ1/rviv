package ftpfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func TestFtpFs_Remove(t *testing.T) {
	server := OpenTestServer()
	client, err := Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
		client.Close()
		server.Stop()
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

	fileObj.Close()

	if client.RemoveFile(testFilename) != nil {
		t.FailNow()
	}
}

func TestFtpFs_RemoveDir(t *testing.T) {
	server := OpenTestServer()
	client, err := Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
		client.Close()
		server.Stop()
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
