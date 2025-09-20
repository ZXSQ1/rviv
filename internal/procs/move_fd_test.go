package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestMoveFiles(t *testing.T) {
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
		for _, src := range srcs {
			RemoveFile(src)
		}

		RemoveDir(destdir)
		destdir.Fsys.Close()
	})

	if CreateDir(destdir) != nil {
		t.FailNow()
	}

	for _, src := range srcs {
		if CreateFile(src) != nil {
			t.FailNow()
		}
	}

	if MoveFiles(srcs, destdir, false) != nil {
		t.FailNow()
	}

	for _, src := range srcs {
		if src.Fsys.IsExist(src.Filename) {
			t.FailNow()
		}
	}

	destEntries, err := List(destdir)

	if err != nil {
		t.FailNow()
	}

	slices.Sort(testSrcFiles)

	for idx := range testSrcFiles {
		if filepath.Base(destEntries[idx].Filename) != testSrcFiles[idx] {
			t.FailNow()
		}
	}
}
