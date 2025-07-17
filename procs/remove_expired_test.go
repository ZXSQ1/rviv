package procs

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/expiry"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestRemoveExpired(t *testing.T) {
	client, err := localfs.Init(os.TempDir())

	if err != nil {
		t.FailNow()
	}

	testPrefix := "test"
	testArchive := "archive[%d]"

	prefix := config.Path{
		Filename: testPrefix,
		Active:   true,
		Fsys:     client,
	}

	archive := config.Path{
		Filename: filepath.Join(
			testPrefix, config.StdPath(testArchive),
		),

		Active: true,
		Fsys:   client,
	}

	t.Cleanup(func() {
		prefix.Fsys.RemoveDir(prefix.Filename)
		prefix.Fsys.Close()
	})

	if prefix.Fsys.CreateDir(prefix.Filename) != nil {
		t.FailNow()
	}

	if archive.Fsys.Create(archive.Filename) != nil {
		t.FailNow()
	}

	testExpiry, err := expiry.ParseExpiry("4s")

	if err != nil {
		t.FailNow()
	}

	if RemoveExpired(prefix, testArchive, testExpiry, false) != nil {
		t.FailNow()
	}

	entries, err := ListDir(config.ListOpts{
		Filename:  prefix,
		Recursive: false,
		Verbose:   false,
	})

	if err != nil {
		t.FailNow()
	}

	if len(entries) < 1 {
		t.FailNow()
	}

	time.Sleep(4 * time.Second)
	testExpiry, err = expiry.ParseExpiry("4s")

	if err != nil {
		t.FailNow()
	}

	if RemoveExpired(prefix, testArchive, testExpiry, false) != nil {
		t.FailNow()
	}

	entries, err = ListDir(config.ListOpts{
		Filename:  prefix,
		Recursive: false,
		Verbose:   false,
	})

	if err != nil {
		t.FailNow()
	}

	if len(entries) > 0 {
		t.FailNow()
	}
}
