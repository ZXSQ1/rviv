package procs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestRemove(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testPrefix := "c"
	testEntries := map[string]bool{
		"a/b": true, "c/d": true, "e": false,
		"f": false, "g/h": false,
	}

	entries := []config.Path{}
	prefix := config.Path{
		Filename: testPrefix,
		Active:   true,
		Fsys:     client,
	}

	t.Cleanup(func() {
		RemoveDir(prefix)
		prefix.Fsys.Close()
	})

	for testEntry, isDir := range testEntries {
		entry := config.Path{
			Filename: filepath.Join(testPrefix, testEntry),
			Active:   true,
			Fsys:     client,
		}

		err := Mkdir(config.MkdirOpts{
			Filenames: []config.Path{
				CopyPath(entry, filepath.Dir(entry.Filename)),
			},

			Parent: true,
		})

		if err != nil {
			t.FailNow()
		}

		if isDir {
			if CreateDir(entry) != nil {
				t.FailNow()
			}
		} else {
			if CreateFile(entry) != nil {
				t.FailNow()
			}
		}

		entries = append(entries, entry)
	}

	err = Remove(config.RemoveOpts{
		Filenames: entries,
		Recursive: false,
	})

	if err == nil {
		t.FailNow()
	}

	if RemoveDir(prefix) != nil {
		t.FailNow()
	}

	for _, entry := range entries {
		isDir := testEntries[strings.TrimLeft(
			strings.Replace(entry.Filename, testPrefix, "", 1), "/",
		)]

		err := Mkdir(config.MkdirOpts{
			Filenames: []config.Path{
				CopyPath(entry, filepath.Dir(entry.Filename)),
			},

			Parent: true,
		})

		if err != nil {
			t.FailNow()
		}

		if isDir {
			if CreateDir(entry) != nil {
				t.FailNow()
			}
		} else {
			if CreateFile(entry) != nil {
				t.FailNow()
			}
		}
	}

	err = Remove(config.RemoveOpts{
		Filenames: entries,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}
}
