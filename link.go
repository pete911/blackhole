package main

import (
	"math/rand"
	"strings"
	"time"
)

const (
	linkChars      = "abcdefghijklmnopqrstuvwxyz0123456789_-"
	lenLinkChars   = len(linkChars)
	minLinkSegment = 5
	maxLinkSegment = 20
)

var absRelLinksRatio = [2]int{10, 2}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetLinks(flags Flags) []string {

	links := make([]string, random(flags.MinLinks, flags.MaxLinks))
	for i := range links {
		links[i] = getLink(flags)
	}
	return links
}

func getLink(flags Flags) string {

	segments := make([]string, random(flags.MinLinkDepth, flags.MaxLinkDepth))
	for i := range segments {
		segments[i] = getLinkSegment(minLinkSegment, maxLinkSegment)
	}

	// generate some relative and absolute links
	if rand.Intn(absRelLinksRatio[0]) > absRelLinksRatio[1] {
		return "/" + strings.Join(segments, "/")
	}
	return strings.Join(segments, "/")
}

func getLinkSegment(minLength, maxLength int) string {

	b := make([]byte, random(minLength, maxLength))
	for i := range b {
		b[i] = linkChars[random(0, lenLinkChars)]
	}
	return string(b)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
