package procs

import (
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestIsExist(t *testing.T) {
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

	if IsExist(filename) == nil {
		t.FailNow()
	}

	if filename.Fsys.CreateDir(filename.Filename) != nil {
		t.FailNow()
	}

	if IsExist(filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.RemoveDir(filename.Filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.CreateFile(filename.Filename) != nil {
		t.FailNow()
	}

	if IsExist(filename) != nil {
		t.FailNow()
	}
}

func TestIsDir(t *testing.T) {
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

	if IsDir(filename) == nil {
		t.FailNow()
	}

	if filename.Fsys.CreateDir(filename.Filename) != nil {
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

	if IsDir(filename) == nil {
		t.FailNow()
	}
}

func TestIsRegular(t *testing.T) {
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

	if IsRegular(filename) == nil {
		t.FailNow()
	}

	if filename.Fsys.CreateDir(filename.Filename) != nil {
		t.FailNow()
	}

	if IsRegular(filename) == nil {
		t.FailNow()
	}

	if filename.Fsys.RemoveDir(filename.Filename) != nil {
		t.FailNow()
	}

	if filename.Fsys.CreateFile(filename.Filename) != nil {
		t.FailNow()
	}

	if IsRegular(filename) != nil {
		t.FailNow()
	}
}
