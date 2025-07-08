package localfs

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestLocalFs_List(t *testing.T) {
	client, err := Init(testPrefix)
	testEntryPrefix := "test"
	testEntries := []string{
		"a", "bah", "shfj", "bag", "feh", "voo",
	}

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.RemoveAll(testPrefix + "/" + testEntryPrefix)
	})

	for _, testEntry := range testEntries {
		if os.MkdirAll(testPrefix+"/"+testEntryPrefix+"/"+testEntry,
			filesystem.PermDir) != nil {

			t.FailNow()
		}
	}

	actualEntries := []string{}
	rawEntries, err := os.ReadDir(testPrefix + "/" + testEntryPrefix)

	if err != nil {
		t.FailNow()
	}

	for _, rawEntry := range rawEntries {
		actualEntries = append(
			actualEntries, testEntryPrefix+"/"+rawEntry.Name(),
		)
	}

	resultEntries, err := client.ListDir(testEntryPrefix)

	if err != nil {
		t.FailNow()
	}

	if !slices.Equal(resultEntries, actualEntries) {
		t.FailNow()
	}
}
