FROM golang:1.18rc1-alpine AS build
RUN apk add --no-cache gcc libc-dev
WORKDIR /go/src/app

COPY . .
RUN go test ./...
ARG version=dev
RUN go build -ldflags "-X main.Version=$version" -o /bin/blackhole

FROM alpine:3.15.0
RUN apk add --no-cache

COPY --from=build /bin/blackhole /usr/local/bin/blackhole
CMD ["blackhole"]
