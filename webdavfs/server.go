package webdavfs

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/net/webdav"
)

var (
	testAddr   = "127.0.0.1:3000"
	testUser   = "test"
	testPass   = "test"
	testPrefix = os.TempDir()
)

type TestServer struct {
	Server   *http.Server
	Listener net.Listener
	Addr     string
}

func (ts *TestServer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ts.Server.Shutdown(ctx)
}

func OpenTestServer() *TestServer {
	handler := &webdav.Handler{
		Prefix:     "/",
		FileSystem: webdav.Dir(filepath.Clean(testPrefix)),
		LockSystem: webdav.NewMemLS(),
	}

	authHandler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			u, p, ok := r.BasicAuth()

			if !ok || u != testUser || p != testPass {
				log.Fatalln("invalid credentials")
			}

			handler.ServeHTTP(w, r)
		},
	)

	listener, err := net.Listen("tcp", testAddr)

	if err != nil {
		log.Fatalln(err.Error())
	}

	server := &http.Server{
		Handler: authHandler,
	}

	go func() {
		if err := server.Serve(listener); err != http.ErrServerClosed {
			log.Fatalln(err.Error())
		}
	}()

	time.Sleep(100 * time.Millisecond)

	return &TestServer{
		Server:   server,
		Listener: listener,
		Addr:     listener.Addr().String(),
	}
}
