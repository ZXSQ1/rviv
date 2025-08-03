package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestCopyAll(t *testing.T) {
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

	for testEntry, isDir := range testEntries {
		entry := config.Path{
			Filename: filepath.Join(testPrefix, testEntry),
			Active:   true,
			Fsys:     client,
		}

		err := Mkdir(config.MkdirOpts{
			Filenames: []config.Path{
				{
					Filename: filepath.Dir(entry.Filename),
					Active:   entry.Active,
					Fsys:     entry.Fsys,
				},
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

	if dest.Fsys.CreateDir(dest.Filename) != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		prefix.Fsys.RemoveDir(prefix.Filename)
		dest.Fsys.RemoveDir(dest.Filename)
		prefix.Fsys.Close()
	})

	if err := CopyAll(entries, dest, false); err != nil {
		t.FailNow()
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
