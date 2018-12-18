GOCMD=go
GOBUILD=$(GOCMD) build
BINARYNAME=srs-playlists

all: clean build test

build:
	$(GOBUILD) -o $(BINARYNAME) cmd/main.go 

test:

clean:
	rm $(BINARYNAME)