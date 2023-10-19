package main

import (
	"testing"
)

func TestHTMLResume(t *testing.T) {
	str, err := GetHTMLResume()
	if err != nil {
		t.Fatal(err)
	}
	if len(str) == 0 {
		t.Fatal("empty buffer")
	}
}
