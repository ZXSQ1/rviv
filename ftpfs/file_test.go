package ftpfs

import (
	"os"
	"slices"
	"testing"
)

func TestFile_Write(t *testing.T) {
	server := openTestServer()
	client, err := Connect(testIp, testPort, testUser, testPass)
	testFilename := "test"
	testContent := "test content for the file: " + testFilename

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

	fileObj, err := client.Open(testFilename)

	if err != nil {
		t.FailNow()
	}

	_, err = fileObj.Write([]byte(testContent))

	if err != nil {
		t.FailNow()
	}

	_, err = fileObj.Write([]byte(testContent))

	if err != nil {
		t.FailNow()
	}

	fileObj.Close()

	bytes, err := os.ReadFile(testPrefix + "/" + testFilename)

	if err != nil {
		t.FailNow()
	}

	if !slices.Equal(bytes, []byte(testContent+testContent)) {
		t.FailNow()
	}
}
