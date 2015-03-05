package test

import (
	".."
	"testing"
)

func TestNewASTNode(t *testing.T) {
	n := gopc.NewASTNode()
	if n.Length() != 0 {
		t.Fatalf("creation failed")
	}
}

func TestASTNodeAdd(t *testing.T) {
	n := &gopc.ASTNode{Type: "tag", TagName: "string"}
	n.Add(&gopc.ASTNode{Type: "char", Pos: 0, Len: 1, Content: `"`})
	n.Add(&gopc.ASTNode{Type: "regex", Content: "hello, world!"})
	n.Add(&gopc.ASTNode{Type: "char", Pos: 14, Len: 1, Content: `"`})
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

func TestASTNodeDelete(t *testing.T) {

}

func TestASTNodeRegexPrint(t *testing.T) {
	n := &gopc.ASTNode{Type: "regex", Len: 0, Pos: 0}
	expected := "regex"
	if n.String() != expected {
		t.Fatalf("regex node string format, got `%s`, expected `%s`", n, expected)
	}
}

func TestASTNodeCharPrint(t *testing.T) {
	n := &gopc.ASTNode{Type: "char", Len: 1, Pos: 13, Content: "c"}
	expected := "char:1:13 'c'"
	if n.String() != expected {
		t.Fatalf("char node string format, got `%s`, expected `%s`", n, expected)
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
		t.Fatalf("char node string format, got `%s`, expected `%s`", n, expected)
	}
}
