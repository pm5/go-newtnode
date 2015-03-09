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

func TestGetLen(t *testing.T) {
	var n *spc.Node

	n = spc.NewNodeChar("c", 1, 0)
	if n.GetLen() != 1 {
		t.Fatalf("Length of character node wrong. Expected 1, got %d", n.GetLen())
	}

	n = spc.NewNodeRegexp("hello")
	if n.GetLen() != 5 {
		t.Fatalf("Length of regexp node wrong. Expected 5, got %d", n.GetLen())
	}

	n = spc.NewNodeTag("string")
	n.Add(spc.NewNodeChar(`"`, 1, 0))
	n.Add(spc.NewNodeRegexp("hello"))
	n.Add(spc.NewNodeChar(`"`, 1, 5))
	if n.GetLen() != 7 {
		t.Fatalf("Length of tag node wrong. Expected 7, got %d", n.GetLen())
	}
}

//func TestLengthChar(t *testing.T) {
//n := spc.NewNodeChar(".", 1, 0)
//if n.Length() != 1 {
//t.Fatalf("Node length wrong. Expected 1, got %d", n.Length())
//}
//}
