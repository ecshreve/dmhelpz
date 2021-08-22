mod-init:
	go mod init github.com/ecshreve/dmhelpz

mod-tidy:
	go mod tidy
	
build:
	go build -o bin/dmhelpz github.com/ecshreve/dmhelpz/cmd/dmhelpz

install:
	go install -i github.com/ecshreve/dmhelpz/cmd/dmhelpz

run-only:
	bin/dmhelpz

run: build run-only

test:
	go test github.com/ecshreve/dmhelpz/...

testv:
	go test -v github.com/ecshreve/dmhelpz/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/dmhelpz/...