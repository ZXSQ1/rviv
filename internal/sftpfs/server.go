package sftpfs

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/ZXSQ1/rviv/internal/logging"
	"github.com/gliderlabs/ssh"
	"github.com/pkg/sftp"
)

var (
	testAddr   = "127.0.0.1:3000"
	testUser   = "test"
	testPass   = "test"
	testPrefix = os.TempDir()
)

func OpenTestServer() *ssh.Server {
	sshServer := &ssh.Server{
		Addr: testAddr,

		PasswordHandler: func(ctx ssh.Context, password string) bool {
			return ctx.User() == testUser && password == testPass
		},

		SubsystemHandlers: map[string]ssh.SubsystemHandler{
			"sftp": func(sess ssh.Session) {
				logging.ReportErr(os.Chdir(testPrefix))
				server, err := sftp.NewServer(sess)

				if err != nil {
					log.Println("SFTP init error:", err)
					return
				}

				if err := server.Serve(); err == io.EOF {
					log.Println("SFTP client disconnected")
				} else if err != nil {
					log.Println("SFTP server error:", err)
				}
			},
		},
	}

	go func() {
		logging.ReportErr(
			sshServer.ListenAndServe(),
		)
	}()

	time.Sleep(100 * time.Millisecond)
	return sshServer
}
