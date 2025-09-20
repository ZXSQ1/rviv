package procs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/localfs"
	"github.com/itchyny/timefmt-go"
)

func TestOrganizeDate(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testOrganizedir := "b"
	testPrefix := "a"
	testEntries := []string{"a", "b", "d", "e", "j", "l", "xylem"}
	testDateFmt := "%m%y"
	expectedDate := timefmt.Format(time.Now(), testDateFmt)

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

	if OrganizeDate(entries, organizedir, testDateFmt, false) != nil {
		t.FailNow()
	}

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

		if entryBase != filepath.Join(expectedDate, testEntries[idx]) {
			t.FailNow()
		}
	}
}
