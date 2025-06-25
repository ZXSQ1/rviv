package webdavfs

import (
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestConnect(t *testing.T) {
	server := OpenTestServer()
	_, err := Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		server.Close()
	})

	if _, err := Connect(&filesystem.ConnInfo{}); err == nil {
		t.FailNow()
	}
}
