package ftpfs

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

	if server.Stop() != nil {
		t.FailNow()
	}

	if client.Close() != nil {
		t.FailNow()
	}

	_, err = Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	if err == nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		client.Close()
		server.Stop()
	})
}
