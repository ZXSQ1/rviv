package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestDiffer(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testSrcdir := "a"
	testSrcEntries := []string{"a", "b", "c", "d", "e"}
	expectedSrcDiffs := []string{"a", "b"}
	expectedSim := []string{"c", "d", "e"}
	testDestdir := "b"
	testDestEntries := []string{"c", "d", "e", "f", "g"}
	expectedDestDiffs := []string{"f", "g"}

	srcEntries := []config.Path{}
	srcdir := config.Path{
		Filename: testSrcdir,
		Active:   true,
		Fsys:     client,
	}

	destEntries := []config.Path{}
	destdir := config.Path{
		Filename: testDestdir,
		Active:   true,
		Fsys:     client,
	}

	for _, testSrcEntry := range testSrcEntries {
		srcEntries = append(srcEntries, config.Path{
			Filename: filepath.Join(testSrcdir, testSrcEntry),
			Active:   true,
			Fsys:     client,
		})
	}

	for _, testDestEntry := range testDestEntries {
		destEntries = append(destEntries, config.Path{
			Filename: filepath.Join(testDestdir, testDestEntry),
			Active:   true,
			Fsys:     client,
		})
	}

	t.Cleanup(func() {
		srcdir.Fsys.RemoveDir(srcdir.Filename)
		destdir.Fsys.RemoveDir(destdir.Filename)
		destdir.Fsys.Close()
	})

	if srcdir.Fsys.CreateDir(srcdir.Filename) != nil {
		t.FailNow()
	}

	for _, srcEntry := range srcEntries {
		if srcEntry.Fsys.Create(srcEntry.Filename) != nil {
			t.FailNow()
		}
	}

	if destdir.Fsys.CreateDir(destdir.Filename) != nil {
		t.FailNow()
	}

	for _, destEntry := range destEntries {
		if destEntry.Fsys.Create(destEntry.Filename) != nil {
			t.FailNow()
		}
	}

	diffs, err := Differ(config.DifferOpts{
		Src:  srcdir,
		Dest: destdir,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range diffs.UniqueSrcEntries {
		diffBase := filepath.Base(diffs.UniqueSrcEntries[idx].Filename)

		if !slices.Contains(expectedSrcDiffs, diffBase) {
			t.FailNow()
		}
	}

	for idx := range diffs.UniqueDestEntries {
		diffBase := filepath.Base(diffs.UniqueDestEntries[idx].Filename)

		if !slices.Contains(expectedDestDiffs, diffBase) {
			t.FailNow()
		}
	}

	for idx := range diffs.SrcCommonEntries {
		diffBase := filepath.Base(diffs.SrcCommonEntries[idx].Filename)

		if !slices.Contains(expectedSim, diffBase) {
			t.FailNow()
		}
	}
}
