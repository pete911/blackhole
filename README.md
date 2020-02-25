# Blackhole

Simple web server generating random links. The aim of this project is to keep bots/web crawlers occupied within 'blackhole'.

## build

 - go - `go build` ([go](https://golang.org/dl/) required)
 - docker - `docker build -t blackhole .` ([docker](https://www.docker.com/) required)

## run

 - go - `./blackhole`
 - docker - `docker run -d -p 8080:8080 --name blackhole blackhole`
 
 creates local server on port `8080`, let bots/crawlers/spiders run against it.

### optional flags

```
  -max-link-depth int
        max. link depth (number of path segments) (default 10)
  -max-links int
        max. number of links to generate (default 50)
  -min-link-depth int
        min. link depth (number of path segments) (default 1)
  -min-links int
        min. number of links to generate (default 10)
  -port int
        blackhole server port (default 8080)
```

Each option can be by env. variable as well. Env. variables have the same name as flags but are in uppercase, prefixed
with `BH_` and hyphens are replaced with underscores. e.g. flag `-max-links` can be set with `BH_MAX_LINKS` env. variable.
