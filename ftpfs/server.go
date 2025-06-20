package ftpfs

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
	"goftp.io/server"
	"goftp.io/server/core"
)

type TestDriverFactory struct{}
type TestDriver struct{}
type TestFileInfo struct {
	stat os.FileInfo
}

func (ff *TestFileInfo) Owner() string      { return testUser }
func (ff *TestFileInfo) Group() string      { return testUser }
func (ff *TestFileInfo) IsDir() bool        { return ff.IsDir() }
func (ff *TestFileInfo) ModTime() time.Time { return ff.ModTime() }
func (ff *TestFileInfo) Mode() fs.FileMode  { return ff.Mode() }
func (ff *TestFileInfo) Name() string       { return ff.Name() }
func (ff *TestFileInfo) Size() int64        { return ff.Size() }
func (ff *TestFileInfo) Sys() any           { return ff.Sys() }

func NewTestFileInfo(stat os.FileInfo) *TestFileInfo {
	return &TestFileInfo{stat}
}

func (driver *TestDriverFactory) NewDriver() (core.Driver, error) {
	return &TestDriver{}, nil
}

func (driver *TestDriver) Stat(name string) (core.FileInfo, error) {
	os.Chdir(testPrefix)
	stat, err := os.Stat(name)

	if err != nil {
		return nil, err
	}

	return NewTestFileInfo(stat), nil
}

func (driver *TestDriver) ListDir(name string,
	fn func(core.FileInfo) error) error {

	os.Chdir(testPrefix)
	fullPath := filepath.Join(testPrefix, filepath.Clean("/"+name))
	entries, err := os.ReadDir(fullPath)

	if err != nil {
		return err
	}

	for _, entry := range entries {
		info, err := entry.Info()

		if err != nil {
			return err
		}

		if err := fn(NewTestFileInfo(info)); err != nil {
			return err
		}
	}

	return nil
}

func (driver *TestDriver) DeleteDir(name string) error {
	os.Chdir(testPrefix)
	return os.RemoveAll(name)
}

func (driver *TestDriver) DeleteFile(name string) error {
	os.Chdir(testPrefix)
	return os.Remove(name)
}

func (driver *TestDriver) Rename(oldname, newname string) error {
	os.Chdir(testPrefix)
	return os.Rename(oldname, newname)
}

func (driver *TestDriver) MakeDir(name string) error {
	os.Chdir(testPrefix)
	return os.MkdirAll(name, filesystem.PermDir)
}

func (driver *TestDriver) GetFile(name string, what int64) (int64,
	io.ReadCloser, error) {

	os.Chdir(testPrefix)
	fileObj, err := os.Open(name)
	return what, fileObj, err
}

func (driver *TestDriver) PutFile(name string, reader io.Reader, _ bool) (
	int64, error) {

	os.Chdir(testPrefix)
	stat, err := os.Stat(name)

	if err != nil {
		return -1, err
	}

	buffer := make([]byte, stat.Size())
	n, err := reader.Read(buffer)

	if err != nil {
		return -1, err
	}

	os.WriteFile(name, buffer, filesystem.PermRegular)

	return int64(n), nil
}

var (
	testAddr   = "127.0.0.1:3000"
	testIp     = "127.0.0.1"
	testPort   = 3000
	testUser   = "test"
	testPass   = "test"
	testPrefix = os.TempDir()
)

func OpenTestServer() *server.Server {
	server := server.NewServer(&core.ServerOpts{
		Factory: &TestDriverFactory{},
		Auth: &core.SimpleAuth{
			Name:     testUser,
			Password: testPass,
		},

		Hostname: testIp,
		Port:     testPort,
	})

	go func() {
		logging.ReportErr(
			server.ListenAndServe(),
		)
	}()

	return server
}
