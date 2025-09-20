package procs

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/expiry"
	"github.com/ZXSQ1/rviv/internal/localfs"
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
		RemoveDir(prefix)
		prefix.Fsys.Close()
	})

	if CreateDir(prefix) != nil {
		t.FailNow()
	}

	if CreateFile(archive) != nil {
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
