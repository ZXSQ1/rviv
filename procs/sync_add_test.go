package procs

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestSyncAdd(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testSrc := "a"
	testSrcEntries := []string{"a", "b", "c", "d"}
	testDest := "b"
	testDestEntries := []string{"d", "e", "f", "g"}

	expectedDestEntries1 := []string{"a", "b", "c", "d", "e", "f", "g"}
	expectedSrcEntries1 := []string{"a", "b", "c", "d"}
	expectedDestEntries2 := []string{"a", "b", "c", "d", "e", "f", "g"}
	expectedSrcEntries2 := []string{"a", "b", "c", "d", "e", "f", "g"}

	srcEntries := []config.Path{}
	src := config.Path{
		Filename: testSrc,
		Active:   true,
		Fsys:     client,
	}

	destEntries := []config.Path{}
	dest := config.Path{
		Filename: testDest,
		Active:   true,
		Fsys:     client,
	}

	for _, testSrcEntry := range testSrcEntries {
		srcEntries = append(srcEntries, config.Path{
			Filename: filepath.Join(testSrc, testSrcEntry),
			Active:   true,
			Fsys:     client,
		})
	}

	for _, testDestEntry := range testDestEntries {
		destEntries = append(destEntries, config.Path{
			Filename: filepath.Join(testDest, testDestEntry),
			Active:   true,
			Fsys:     client,
		})
	}

	t.Cleanup(func() {
		src.Fsys.RemoveDir(src.Filename)
		dest.Fsys.RemoveDir(dest.Filename)
		dest.Fsys.Close()
	})

	if src.Fsys.CreateDir(src.Filename) != nil {
		t.FailNow()
	}

	for _, srcEntry := range srcEntries {
		if srcEntry.Fsys.Create(srcEntry.Filename) != nil {
			t.FailNow()
		}
	}

	if dest.Fsys.CreateDir(dest.Filename) != nil {
		t.FailNow()
	}

	for _, destEntry := range destEntries {
		if destEntry.Fsys.Create(destEntry.Filename) != nil {
			t.FailNow()
		}
	}

	err = SyncAdd(src, dest, true, false)

	if err != nil {
		t.FailNow()
	}

	destResultEntries, err := ListDir(config.ListOpts{
		Filename:  dest,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	srcResultEntries, err := ListDir(config.ListOpts{
		Filename:  src,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range destResultEntries {
		if !slices.Contains(expectedDestEntries1, filepath.Base(
			destResultEntries[idx].Filename)) {

			t.FailNow()
		}
	}

	for idx := range srcResultEntries {
		if !slices.Contains(expectedSrcEntries1, filepath.Base(
			srcResultEntries[idx].Filename)) {

			t.FailNow()
		}
	}

	if dest.Fsys.RemoveDir(dest.Filename) != nil {
		t.FailNow()
	}

	if dest.Fsys.CreateDir(dest.Filename) != nil {
		t.FailNow()
	}

	for _, destEntry := range destEntries {
		if destEntry.Fsys.Create(destEntry.Filename) != nil {
			t.FailNow()
		}
	}

	err = SyncAdd(src, dest, false, false)

	if err != nil {
		println(err.Error())
		t.FailNow()
	}

	destResultEntries, err = ListDir(config.ListOpts{
		Filename:  dest,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	srcResultEntries, err = ListDir(config.ListOpts{
		Filename:  src,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range destResultEntries {
		if !slices.Contains(expectedDestEntries2, filepath.Base(
			destResultEntries[idx].Filename)) {

			t.FailNow()
		}
	}

	for idx := range srcResultEntries {
		if !slices.Contains(expectedSrcEntries2, filepath.Base(
			srcResultEntries[idx].Filename)) {

			t.FailNow()
		}
	}
}
