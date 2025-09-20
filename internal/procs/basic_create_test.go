package procs

import (
	"path/filepath"
	"testing"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/localfs"
)

func TestCreateFile(t *testing.T) {
	client, err := localfs.Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	filename := config.Path{
		Filename: "test/fsdf",
		Active:   true,
		Fsys:     client,
	}

	t.Cleanup(func() {
		client.RemoveDir(filename.Filename)
		client.RemoveDir(filepath.Dir(filename.Filename))
		client.Close()
	})

	if CreateFile(filename) != nil {
		t.FailNow()
	}

	if IsRegular(filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.RemoveFile(filename.Filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.CreateDir(filename.Filename) != nil {
		t.FailNow()
	}

	if CreateFile(filename) == nil {
		t.FailNow()
	}
}

func TestCreateDir(t *testing.T) {
	client, err := localfs.Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	filename := config.Path{
		Filename: "test/sdfj",
		Active:   true,
		Fsys:     client,
	}

	t.Cleanup(func() {
		client.RemoveFile(filename.Filename)
		client.RemoveDir(filepath.Dir(filename.Filename))
		client.Close()
	})

	if CreateDir(filename) != nil {
		t.FailNow()
	}

	if IsDir(filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.RemoveDir(filename.Filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.CreateFile(filename.Filename) != nil {
		t.FailNow()
	}

	if CreateDir(filename) == nil {
		t.FailNow()
	}
}
