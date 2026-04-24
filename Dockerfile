FROM golang:1.26-alpine AS build
RUN apk add --no-cache gcc libc-dev
WORKDIR /go/src/app

COPY . .
RUN go test ./...
ARG version=dev
RUN go build -ldflags "-X main.Version=$version" -o /bin/blackhole

FROM alpine:3.23.4

COPY --from=build /bin/blackhole /usr/local/bin/blackhole
EXPOSE 8080
USER nobody:nobody
CMD ["/usr/local/bin/blackhole"]
