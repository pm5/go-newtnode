package test

import (
	".."
	"testing"
)

func TestNewNodeTag(t *testing.T) {
	n := newtnode.NewNodeTag("string")
	if n.Type != "tag" {
		t.Fatalf("Tag node creation failed")
	}
	if n.TagName != "string" {
		t.Fatalf("Tag name failed")
	}
}

func TestNewNodeRegex(t *testing.T) {
	n := newtnode.NewNodeRegex("cons")
	if n.Type != "regex" {
		t.Fatalf("Regex node creation failed")
	}
}

func TestNewNodeChar(t *testing.T) {
	n := newtnode.NewNodeChar(`"`, 1, 13)
	if n.Type != "char" {
		t.Fatalf("Regex node creation failed")
	}
	if n.Pos != 13 {
		t.Fatalf("Regex node creation position failed")
	}
	if n.Len != 1 {
		t.Fatalf("Regex node creation length failed")
	}
}

func TestNodeAdd(t *testing.T) {
	n := newtnode.NewNodeTag("string")
	n.Add(newtnode.NewNodeChar(`"`, 1, 0))
	n.Add(newtnode.NewNodeRegex("hello, world!"))
	n.Add(newtnode.NewNodeChar(`"`, 1, 14))
	if n.Length() != 3 {
		t.Fatalf("add child failed")
	}
	expected := `string
  char:1:0 '"'
  regex 'hello, world!'
  char:1:14 '"'`
	if n.String() != expected {
		t.Fatalf("add child, got `%s`, expected `%s`", n, expected)
	}
}

func TestNodeDelete(t *testing.T) {
	n := newtnode.NewNodeTag("string")
	n.Add(newtnode.NewNodeChar(`"`, 1, 0))
	n.Add(newtnode.NewNodeRegex("foobar"))
	n.Add(newtnode.NewNodeChar(`"`, 1, 7))
	n.Delete(1)
	expected := `string
  char:1:0 '"'
  char:1:7 '"'`
	if n.String() != expected {
		t.Fatalf("delete child failed, got `%s`, expected `%s`", n, expected)
	}
	if n.Length() != 2 {
		t.Fatalf("delete child length failed, got %d, expected 2", n.Length())
	}
}

func TestNodeRegexString(t *testing.T) {
	n := newtnode.NewNodeRegex("")
	expected := "regex"
	if n.String() != expected {
		t.Fatalf("regex node string format, got `%s`, expected `%s`", n, expected)
	}
}

func TestNodeCharString(t *testing.T) {
	n := newtnode.NewNodeChar("c", 1, 13)
	expected := "char:1:13 'c'"
	if n.String() != expected {
		t.Fatalf("char node string format, got `%s`, expected `%s`", n, expected)
	}
}
