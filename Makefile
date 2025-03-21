
VERSION = 0.0.1

SRC = .
BINFILE = ./bin/rviv-$(VERSION)
SRCFILE = $(SRC)/app.go
TESTOPTS = -timeout=5s -v

build:
	go build -o $(BINFILE) $(SRCFILE)

test:
	go test $(TESTOPTS) ./local/*.go
	go test $(TESTOPTS) ./ftp/*.go
	go test $(TESTOPTS) ./ssh/*.go
	go test $(TESTOPTS) ./webdav/*.go
	go test $(TESTOPTS) ./utils/*.go
	go test $(TESTOPTS) ./archive/*.go
	go test $(TESTOPTS) ./files/*.go

clean:
	rm $(BINFILE)
