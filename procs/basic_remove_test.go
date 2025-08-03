package procs

import (
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestRemoveFile(t *testing.T) {
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
		client.Close()
	})

	if RemoveFile(filename) == nil {
		t.FailNow()
	}

	if filename.Fsys.CreateDir(filename.Filename) != nil {
		t.FailNow()
	}

	if RemoveFile(filename) == nil {
		t.FailNow()
	}

	if filename.Fsys.RemoveDir(filename.Filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.CreateFile(filename.Filename) != nil {
		t.FailNow()
	}

	if RemoveFile(filename) != nil {
		t.FailNow()
	}
}

func TestRemoveDir(t *testing.T) {
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
		client.Close()
	})

	if RemoveDir(filename) == nil {
		t.FailNow()
	}

	if filename.Fsys.CreateFile(filename.Filename) != nil {
		t.FailNow()
	}

	if RemoveDir(filename) == nil {
		t.FailNow()
	}

	if filename.Fsys.RemoveFile(filename.Filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.CreateDir(filename.Filename) != nil {
		t.FailNow()
	}

	if RemoveDir(filename) != nil {
		t.FailNow()
	}
}
