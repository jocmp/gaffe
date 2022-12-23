.PHONY: build all clean test
GOFLAGS=-mod=vendor
GOPROXY="off"
BINARY_NAME=gaffe

all: clean
clean:
	@rm -rf target/

build: clean
	@env GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o target/$(BINARY_NAME).linux
	@env GOOS=darwin go build $(GOFLAGS) -o target/$(BINARY_NAME).darwin
