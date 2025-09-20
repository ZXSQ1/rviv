package ftpfs

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func TestFtpFs_List(t *testing.T) {
	server := OpenTestServer()
	client, err := Connect(&filesystem.ConnInfo{
		Addr: testAddr,
		User: testUser,
		Pass: testPass,
	})

	testEntryPrefix := "test"
	testEntries := []string{"hi", "abcd", "jsdf", "foo", "bar", "feh"}

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.RemoveAll(testPrefix + "/" + testEntryPrefix)
		client.Close()
		server.Stop()
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

	resultEntries, err := client.List(testEntryPrefix)

	if err != nil {
		t.FailNow()
	}

	slices.Sort(actualEntries)
	slices.Sort(resultEntries)

	if !slices.Equal(actualEntries, resultEntries) {
		t.FailNow()
	}
}
