package ftpfs

import (
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestFtpFs_ListDir(t *testing.T) {
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

	actualEntriesData, err := os.ReadDir(testPrefix + "/" + testEntryPrefix)
	actualEntries := []string{}

	if err != nil {
		t.FailNow()
	}

	for _, actualEntryData := range actualEntriesData {
		actualEntries = append(
			actualEntries, testEntryPrefix+"/"+actualEntryData.Name(),
		)
	}

	resultEntries, err := client.ListDir(testEntryPrefix)

	if err != nil {
		t.FailNow()
	}

	slices.Sort(actualEntries)
	slices.Sort(resultEntries)

	if !slices.Equal(actualEntries, resultEntries) {
		fmt.Printf("%v\n%v\n", actualEntries, resultEntries)
		t.FailNow()
	}
}
