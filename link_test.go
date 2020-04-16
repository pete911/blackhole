package main

import (
	"strings"
	"testing"
)

func TestGetLinks(t *testing.T) {

	flags := Flags{MinLinks: 5, MaxLinks: 20, MinLinkDepth: 1, MaxLinkDepth: 10}
	for i := 0; i < 1000; i++ {
		l := GetLinks(flags)
		if len(l) < flags.MinLinks {
			t.Errorf("%d links, expected %d or more", len(l), flags.MinLinks)
		}
		if len(l) > flags.MaxLinks {
			t.Errorf("%d links, expected %d or less", len(l), flags.MaxLinks)
		}
	}
}

func TestGetLinkSegment(t *testing.T) {

	for i := 0; i < 1000; i++ {
		s := getLinkSegment(5, 25)
		if len(s) < 5 || len(s) > 25 {
			t.Errorf("link segment is %d long, expected between 5 - 25", len(s))
		}
		if strings.Contains(s, "/") {
			t.Errorf("link segment %s, contains unexpected '/' character", s)
		}
	}
}

func TestGetLink(t *testing.T) {

	flags := Flags{MinLinkDepth: 1, MaxLinkDepth: 10}
	for i := 0; i < 1000; i++ {
		l := getLink(flags)
		depth := strings.Split(l, "/")
		if len(depth) < flags.MinLinkDepth {
			t.Errorf("link depth is %d, expected %d or more", len(depth), flags.MinLinkDepth)
		}
		if len(depth) > flags.MaxLinkDepth {
			t.Errorf("link depth is %d, expected %d or less", len(depth), flags.MaxLinkDepth)
		}
	}
}
