package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestMoveAll(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testDest := "b"
	testPrefix := "a"
	testEntries := map[string]bool{
		"a/b":         true,
		"fhjsd/fsdkj": false,
		"foo/bar":     true,
		"fsjdk/fd":    true,
		"go/good":     false,
		"jk/das":      false,
	}

	entries := []config.Path{}
	prefix := config.Path{
		Filename: testPrefix,
		Active:   true,
		Fsys:     client,
	}

	dest := config.Path{
		Filename: testDest,
		Active:   true,
		Fsys:     client,
	}

	t.Cleanup(func() {
		RemoveDir(prefix)
		RemoveDir(dest)
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

	if CreateDir(dest) != nil {
		t.FailNow()
	}

	if err := MoveAll(entries, dest, false); err != nil {
		t.FailNow()
	}

	for _, entry := range entries {
		if IsExist(entry, nil) == nil {
			t.FailNow()
		}
	}

	destEntries, err := ListDir(config.ListOpts{
		Filename:  dest,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	srcTextEntries := []string{}
	destTextEntries := []string{}

	for idx := range entries {
		srcTextEntries = append(srcTextEntries, filepath.Base(
			entries[idx].Filename))

		destTextEntries = append(destTextEntries, filepath.Base(
			destEntries[idx].Filename))
	}

	slices.Sort(srcTextEntries)
	slices.Sort(destTextEntries)

	for idx := range srcTextEntries {
		srcfile := srcTextEntries[idx]
		destfile := destTextEntries[idx]

		if srcfile != destfile {
			t.FailNow()
		}
	}
}
