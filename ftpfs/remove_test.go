package ftpfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestFtpFs_Remove(t *testing.T) {
	server := openTestServer()
	client, err := Connect(testIp, testPort, testUser, testPass)
	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testFilename)
		client.Close()
		server.Shutdown()
	})

	if os.Mkdir(testPrefix+"/"+testFilename, filesystem.DirPerm) != nil {
		t.FailNow()
	}

	if client.Remove(testFilename) == nil {
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

	if client.Remove(testFilename) != nil {
		t.FailNow()
	}
}
