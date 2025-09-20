package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestGlob(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testPrefix := "test"
	testGlob := testPrefix + "/a/d*"
	testEntries := []string{
		"a/b", "a/j", "a/dsfsd", "a/ds", "a/d",
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
		RemoveDir(pathPrefix)
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

	entries, err := Glob(config.GlobOpts{
		Entries: []config.Path{
			{
				Filename: testGlob,
				Active:   true,
				Fsys:     client,
			},
		},
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range entries {
		if entries[idx].Filename != pathEntries[1:4][idx].Filename {
			t.FailNow()
		}
	}
}
