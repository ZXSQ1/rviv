package procs

import (
	"testing"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/localfs"
)

func TestStat(t *testing.T) {
	client, err := localfs.Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	filename := config.Path{
		Filename: "test",
		Active:   true,
		Fsys:     client,
	}

	t.Cleanup(func() {
		client.RemoveFile(filename.Filename)
		client.Close()
	})

	if _, err := Stat(filename); err == nil {
		t.FailNow()
	}

	if filename.Fsys.CreateDir(filename.Filename) != nil {
		t.FailNow()
	}

	if _, err := Stat(filename); err != nil {
		t.FailNow()
	}
}
