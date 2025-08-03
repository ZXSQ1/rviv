package procs

import (
	"testing"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/localfs"
)

func TestOpen(t *testing.T) {
	client, err := localfs.Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	testContent := "the fox jumps over the dog"
	filename := config.Path{
		Filename: "test",
		Active:   true,
		Fsys:     client,
	}

	t.Cleanup(func() {
		client.RemoveFile(filename.Filename)
		client.Close()
	})

	if filename.Fsys.CreateFile(filename.Filename) != nil {
		t.FailNow()
	}

	fileObj, err := Open(filename, filesystem.ModeWrite)

	if err != nil {
		t.FailNow()
	}

	if _, err := fileObj.Write([]byte(testContent)); err != nil {
		t.FailNow()
	}

	if _, err := fileObj.Read([]byte{}); err == nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	fileObj, err = Open(filename, filesystem.ModeRead)

	if err != nil {
		t.FailNow()
	}

	if _, err := fileObj.Write([]byte{}); err == nil {
		t.FailNow()
	}

	buffer := make([]byte, len(testContent))

	if _, err := fileObj.Read(buffer); err == nil {
		t.FailNow()
	}

	if fileObj.Close() != nil {
		t.FailNow()
	}

	if string(buffer) != testContent {
		t.FailNow()
	}
}
