package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	port         = flag.Int("port", getIntEnv("BH_PORT", 8080), "blackhole server port")
	maxLinks     = flag.Int("max-links", getIntEnv("BH_MAX_LINKS", 50), "max. number of links to generate")
	minLinks     = flag.Int("min-links", getIntEnv("BH_MIN_LINKS", 10), "min. number of links to generate")
	maxLinkDepth = flag.Int("max-link-depth", getIntEnv("BH_MAX_LINK_DEPTH", 10), "max. link depth (number of path segments)")
	minLinkDepth = flag.Int("min-link-depth", getIntEnv("BH_MIN_LINK_DEPTH", 1), "min. link depth (number of path segments)")
)

type Flags struct {
	Port         int
	MaxLinks     int
	MinLinks     int
	MaxLinkDepth int
	MinLinkDepth int
}

func ParseFlags() Flags {

	flag.Parse()
	f := Flags{
		Port:         intValue(port),
		MaxLinks:     intValue(maxLinks),
		MinLinks:     intValue(minLinks),
		MaxLinkDepth: intValue(maxLinkDepth),
		MinLinkDepth: intValue(minLinkDepth),
	}

	if err := f.validate(); err != nil {
		failFlags(err)
	}
	return f
}

func (f Flags) validate() error {

	if f.Port < 0 || f.Port > 65535 {
		return fmt.Errorf("invalid port number: %d", f.Port)
	}
	return nil
}

func intValue(v *int) int {

	if v == nil {
		return 0
	}
	return *v
}

func getIntEnv(envName string, defaultValue int) int {

	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}

	if intValue, err := strconv.Atoi(env); err == nil {
		return intValue
	}
	return defaultValue
}

func failFlags(err error) {

	flag.CommandLine.SetOutput(os.Stderr)
	fmt.Fprintln(flag.CommandLine.Output(), err.Error())
	flag.Usage()
	os.Exit(2)
}
