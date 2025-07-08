
VERSION = alpha

BINFILE = ./bin/rviv-$(VERSION)
SRCFILE = ./app.go
TESTOPTS = -v -timeout 2m

.SILENT: build test run clean

build:
	go build -o $(BINFILE) $(SRCFILE)

test:
	go test $(TESTOPTS) procs/*.go
	go test $(TESTOPTS) config/*.go
	go test $(TESTOPTS) localfs/*.go
	go test $(TESTOPTS) ftpfs/*.go
	go test $(TESTOPTS) sftpfs/*.go
	go test $(TESTOPTS) webdavfs/*.go
	go test $(TESTOPTS) compressor/*.go

run:
	go run $(SRCFILE)

clean:
	rm $(BINFILE)
