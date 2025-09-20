package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestCopyFiles(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testSrcFiles := []string{"a", "b", "asd", "das"}
	testDestDir := "test"

	srcs := []config.Path{}
	destdir := config.Path{
		Filename: testDestDir,
		Active:   true,
		Fsys:     client,
	}

	for _, src := range testSrcFiles {
		srcs = append(srcs, config.Path{
			Filename: src,
			Active:   true,
			Fsys:     client,
		})
	}

	t.Cleanup(func() {
		RemoveDir(destdir)
		destdir.Fsys.Close()
	})

	if err := CreateDir(destdir); err != nil {
		println(err)
		t.FailNow()
	}

	for _, src := range srcs {
		if CreateFile(src) != nil {
			t.FailNow()
		}
	}

	if CopyFiles(srcs, destdir, false) != nil {
		t.FailNow()
	}

	destEntries, err := destdir.Fsys.List(destdir.Filename)

	if err != nil {
		t.FailNow()
	}

	slices.Sort(testSrcFiles)
	slices.Sort(destEntries)

	for idx := range testSrcFiles {
		if filepath.Base(destEntries[idx]) != testSrcFiles[idx] {
			t.FailNow()
		}
	}
}
