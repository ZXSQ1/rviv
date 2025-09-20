package ftpfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func TestFtpFs_IsExist(t *testing.T) {
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

	if os.MkdirAll(testPrefix+"/"+testFilename, filesystem.PermDir) != nil {
		t.FailNow()
	}

	if !client.IsExist(testFilename) {
		t.FailNow()
	}

	if os.RemoveAll(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}

	if client.IsExist(testFilename) {
		t.FailNow()
	}
}
