package test

import (
	".."
	"testing"
)

func TestASTNodeRegexPrint(t *testing.T) {
	n := &gopc.ASTNode{Type: "regex", Len: 0, Pos: 0}
	expected := "regex"
	if n.String() != expected {
		t.Errorf("regex node string format, got `%s`, expected `%s`", n, expected)
	}
}

func TestASTNodeCharPrint(t *testing.T) {
	n := &gopc.ASTNode{Type: "char", Len: 1, Pos: 13, Content: "c"}
	expected := "char:1:13 'c'"
	if n.String() != expected {
		t.Errorf("char node string format, got `%s`, expected `%s`", n, expected)
	}
}

func TestASTNodeTagPrint(t *testing.T) {
	n := &gopc.ASTNode{
		Type:    "tag",
		TagName: "string",
		Children: []*gopc.ASTNode{
			&gopc.ASTNode{Type: "char", Len: 1, Pos: 0, Content: `"`},
			&gopc.ASTNode{Type: "regex", Content: "hello, world!"},
			&gopc.ASTNode{Type: "char", Len: 1, Pos: 14, Content: `"`},
		},
	}
	expected := `string
  char:1:0 '"'
  regex 'hello, world!'
  char:1:14 '"'`
	if n.String() != expected {
		t.Errorf("char node string format, got `%s`, expected `%s`", n, expected)
	}
}

func TestASTNodeAdd(t *testing.T) {
	a := &gopc.ASTNode{Type: "tag", TagName: "string"}
	a.Add(&gopc.ASTNode{Type: "char", Pos: 0, Len: 1, Content: `"`})
}
