package sftpfs

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestSFtpFs_ListDir(t *testing.T) {
	server := OpenTestServer()
	client, err := Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	if err != nil {
		t.FailNow()
	}

	testEntryPrefix := "test"
	testEntries := []string{"a", "hfsjd", "b", "fh", "jk", "alsd"}

	t.Cleanup(func() {
		os.RemoveAll(testEntryPrefix)
		client.Close()
		server.Close()
	})

	for _, testEntry := range testEntries {
		if os.MkdirAll(testPrefix+"/"+testEntryPrefix+"/"+testEntry,
			filesystem.PermDir) != nil {

			t.FailNow()
		}
	}

	actualEntries := slices.Clone(testEntries)

	for idx, actualEntry := range actualEntries {
		actualEntries[idx] = testEntryPrefix + "/" + actualEntry
	}

	resultEntries, err := client.ListDir(testEntryPrefix)

	if err != nil {
		t.FailNow()
	}

	slices.Sort(actualEntries)
	slices.Sort(resultEntries)

	if !slices.Equal(actualEntries, resultEntries) {
		t.FailNow()
	}
}
