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

func TestLen(t *testing.T) {
	var n *spc.Node

	n = spc.NewNodeChar("c", 0)
	if n.Len() != 1 {
		t.Fatalf("Length of character node wrong. Expected 1, got %d", n.Len())
	}

	n = spc.NewNodeRegexp("hello")
	if n.Len() != 5 {
		t.Fatalf("Length of regexp node wrong. Expected 5, got %d", n.Len())
	}

	n = spc.NewNodeTag("string")
	n.Add(spc.NewNodeChar(`"`, 0))
	n.Add(spc.NewNodeRegexp("hello"))
	n.Add(spc.NewNodeChar(`"`, 5))
	if n.Len() != 7 {
		t.Fatalf("Length of tag node wrong. Expected 7, got %d", n.Len())
	}
}

func TestString(t *testing.T) {
	n := spc.NewNodeTag("string")
	n.Add(spc.NewNodeChar("a", 0))
	if n.String() != `string
  char:1:0 'a'` {
		t.Fatalf("Tag to string failed. got `%s`.", n.String())
	}
}
