package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestListDir(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testPrefix := "test"
	testEntries := []string{
		"a/b", "sdf/a", "a/j", "fsjdkf/fsd", "asd/d", "asd/f", "asd/g",
		"a/c/csd", "a/c/d",
	}

	slices.Sort(testEntries)
	testPathEntries := []config.Path{}
	testPathPrefix := config.Path{
		Filename: testPrefix,
		Active:   true,
		Fsys:     client,
	}

	for _, testEntry := range testEntries {
		testPathEntries = append(testPathEntries, config.Path{
			Filename: testPrefix + "/" + testEntry,
			Active:   true,
			Fsys:     client,
		})
	}

	t.Cleanup(func() {
		client.RemoveDir(testPrefix)
		client.Close()
	})

	if testPathPrefix.Fsys.CreateDir(testPathPrefix.Filename) != nil {
		t.FailNow()
	}

	for _, entry := range testPathEntries {
		entryName := entry.Filename
		entryDir := filepath.Dir(entryName)

		if !entry.Fsys.IsExist(entryDir) &&
			entry.Fsys.CreateDir(entryDir) != nil {

			t.FailNow()
		}

		if entry.Fsys.CreateDir(entryName) != nil {
			t.FailNow()
		}
	}

	paths, err := ListDir(config.ListOpts{
		Filename:  testPathPrefix,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range paths {
		if paths[idx].Filename != testPathEntries[idx].Filename {
			t.FailNow()
		}
	}
}
