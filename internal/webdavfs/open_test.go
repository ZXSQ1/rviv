package webdavfs

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func TestWebDavFs_Open(t *testing.T) {
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
	testContent := "test"

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
		client.Close()
		server.Close()
	})

	createdObj, err := os.Create(testPrefix + "/" + testFilename)

	if err != nil {
		t.FailNow()
	}

	if createdObj.Close() != nil {
		t.FailNow()
	}

	fileObj, err := client.Open(testFilename, filesystem.ModeWrite)

	if err != nil {
		t.FailNow()
	}

	n, err := fileObj.Write([]byte(testContent))

	if n != len(testContent) || err != nil {
		t.FailNow()
	}

	n, err = fileObj.Write([]byte(testContent))

	if n != len(testContent) || err != nil {
		t.FailNow()
	}

	_, err = fileObj.Read([]byte(testContent))

	if err == nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	fileObj, err = client.Open(testFilename, filesystem.ModeRead)

	if err != nil {
		t.FailNow()
	}

	buffer := make([]byte, len(testContent))
	n, err = fileObj.Read(buffer)

	if n != len(testContent) || err != nil || string(buffer) != testContent {
		t.FailNow()
	}

	n, err = fileObj.Read(buffer)

	if n != len(testContent) || !errors.Is(err, io.EOF) || string(buffer) != testContent {
		t.FailNow()
	}

	_, err = fileObj.Write(buffer)

	if err == nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}
}
