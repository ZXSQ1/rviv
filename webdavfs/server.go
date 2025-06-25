package webdavfs

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
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
	server   *http.Server
	listener net.Listener
	addr     string
}

func (ts *TestServer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ts.server.Shutdown(ctx)
}

func OpenTestServer() *TestServer {
	handler := &webdav.Handler{
		Prefix:     "/",
		FileSystem: webdav.Dir(testPrefix),
		LockSystem: webdav.NewMemLS(),
	}

	authHandler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			u, p, ok := r.BasicAuth()

			if !ok || u != testUser || p != testPass {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)

				return
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
		server:   server,
		listener: listener,
		addr:     listener.Addr().String(),
	}
}
