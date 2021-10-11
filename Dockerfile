FROM golang:1.17.2-alpine AS build
RUN apk add --no-cache gcc libc-dev
WORKDIR /go/src/app

COPY . .
RUN go test ./...
ARG version=dev
RUN go build -ldflags "-X main.Version=$version" -o /bin/blackhole

FROM alpine:3.14.2
RUN apk add --no-cache

COPY --from=build /bin/blackhole /usr/local/bin/blackhole
CMD ["blackhole"]
