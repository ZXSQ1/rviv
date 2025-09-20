package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/localfs"
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
	pathEntries := []config.Path{}
	pathPrefix := config.Path{
		Filename: testPrefix,
		Active:   true,
		Fsys:     client,
	}

	for _, testEntry := range testEntries {
		pathEntries = append(pathEntries, config.Path{
			Filename: testPrefix + "/" + testEntry,
			Active:   true,
			Fsys:     client,
		})
	}

	t.Cleanup(func() {
		client.RemoveDir(testPrefix)
		client.Close()
	})

	if CreateDir(pathPrefix) != nil {
		t.FailNow()
	}

	for _, entry := range pathEntries {
		entryDir := CopyPath(entry, filepath.Dir(entry.Filename))

		if IsExist(entryDir, CreateDir) != nil {
			t.FailNow()
		}

		if CreateDir(entry) != nil {
			t.FailNow()
		}
	}

	paths, err := ListDir(config.ListOpts{
		Filename:  pathPrefix,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range paths {
		if paths[idx].Filename != pathEntries[idx].Filename {
			t.FailNow()
		}
	}

	if RemoveDir(pathPrefix) != nil {
		t.FailNow()
	}

	if CreateDir(pathPrefix) != nil {
		t.FailNow()
	}

	for _, entry := range pathEntries {
		entryDir := CopyPath(entry, filepath.Dir(entry.Filename))

		if IsExist(entryDir, CreateDir) != nil {
			t.FailNow()
		}

		if CreateDir(entry) != nil {
			t.FailNow()
		}
	}

	paths, err = ListDir(config.ListOpts{
		Filename:  pathPrefix,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range paths {
		if paths[idx].Filename != pathEntries[idx].Filename {
			t.FailNow()
		}
	}
}
