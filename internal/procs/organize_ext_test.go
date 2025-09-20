package procs

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/localfs"
)

func TestOrganizeExt(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testOrganizedir := "b"
	testPrefix := "a"
	testEntries := []string{"a.zip", "b.zip", "d.tar", "xylem.png"}
	expectedEntries := []string{
		"Zip/a.zip", "Zip/b.zip", "Tar/d.tar", "Png/xylem.png",
	}

	entries := []config.Path{}
	prefix := config.Path{
		Filename: testPrefix,
		Active:   true,
		Fsys:     client,
	}

	organizedir := config.Path{
		Filename: testOrganizedir,
		Active:   true,
		Fsys:     client,
	}

	for _, testEntry := range testEntries {
		entries = append(entries, config.Path{
			Filename: filepath.Join(testPrefix, testEntry),
			Active:   true,
			Fsys:     client,
		})
	}

	t.Cleanup(func() {
		RemoveDir(prefix)
		RemoveDir(organizedir)
		prefix.Fsys.Close()
	})

	if CreateDir(prefix) != nil {
		t.FailNow()
	}

	if CreateDir(organizedir) != nil {
		t.FailNow()
	}

	for _, entry := range entries {
		entryDir := CopyPath(entry, filepath.Dir(entry.Filename))
		if IsExist(entryDir, CreateDir) != nil {
			t.FailNow()
		}

		if CreateFile(entry) != nil {
			t.FailNow()
		}
	}

	if OrganizeExt(entries, organizedir, false) != nil {
		t.FailNow()
	}

	slices.Sort(expectedEntries)
	organizeEntries, err := ListDir(config.ListOpts{
		Filename:  organizedir,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range organizeEntries {
		entryBase := strings.TrimLeft(strings.Replace(
			organizeEntries[idx].Filename, organizedir.Filename, "", 1,
		), "/")

		if entryBase != expectedEntries[idx] {
			t.FailNow()
		}
	}
}
