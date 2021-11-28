ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

run: fmt
	go run $(ALL_GO_SOURCES) -dir=testdata

test: fmt
	go test -v -race

fmt:
	go fmt *.go

modinit:
	go mod init github.com/siongui/go-fix-google-photo-timestamp

modtidy:
	go mod tidy
