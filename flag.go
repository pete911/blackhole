package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Flags struct {
	Port         int
	ProfilePort  int
	MaxLinks     int
	MinLinks     int
	MaxLinkDepth int
	MinLinkDepth int
}

func ParseFlags() (Flags, error) {

	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	var (
		port         = f.Int("port", getIntEnv("BH_PORT", 8080), "blackhole server port")
		profilePort  = f.Int("profile-port", getIntEnv("BH_PROFILE_PORT", 0), "blackhole pprofile server port, 0 means pprof server is disabled")
		maxLinks     = f.Int("max-links", getIntEnv("BH_MAX_LINKS", 50), "max. number of links to generate")
		minLinks     = f.Int("min-links", getIntEnv("BH_MIN_LINKS", 10), "min. number of links to generate")
		maxLinkDepth = f.Int("max-link-depth", getIntEnv("BH_MAX_LINK_DEPTH", 10), "max. link depth (number of path segments)")
		minLinkDepth = f.Int("min-link-depth", getIntEnv("BH_MIN_LINK_DEPTH", 1), "min. link depth (number of path segments)")
	)
	if err := f.Parse(os.Args[1:]); err != nil {
		return Flags{}, err
	}

	//flag.Parse()
	flags := Flags{
		Port:         intValue(port),
		ProfilePort:  intValue(profilePort),
		MaxLinks:     intValue(maxLinks),
		MinLinks:     intValue(minLinks),
		MaxLinkDepth: intValue(maxLinkDepth),
		MinLinkDepth: intValue(minLinkDepth),
	}

	err := flags.validate()
	return flags, err
}

func (f Flags) validate() error {

	if f.Port < 0 || f.Port > 65535 {
		return fmt.Errorf("invalid port number: %d", f.Port)
	}
	if f.ProfilePort < 0 || f.ProfilePort > 65535 {
		return fmt.Errorf("invalid profile port number: %d", f.Port)
	}
	if f.ProfilePort == f.Port {
		return fmt.Errorf("profile port and port cannot be the same: %d", f.Port)
	}
	return nil
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

func intValue(v *int) int {

	if v == nil {
		return 0
	}
	return *v
}
