package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var absRelLinksRatio = [2]int{10, 2}

func GetLinks(flags Flags) []string {

	var links []string
	for i := 0; i < random(flags.MinLinks, flags.MaxLinks); i++ {
		links = append(links, getPath(flags))
	}
	return links
}

func getPath(flags Flags) string {

	var path []string
	for i := 0; i < random(flags.MinLinkDepth, flags.MaxLinkDepth); i++ {
		path = append(path, fmt.Sprintf("%d", rand.Intn(10000)))
	}

	// generate some relative links
	if rand.Intn(absRelLinksRatio[0]) > absRelLinksRatio[1] {
		return "/" + strings.Join(path, "/")
	}
	return strings.Join(path, "/")
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
