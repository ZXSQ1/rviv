package sftpfs

import (
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestConnect(t *testing.T) {
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

	_, err = Connect(&filesystem.ConnInfo{})

	if err == nil {
		t.FailNow()
	}
}
