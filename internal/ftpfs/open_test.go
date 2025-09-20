package ftpfs

import (
	"io"
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func TestFtpFs_Open(t *testing.T) {
	server := OpenTestServer()
	client, err := Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	testFilename := "test"
	testContent := "test"

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testPrefix + "/" + testFilename)
		client.Close()
		server.Stop()
	})

	if client.CreateFile(testFilename) != nil {
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

	_, err = fileObj.Read(make([]byte, filesystem.BufferSize))

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

	if n != len(testContent) || err != nil {
		t.FailNow()
	}

	n, err = fileObj.Read(buffer)

	if n != len(testContent) || err != nil {
		t.FailNow()
	}

	n, err = fileObj.Read(buffer)

	if n != 0 || err.Error() != io.EOF.Error() {
		t.FailNow()
	}

	_, err = fileObj.Write([]byte{})

	if err == nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	if os.Remove(testPrefix+"/"+testFilename) != nil {
		t.FailNow()
	}

	_, err = client.Open(testFilename, filesystem.ModeWrite)

	if err == nil {
		t.FailNow()
	}
}
