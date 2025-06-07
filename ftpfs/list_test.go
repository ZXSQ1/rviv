package ftpfs

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestFtpFs_List(t *testing.T) {
	server := openTestServer()
	client, err := Connect(testIp, testPort, testUser, testPass)
	filePrefix := "test"
	files := []string{
		"a", "a/b",
		"foo", "bar",
		"foo/bar", "bar/foo",
		"abcd",
	}

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.RemoveAll(testPrefix + "/" + filePrefix)
		client.Close()
		server.Shutdown()
	})

	for _, file := range files {
		if os.MkdirAll(testPrefix+"/"+filePrefix+"/"+file, filesystem.DirPerm) != nil {
			t.FailNow()
		}
	}

	serverList, err := client.ListDir(filePrefix)

	if err != nil {
		t.FailNow()
	}

	localInfoList, err := os.ReadDir(testPrefix + "/" + filePrefix)
	localList := []string{}

	if err != nil {
		t.FailNow()
	}

	for _, localEntry := range localInfoList {
		localList = append(localList, localEntry.Name())
	}

	if !slices.Equal(serverList, localList) {
		t.FailNow()
	}
}
