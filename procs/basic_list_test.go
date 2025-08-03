package procs

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestList(t *testing.T) {
	client, err := localfs.Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	testFilename := "test"
	testEntries := []string{"a", "b/ff", "c/salad", "d/fsd", "hi", "kfs"}

	filename := config.Path{
		Filename: "test",
		Active:   true,
		Fsys:     client,
	}

	entries := []config.Path{}

	for _, testEntry := range testEntries {
		entryDir := filepath.Dir(testEntry)
		entry := config.Path{
			Filename: filepath.Join(testFilename, testEntry),
			Active:   true,
			Fsys:     client,
		}

		filename.Fsys.CreateDir(entryDir)

		if entry.Fsys.CreateFile(entry.Filename) != nil {
			t.FailNow()
		}

		entries = append(entries, entry)
	}

	t.Cleanup(func() {
		client.RemoveFile(filename.Filename)
		client.Close()
	})

	resultEntries, err := List(filename)

	if err != nil {
		t.FailNow()
	}

	for idx := range resultEntries {
		baseResultEntry := strings.Split(
			strings.Replace(
				resultEntries[idx].Filename, filename.Filename+",", "", 1,
			), string(filepath.Separator),
		)[0]

		if baseResultEntry != strings.Split(
			testEntries[idx], string(filepath.Separator))[0] {

			t.FailNow()
		}
	}
}

func TestListRecursive(t *testing.T) {
	client, err := localfs.Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	testFilename := "test"
	testEntries := []string{"a", "b", "c", "d/hj", "hi/fsd", "kfs"}

	filename := config.Path{
		Filename: "test",
		Active:   true,
		Fsys:     client,
	}

	entries := []config.Path{}

	for _, testEntry := range testEntries {
		entryDir := filepath.Dir(testEntry)
		entry := config.Path{
			Filename: filepath.Join(testFilename, testEntry),
			Active:   true,
			Fsys:     client,
		}

		filename.Fsys.CreateDir(entryDir)

		if entry.Fsys.CreateFile(entry.Filename) != nil {
			t.FailNow()
		}

		entries = append(entries, entry)
	}

	t.Cleanup(func() {
		client.RemoveFile(filename.Filename)
		client.Close()
	})

	resultEntries, err := ListRecursive(filename)

	if err != nil {
		t.FailNow()
	}

	for idx := range resultEntries {
		baseResultEntry := strings.Replace(
			resultEntries[idx].Filename, filename.Filename+",", "", 1,
		)

		if baseResultEntry != testEntries[idx] {
			t.FailNow()
		}
	}
}
