IMAGE := pete911/blackhole
VERSION ?= dev

test:
	go test -v -race ./...

build: test
	go build

image:
	docker build -t ${IMAGE}:${VERSION} .
	docker tag ${IMAGE}:${VERSION} ${IMAGE}:latest

push-image:
	docker push ${IMAGE}:${VERSION}
	docker push ${IMAGE}:latest