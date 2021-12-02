.PHONY: clean build

DIST = ./bin
ENTRY = ./cmd

clean:
	rm -rf ${DIST}

build-unix: clean
	go build -ldflags "-s -w" -o ${DIST}/unix/downloader ${ENTRY}

.DEFAULT_GOAL := build-unix
