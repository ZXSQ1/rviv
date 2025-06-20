package ftpfs

import (
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ZXSQ1/rviv/logging"
	ftpserver "github.com/fclairamb/ftpserverlib"
	"github.com/spf13/afero"
)

type TestDriver struct {
	BasePath string
}

type TestClientDriver struct {
	BasePath string
}

var (
	testAddr   = "127.0.0.1:3000"
	testUser   = "test"
	testPass   = "test"
	testPrefix = os.TempDir()
)

func (driver *TestDriver) GetSettings() (*ftpserver.Settings, error) {
	return &ftpserver.Settings{
		ListenAddr: testAddr,
		PassiveTransferPortRange: &ftpserver.PortRange{
			Start: 3001,
			End:   3012,
		},
	}, nil
}

func (driver *TestDriver) ClientConnected(cc ftpserver.ClientContext) (
	string, error) {

	return "", nil
}

func (driver *TestDriver) ClientDisconnected(cc ftpserver.ClientContext) {}
func (driver *TestDriver) AuthUser(cc ftpserver.ClientContext,
	user, pass string) (ftpserver.ClientDriver, error) {

	if user == testUser && pass == testPass {
		return &TestClientDriver{
			BasePath: driver.BasePath,
		}, nil
	}

	return nil, fmt.Errorf("invalid credentials")
}

func (driver *TestDriver) GetTLSConfig() (*tls.Config, error) {
	return nil, nil
}

func (driver *TestClientDriver) Chmod(name string, mode os.FileMode) error {
	return os.Chmod(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	), mode)
}

func (driver *TestClientDriver) Chown(name string, uid, gid int) error {
	return os.Chown(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	), uid, gid)
}

func (driver *TestClientDriver) Chtimes(name string, atime,
	mtime time.Time) error {

	return os.Chtimes(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	), atime, mtime)
}

func (driver *TestClientDriver) Create(name string) (afero.File, error) {
	return os.Create(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	))
}

func (driver *TestClientDriver) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	), perm)
}

func (driver *TestClientDriver) MkdirAll(name string, perm os.FileMode) error {
	return os.MkdirAll(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	), perm)
}

func (driver *TestClientDriver) Name() string {
	return "Local FTP Server"
}

func (driver *TestClientDriver) Open(name string) (afero.File, error) {
	return os.Open(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	))
}

func (driver *TestClientDriver) OpenFile(name string, flag int,
	perm os.FileMode) (afero.File, error) {

	return os.OpenFile(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	), flag, perm)
}

func (driver *TestClientDriver) Remove(name string) error {
	return os.Remove(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	))
}

func (driver *TestClientDriver) RemoveAll(name string) error {
	return os.RemoveAll(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	))
}

func (driver *TestClientDriver) Rename(oldname, newname string) error {
	return os.Rename(
		filepath.Join(driver.BasePath, filepath.Clean("/"+oldname)),
		filepath.Join(driver.BasePath, filepath.Clean("/"+newname)),
	)
}

func (driver *TestClientDriver) Stat(name string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(
		driver.BasePath,
		filepath.Clean("/"+name),
	))
}

func OpenTestServer() *ftpserver.FtpServer {
	server := ftpserver.NewFtpServer(&TestDriver{BasePath: testPrefix})

	go func() {
		logging.ReportErr(
			server.ListenAndServe(),
		)
	}()

	time.Sleep(500 * time.Millisecond)

	return server
}
