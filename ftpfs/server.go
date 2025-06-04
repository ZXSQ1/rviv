package ftpfs

import (
	"log"
	"os"
	"time"

	"goftp.io/server"
	"goftp.io/server/core"
)

const (
	testIp   = "0.0.0.0"
	testPort = 3000
	testUser = "test"
	testPass = "test"
)

var (
	testPrefix = os.TempDir()
)

func openTestServer() *server.Server {
	factory := &server.FileDriverFactory{
		RootPath: testPrefix,
		Perm:     server.NewSimplePerm("root", "root"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Auth:     &server.SimpleAuth{Name: testUser, Password: testPass},
		Port:     testPort,
		Hostname: testIp,
		Logger:   &core.DiscardLogger{},
	}

	server := server.NewServer(opts)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	time.Sleep(100 * time.Millisecond)
	return server
}
