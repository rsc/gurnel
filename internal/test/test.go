package test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func SetupTestDir(t *testing.T) (dir string, cleanup func()) {
	dir, err := ioutil.TempDir("", "gurnel_test")
	if err != nil {
		t.Fatalf("creating test dir: %s", err)
	}
	if err = os.Chdir(dir); err != nil {
		t.Fatalf("changing to test dir: %s", err)
	}

	cleanup = func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatalf("removing test dir: %s", err)
		}
	}
	return dir, cleanup
}

func WriteFile(t *testing.T, path, contents string) func() {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0600)
	if err != nil {
		t.Fatalf("opening file: %s", err)
	}

	if _, err = f.WriteString(contents); err != nil {
		t.Fatalf("writing file: %s", err)
	}

	return func() {
		if err := f.Close(); err != nil {
			t.Fatalf("closing file: %s", err)
		}
	}
}

func CheckErr(t *testing.T, expected string, actual error) {
	if expected == "" {
		if actual != nil {
			t.Fatalf("expected no error. got %s", actual)
		}
	} else {
		if !strings.Contains(actual.Error(), expected) {
			t.Fatalf("expected an error containing %s. got %s", expected, actual)
		}
	}
}

func CheckOutput(t *testing.T, expected []string, actual string) {
	for _, expectedOut := range expected {
		if !strings.Contains(strings.ToLower(actual), strings.ToLower(expectedOut)) {
			t.Fatalf("expected output containing %s. got %q", expectedOut, actual)
		}
	}
}

type FixedClock struct{}

func (c *FixedClock) Now() time.Time {
	return time.Date(2008, time.April, 12, 16, 0, 0, 0, time.UTC)
}
