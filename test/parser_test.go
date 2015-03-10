package test

import (
	".."
	"testing"
)

func TestNewParserChar(t *testing.T) {
	p := spc.NewParserChar(`"`)
	if p == nil {
		t.Fatalf("Parser creation failed")
	}
}

func TestCharParse(t *testing.T) {
	p := spc.NewParserChar(`o`)
	n, err := p.Parse(`"hello"`, 5)
	if err != nil {
		t.Fatalf("Parse char failed: %s", err)
	}
	if n == nil {
		t.Fatalf("Parse char failed. Expected `o`, got nil.")
	}
	if n.Type != "char" {
		t.Fatalf("Parser char parse failed. Expected `char`, got `%s`", n.Type)
	}
	if n.Content != `o` {
		t.Fatalf("Parser char parse failed. Expected `\"`, got `%s`", n.Content)
	}
	if n.Len() != 1 {
		t.Fatalf("Parser char parse failed. Expected `1`, got `%d`", n.Len())
	}
	if n.Pos != 5 {
		t.Fatalf("Parser char parse failed. Expected `1`, got `%d`", n.Pos)
	}
}

func TestNewParserRegexp(t *testing.T) {
	p := spc.NewParserRegexp("[a-zA-Z0-9_]")
	if p == nil {
		t.Fatalf("Parser regexp creation failed")
	}
}

func TestRegexpParse(t *testing.T) {
	p := spc.NewParserRegexp(`[a-zA-Z\s]*`)
	expected := `just what do you think youre doing`
	n, err := p.Parse(expected+`, dave?`, 0)
	if err != nil || n == nil {
		t.Fatalf("Parser regexp parse failed")
	}
	if n.Type != "regexp" {
		t.Fatalf("Parser regexp parse failed. Expected `regexp`, got `%s`", n.Type)
	}
	if n.Content != expected {
		t.Fatalf("Parser regexp parse failed. Expected `%s`, got `%s`", expected, n.Content)
	}
}

func TestNewParserTag(t *testing.T) {
	p := spc.NewParserTag("string",
		spc.NewParserChar(`"`),
		spc.NewParserRegexp(`[a-zA-Z\s\.]+`),
		spc.NewParserChar(`"`),
	)
	if p == nil {
		t.Fatalf("Parser tag creation failed")
	}
}

func TestTagParse(t *testing.T) {
	p := spc.NewParserTag("string",
		spc.NewParserChar(`"`),
		spc.NewParserRegexp(`[a-zA-Z\s\.,!?]+`),
		spc.NewParserChar(`"`),
	)
	n, err := p.Parse("\"hello, world!\"", 0)
	if err != nil {
		t.Fatalf("Parser tag parse failed: %s", err)
	}
	if n.Type != "tag" {
		t.Fatalf("Parser tag parse failed. Expected `tag`, got `%s`", n.Type)
	}
	if n.Length() != 3 {
		t.Fatalf("Parser tag parse failed. Expected 3, got %d", n.Length())
	}
	if n.Children[1].Content != "hello, world!" {
		t.Fatalf("Parser tag parse failed. Expected `hello, world!`, got `%s`", n.Children[1].Content)
	}
}

func TestTagAdd(t *testing.T) {
	w := spc.NewParserRegexp(`[a-zA-Z]+`)
	s := spc.NewParserTag("sentence")
	s.Add(w, true)
	s.Add(spc.NewParserRegexp(`[\.!?]`), false)
	n, err := s.Parse("Do you like green eggs and ham?", 0)
	if err != nil {
		t.Fatalf("Parser add failed: %s", err)
	}
	if n.Len() != 31 {
		t.Fatalf("Parse sentence failed. Expected 31, got %d.", n.Len())
	}
}

func TestNewParserOr(t *testing.T) {
	p := spc.NewParserOr(
		spc.NewParserChar(`+`),
		spc.NewParserChar(`-`),
	)
	var n *spc.Node
	var err error
	n, err = p.Parse("+13", 0)
	if err != nil {
		t.Fatalf("OR parse failed: %s", err)
	}
	if n.Content != "+" {
		t.Fatalf("OR parse failed. Expected `+`, got `%s`.", n.Content)
	}
	n, err = p.Parse("-200", 0)
	if err != nil {
		t.Fatalf("OR parse failed: %s", err)
	}
	if n.Content != "-" {
		t.Fatalf("OR parse failed. Expected `-`, got `%s`.", n.Content)
	}
}
