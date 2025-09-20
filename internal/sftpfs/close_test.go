package sftpfs

import (
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestSFtpFs_Close(t *testing.T) {
	server := OpenTestServer()
	client, err := Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		client.Close()
		server.Close()
	})

	if client.Close() != nil {
		t.FailNow()
	}

	if client.Close() == nil {
		t.FailNow()
	}
}
