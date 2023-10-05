VERSION=`git describe --tags --match v[0-9]* 2> /dev/null`

.PHONY: build test clean

build:
	go build ./...

test:
	go test ./...

install:
	go install

dist: dist-linux
	rm -f gophoria

dist-linux: dist-linux-amd64 dist-linux-arm64 dist-linux-i386
dist-linux-amd64:
	rm -f gophoria && GOARCH=amd64 GOOS=linux go build ./... && tar -czvf "gophoria-${VERSION}-linux-amd64.tar.gz" gophoria
dist-linux-arm64:
	rm -f gophoria && GOARCH=arm64 GOOS=linux go build ./... && tar -czvf "gophoria-${VERSION}-linux-arm64.tar.gz" gophoria

clean:
	rm -rf gophoria*
