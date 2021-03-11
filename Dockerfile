FROM golang:1.16.1-alpine AS build
RUN apk add --no-cache gcc libc-dev
WORKDIR /go/src/app

COPY flag.go link.go main.go ./
RUN go build -o /bin/blackhole


FROM alpine:3.13.2
RUN apk add --no-cache

COPY --from=build /bin/blackhole /usr/local/bin/blackhole
CMD ["blackhole"]
