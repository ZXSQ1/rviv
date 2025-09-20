package procs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestCopyDir(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testSrcPrefix := "test1"
	testDestPrefix := "test2"
	testSrcEntries := []string{
		"a", "b", "c", "d", "go-has-a-lot-of-err-ne-nil", "wa",
	}

	srcEntries := []config.Path{}
	srcPrefix := config.Path{
		Filename: testSrcPrefix,
		Active:   true,
		Fsys:     client,
	}

	destPrefix := config.Path{
		Filename: testDestPrefix,
		Active:   true,
		Fsys:     client,
	}

	for _, testSrcEntry := range testSrcEntries {
		srcEntries = append(srcEntries, config.Path{
			Filename: filepath.Join(srcPrefix.Filename, testSrcEntry),
			Active:   true,
			Fsys:     client,
		})
	}

	t.Cleanup(func() {
		RemoveDir(srcPrefix)
		RemoveDir(destPrefix)
		destPrefix.Fsys.Close()
	})

	if CreateDir(srcPrefix) != nil {
		t.FailNow()
	}

	if CreateDir(destPrefix) != nil {
		t.FailNow()
	}

	for _, srcEntry := range srcEntries {
		if CreateFile(srcEntry) != nil {
			t.FailNow()
		}
	}

	if err := CopyDir(srcPrefix, destPrefix, false); err != nil {
		t.FailNow()
	}

	destEntries, err := ListDir(config.ListOpts{
		Filename:  destPrefix,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range destEntries {
		if filepath.Base(srcEntries[idx].Filename) != filepath.Base(
			destEntries[idx].Filename) {

			t.FailNow()
		}

		if filepath.Dir(destEntries[idx].Filename) != filepath.Join(
			testDestPrefix, testSrcPrefix) {

			t.FailNow()
		}
	}

	if RemoveDir(destPrefix) != nil {
		t.FailNow()
	}

	if err := CopyDir(srcPrefix, destPrefix, false); err != nil {
		t.FailNow()
	}

	destEntries, err = ListDir(config.ListOpts{
		Filename:  destPrefix,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range destEntries {
		if filepath.Base(srcEntries[idx].Filename) != filepath.Base(
			destEntries[idx].Filename) {

			t.FailNow()
		}

		if filepath.Dir(destEntries[idx].Filename) != testDestPrefix {
			t.FailNow()
		}
	}
}
