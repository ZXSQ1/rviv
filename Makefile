
VERSION = alpha

BINFILE = ./bin/rviv-$(VERSION)
SRCFILE = ./app.go
TESTOPTS = -v -timeout 2m

build:
	go build -o $(BINFILE) $(SRCFILE)

run:
	go run $(SRCFILE)

test:
	go test $(TESTOPTS) ./ftpfs
	go test $(TESTOPTS) ./localfs
	go test $(TESTOPTS) ./webdavfs
	go test $(TESTOPTS) ./sftpfs
	#go test $(TESTOPTS) ./processes

clean:
	rm $(BINFILE)
