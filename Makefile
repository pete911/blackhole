test:
	go test -v -race ./...

build: test
	go build

