package test

import (
	".."
	"testing"
)

func TestNewNodeTag(t *testing.T) {
	n := spc.NewNodeTag("string")
	if n.Type != "tag" {
		t.Fatalf("Tag node creation failed")
	}
	if n.TagName != "string" {
		t.Fatalf("Tag name failed")
	}
}
