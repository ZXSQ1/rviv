package procs

import (
	"path/filepath"
	"slices"
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
		Filename: testFilename,
		Active:   true,
		Fsys:     client,
	}

	for idx, testEntry := range testEntries {
		testEntries[idx] = filepath.Join(testFilename, testEntry)
		testEntry = testEntries[idx]
		entryDir := filepath.Dir(testEntry)
		filename.Fsys.CreateDir(entryDir)

		if client.CreateFile(testEntry) != nil {
			t.FailNow()
		}
	}

	t.Cleanup(func() {
		client.RemoveDir(filename.Filename)
		client.Close()
	})

	resultEntries, err := List(filename)

	if err != nil {
		t.FailNow()
	}

	rawResultEntries := []string{}

	for _, resultEntry := range resultEntries {
		rawResultEntries = append(rawResultEntries, resultEntry.Filename)
	}

	slices.Sort(rawResultEntries)
	slices.Sort(testEntries)

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
		Filename: testFilename,
		Active:   true,
		Fsys:     client,
	}

	for idx, testEntry := range testEntries {
		testEntries[idx] = filepath.Join(testFilename, testEntry)
		testEntry = testEntries[idx]
		entryDir := filepath.Dir(testEntry)
		filename.Fsys.CreateDir(entryDir)

		if client.CreateFile(testEntry) != nil {
			t.FailNow()
		}
	}

	t.Cleanup(func() {
		client.RemoveDir(filename.Filename)
		client.Close()
	})

	resultEntries, err := ListRecursive(filename)

	if err != nil {
		t.FailNow()
	}

	rawResultEntries := []string{}

	for _, resultEntry := range resultEntries {
		rawResultEntries = append(rawResultEntries, resultEntry.Filename)
	}

	slices.Sort(rawResultEntries)
	slices.Sort(testEntries)

	for idx := range resultEntries {
		baseResultEntry := strings.Replace(
			resultEntries[idx].Filename, filename.Filename+",", "", 1,
		)

		if baseResultEntry != testEntries[idx] {
			t.FailNow()
		}
	}
}
