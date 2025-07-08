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
		if entries[idx].Filename != testPathEntries[1:4][idx].Filename {
			t.FailNow()
		}
	}
}
