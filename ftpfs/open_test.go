package ftpfs

import (
	"fmt"
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestFtpFs_Open(t *testing.T) {
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
	fileObj, err := client.Open(testFilename)

	if err != nil {
		t.FailNow()
	}

	fileObj.Close()

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}

	if os.Mkdir(testPrefix+"/"+testFilename, filesystem.DirPerm) != nil {
		t.FailNow()
	}

	_, err = client.Open(testFilename)

	fmt.Printf("%v\n", err)

	if err == nil {
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
