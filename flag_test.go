package main

import (
	"os"
	"testing"
)

func TestDefaultFlags(t *testing.T) {

	rollback := setInput(nil, nil)
	defer rollback()

	flags, err := ParseFlags()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if flags.Port != 8080 {
		t.Errorf("port: want 8080, got %d", flags.Port)
	}
	if flags.ProfilePort != 0 {
		t.Errorf("profile port: want 0, got %d", flags.Port)
	}
	if flags.MaxLinks != 50 {
		t.Errorf("max links: want 50, got %d", flags.MaxLinks)
	}
	if flags.MinLinks != 10 {
		t.Errorf("host: want 10, got %d", flags.MinLinks)
	}
	if flags.MaxLinkDepth != 10 {
		t.Errorf("host: want 10, got %d", flags.MaxLinkDepth)
	}
	if flags.MinLinkDepth != 1 {
		t.Errorf("host: want 1, got %d", flags.MinLinkDepth)
	}
}

func TestFlags(t *testing.T) {

	rollback := setInput([]string{"flag",
		"--port", "443",
		"--max-links", "80",
		"--min-links", "8",
		"--max-link-depth", "8",
		"--min-link-depth", "2",
	}, nil)
	defer func() { rollback() }()

	flags, err := ParseFlags()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if flags.Port != 443 {
		t.Errorf("port: want 443, got %d", flags.Port)
	}
	if flags.MaxLinks != 80 {
		t.Errorf("max links: want 80, got %d", flags.MaxLinks)
	}
	if flags.MinLinks != 8 {
		t.Errorf("host: want 8, got %d", flags.MinLinks)
	}
	if flags.MaxLinkDepth != 8 {
		t.Errorf("host: want 8, got %d", flags.MaxLinkDepth)
	}
	if flags.MinLinkDepth != 2 {
		t.Errorf("host: want 2, got %d", flags.MinLinkDepth)
	}
}

func TestFlagsFromEnvVar(t *testing.T) {

	rollback := setInput([]string{"flag"}, map[string]string{
		"BH_PORT":           "80",
		"BH_MAX_LINKS":      "7",
		"BH_MIN_LINKS":      "4",
		"BH_MAX_LINK_DEPTH": "11",
		"BH_MIN_LINK_DEPTH": "3",
	})
	defer func() { rollback() }()

	flags, err := ParseFlags()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if flags.Port != 80 {
		t.Errorf("port: want 80, got %d", flags.Port)
	}
	if flags.MaxLinks != 7 {
		t.Errorf("max links: want 7, got %d", flags.MaxLinks)
	}
	if flags.MinLinks != 4 {
		t.Errorf("host: want 4, got %d", flags.MinLinks)
	}
	if flags.MaxLinkDepth != 11 {
		t.Errorf("host: want 11, got %d", flags.MaxLinkDepth)
	}
	if flags.MinLinkDepth != 3 {
		t.Errorf("host: want 3, got %d", flags.MinLinkDepth)
	}
}

func TestFlagsOverrideEnvVar(t *testing.T) {

	rollback := setInput([]string{"flag",
		"--port", "443",
		"--max-links", "80",
		"--min-links", "8",
		"--max-link-depth", "8",
		"--min-link-depth", "2",
	}, map[string]string{
		"BH_PORT":           "80",
		"BH_MAX_LINKS":      "7",
		"BH_MIN_LINKS":      "4",
		"BH_MAX_LINK_DEPTH": "11",
		"BH_MIN_LINK_DEPTH": "3",
	})
	defer func() { rollback() }()

	flags, err := ParseFlags()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if flags.Port != 443 {
		t.Errorf("port: want 443, got %d", flags.Port)
	}
	if flags.MaxLinks != 80 {
		t.Errorf("max links: want 80, got %d", flags.MaxLinks)
	}
	if flags.MinLinks != 8 {
		t.Errorf("host: want 8, got %d", flags.MinLinks)
	}
	if flags.MaxLinkDepth != 8 {
		t.Errorf("host: want 8, got %d", flags.MaxLinkDepth)
	}
	if flags.MinLinkDepth != 2 {
		t.Errorf("host: want 2, got %d", flags.MinLinkDepth)
	}
}

func TestFlagsValidate(t *testing.T) {

	rollback := setInput([]string{"flag",
		"--port", "70000",
	}, nil)
	defer func() { rollback() }()

	if _, err := ParseFlags(); err == nil {
		t.Error("expected error")
	}
}

// --- helper functions ---

func setInput(args []string, env map[string]string) (rollback func()) {

	osArgs := os.Args
	rollback = func() {
		os.Args = osArgs
		for k := range env {
			os.Unsetenv(k)
		}
	}

	if args == nil {
		args = []string{"test"}
	}

	os.Args = args
	for k, v := range env {
		os.Setenv(k, v)
	}
	return rollback
}
