
VERSION = alpha

BINFILE = ./bin/rviv-$(VERSION)
SRCFILE = ./cmd/app.go
TESTOPTS = -v -timeout 2m

.SILENT: build test run clean

build:
	go build -o $(BINFILE) $(SRCFILE)

test-fs:
	go test $(TESTOPTS) ./internal/localfs/*.go
	go test $(TESTOPTS) ./internal/ftpfs/*.go
	go test $(TESTOPTS) ./internal/sftpfs/*.go
	go test $(TESTOPTS) ./internal/webdavfs/*.go

test-utils:
	go test $(TESTOPTS) ./internal/expiry/*.go

test-config:
	go test $(TESTOPTS) ./internal/config/*.go

test-procs:
	go test $(TESTOPTS) ./internal/procs/*.go

test: test-fs test-utils test-config test-procs

run:
	go run $(SRCFILE)

clean:
	rm $(BINFILE)
