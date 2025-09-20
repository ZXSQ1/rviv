package sftpfs

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func TestSFtpFs_Create(t *testing.T) {
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

	if client.CreateFile(testFilename) != nil {
		t.FailNow()
	}

	if client.CreateFile(testFilename) == nil {
		t.FailNow()
	}
}
