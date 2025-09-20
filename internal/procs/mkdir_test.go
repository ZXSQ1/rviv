package procs

import (
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestMkdir(t *testing.T) {
	client, err := localfs.Init("/tmp")

	if err != nil {
		t.FailNow()
	}

	testPrefix := "a"
	testDirs := []string{
		testPrefix + "/asdfsd/fsd", testPrefix + "/b",
		testPrefix + "/fsd/fsdf/a/s",
	}

	dirs := []config.Path{}
	prefix := config.Path{
		Filename: testPrefix,
		Active:   true,
		Fsys:     client,
	}

	for _, testDir := range testDirs {
		dirs = append(dirs, config.Path{
			Filename: testDir,
			Active:   true,
			Fsys:     client,
		})
	}

	t.Cleanup(func() {
		prefix.Fsys.RemoveDir(prefix.Filename)
		prefix.Fsys.Close()
	})

	err = Mkdir(config.MkdirOpts{
		Filenames: dirs,
		Parent:    false,
	})

	if err == nil {
		t.FailNow()
	}

	paths, err := ListDir(config.ListOpts{
		Filename:  prefix,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range paths {
		if paths[idx].Filename != testDirs[idx] {
			t.FailNow()
		}
	}

	if prefix.Fsys.RemoveDir(prefix.Filename) != nil {
		t.FailNow()
	}

	err = Mkdir(config.MkdirOpts{
		Filenames: dirs,
		Parent:    true,
	})

	if err != nil {
		t.FailNow()
	}

	paths, err = ListDir(config.ListOpts{
		Filename:  prefix,
		Recursive: true,
	})

	if err != nil {
		t.FailNow()
	}

	for idx := range paths {
		if paths[idx].Filename != testDirs[idx] {
			t.FailNow()
		}
	}
}
