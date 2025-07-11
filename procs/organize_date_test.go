package procs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
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
		prefix.Fsys.RemoveDir(prefix.Filename)
		organizedir.Fsys.RemoveDir(organizedir.Filename)
		prefix.Fsys.Close()
	})

	if prefix.Fsys.CreateDir(prefix.Filename) != nil {
		t.FailNow()
	}

	if organizedir.Fsys.CreateDir(organizedir.Filename) != nil {
		t.FailNow()
	}

	for _, entry := range entries {
		if !entry.Fsys.IsExist(filepath.Dir(entry.Filename)) {
			if entry.Fsys.CreateDir(filepath.Dir(entry.Filename)) != nil {
				t.FailNow()
			}
		}

		if entry.Fsys.Create(entry.Filename) != nil {
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
