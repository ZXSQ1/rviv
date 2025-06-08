
VERSION = alpha

BINFILE = ./bin/rviv-$(VERSION)
SRCFILE = ./app.go
TESTOPTS = -v -timeout 2m

build:
	go build -o $(BINFILE) $(SRCFILE)

run:
	go run $(SRCFILE)

clean:
	rm $(BINFILE)
