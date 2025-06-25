package webdavfs

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *WebDavFs) Open(filename string, mode filesystem.OpenMode) (
	io.ReadWriteCloser, error) {

	if !client.IsExist(filename) {
		return nil, filesystem.ErrNotExist
	}

	switch mode {
	case filesystem.ModeWrite:
		done := make(chan error, 1)

		writer, err := client.OpenWrite(filename, done)

		if err != nil {
			return nil, err
		}

		return &File{
			mode:   mode,
			done:   done,
			writer: writer,
		}, nil

	case filesystem.ModeRead:
		reader, err := client.OpenRead(filename)

		if err != nil {
			return nil, err
		}

		return &File{
			mode:   mode,
			reader: reader,
		}, nil
	}

	return nil, os.ErrInvalid
}

func (client *WebDavFs) OpenRead(filename string) (
	io.ReadCloser, error) {

	uri := "http://" + client.connInfo.Addr + "/" + filename
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.connInfo.User, client.connInfo.Pass)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		err = resp.Body.Close()

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("open failed")
	}

	return resp.Body, nil
}

func (client *WebDavFs) OpenWrite(filename string, done chan error) (
	io.WriteCloser, error) {

	pr, pw := io.Pipe()
	uri := "http://" + client.connInfo.Addr + "/" + filename
	req, err := http.NewRequest("PUT", uri, pr)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.connInfo.User, client.connInfo.Pass)

	go func() {
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			done <- err
			return
		}

		if resp.StatusCode >= 300 {
			done <- os.ErrInvalid
			return
		}

		err = resp.Body.Close()

		if err != nil {
			done <- err
			return
		}

		done <- nil
	}()

	return pw, nil
}
