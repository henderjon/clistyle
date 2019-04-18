package clistyle

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestString(t *testing.T) {
	e := Style("test this string", None)

	expected := "test this string"
	if diff := cmp.Diff(e, expected); diff != "" {
		t.Error("String(); (-got +want)", diff)
	}
}

func TestStyle(t *testing.T) {
	e := Style("test this string", Bold|YellowBG)

	// log.Fatalf("%#v", e)
	// `\x1b[;1;43mtest this string\x1b[;0m`
	expected := "\x1b[;1;43mtest this string\x1b[;0m"
	if diff := cmp.Diff(e, expected); diff != "" {
		t.Error("Style(); (-got +want)", diff)
	}
}
