
VERSION = alpha

BINFILE = ./bin/rviv-$(VERSION)
SRCFILE = ./app.go

build:
	go build -o $(BINFILE) $(SRCFILE)

clean:
	rm $(BINFILE)
