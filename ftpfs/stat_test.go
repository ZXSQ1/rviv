package ftpfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestFtpFs_Stat(t *testing.T) {
	server := openTestServer()
	client, err := Connect(testIp, testPort, testUser, testPass)
	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		client.Close()
		server.Shutdown()
	})

	createdObj, err := os.Create(testPrefix + "/" + testFilename)

	if err != nil {
		t.FailNow()
	}

	createdObj.Close()
	_, err = client.Stat(testFilename)

	if err != nil {
		t.FailNow()
	}

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}

	if os.Mkdir(testPrefix+"/"+testFilename, filesystem.DirPerm) != nil {
		t.FailNow()
	}

	_, err = client.Stat(testFilename)

	if err != nil {
		t.FailNow()
	}

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}

	_, err = client.Open(testFilename)

	if err == nil {
		t.FailNow()
	}
}
