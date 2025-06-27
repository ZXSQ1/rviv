package localfs

import (
	"os"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestLocalFs_List(t *testing.T) {
	client := Init("/")
	testPrefix := os.TempDir() + "/test"
	testEntries := []string{
		"a", "bah", "shfj", "bag", "feh", "voo",
	}

	t.Cleanup(func() {
		os.RemoveAll(testPrefix)
	})

	for _, testEntry := range testEntries {
		if os.MkdirAll(testPrefix+"/"+testEntry, filesystem.PermDir) != nil {
			t.FailNow()
		}
	}

	actualEntries := []string{}
	rawEntries, err := os.ReadDir(testPrefix)

	if err != nil {
		t.FailNow()
	}

	for _, rawEntry := range rawEntries {
		actualEntries = append(actualEntries, testPrefix+"/"+rawEntry.Name())
	}

	resultEntries, err := client.ListDir(testPrefix)

	if err != nil {
		t.FailNow()
	}

	if !slices.Equal(resultEntries, actualEntries) {
		t.FailNow()
	}
}
