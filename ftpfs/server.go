package ftpfs

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/ZXSQ1/rviv/logging"
	ftpserver "github.com/fclairamb/ftpserverlib"
	"github.com/spf13/afero"
)

type TestDriver struct{}
type TestClientDriver struct{}

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
		return &TestClientDriver{}, nil
	}

	return nil, fmt.Errorf("invalid credentials")
}

func (driver *TestDriver) GetTLSConfig() (*tls.Config, error) {
	return nil, nil
}

func (driver *TestClientDriver) Chmod(name string, mode os.FileMode) error {
	os.Chdir(testPrefix)
	return os.Chmod(name, mode)
}

func (driver *TestClientDriver) Chown(name string, uid, gid int) error {
	os.Chdir(testPrefix)
	return os.Chown(name, uid, gid)
}

func (driver *TestClientDriver) Chtimes(name string, atime,
	mtime time.Time) error {

	os.Chdir(testPrefix)
	return os.Chtimes(name, atime, mtime)
}

func (driver *TestClientDriver) Create(name string) (afero.File, error) {
	os.Chdir(testPrefix)
	return os.Create(name)
}

func (driver *TestClientDriver) Mkdir(name string, perm os.FileMode) error {
	os.Chdir(testPrefix)
	return os.Mkdir(name, perm)
}

func (driver *TestClientDriver) MkdirAll(name string, perm os.FileMode) error {
	os.Chdir(testPrefix)
	return os.MkdirAll(name, perm)
}

func (driver *TestClientDriver) Name() string {
	return "Local FTP Server"
}

func (driver *TestClientDriver) Open(name string) (afero.File, error) {
	os.Chdir(testPrefix)
	return os.Open(name)
}

func (driver *TestClientDriver) OpenFile(name string, flag int,
	perm os.FileMode) (afero.File, error) {

	os.Chdir(testPrefix)
	return os.OpenFile(name, flag, perm)
}

func (driver *TestClientDriver) Remove(name string) error {
	os.Chdir(testPrefix)
	return os.Remove(name)
}

func (driver *TestClientDriver) RemoveAll(name string) error {
	os.Chdir(testPrefix)
	return os.RemoveAll(name)
}

func (driver *TestClientDriver) Rename(oldname, newname string) error {
	os.Chdir(testPrefix)
	return os.Rename(oldname, newname)
}

func (driver *TestClientDriver) Stat(name string) (os.FileInfo, error) {
	os.Chdir(testPrefix)
	return os.Stat(name)
}

func OpenTestServer() *ftpserver.FtpServer {
	server := ftpserver.NewFtpServer(&TestDriver{})

	go func() {
		logging.ReportErr(
			server.ListenAndServe(),
		)
	}()

	return server
}
