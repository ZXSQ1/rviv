package ftpfs

import (
	"os"
	"testing"
)

func TestFileInfo(t *testing.T) {
	server := openTestServer()
	client, err := Connect(testIp, testPort, testUser, testPass)
	testFilename := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
		client.Close()
		server.Shutdown()
	})

	if client.Create(testFilename) != nil {
		t.FailNow()
	}

	serverInfo, err := client.Stat(testFilename)

	if err != nil {
		t.FailNow()
	}

	localInfo, err := os.Stat(testPrefix + "/" + testFilename)

	if err != nil {
		t.FailNow()
	}

	if serverInfo.Name() != localInfo.Name() {
		t.FailNow()
	}

	if serverInfo.Size() != localInfo.Size() {
		t.FailNow()
	}

	// testing for time would not be of any use; ftp server
	// does not support many operations including time for
	// entries. the implementation of the ModTime method
	// is similar to the implementation of the other methods;
	// not necessary to test

	if serverInfo.Mode() != localInfo.Mode() {
		t.FailNow()
	}

	if serverInfo.IsDir() != localInfo.IsDir() {
		t.FailNow()
	}
}
