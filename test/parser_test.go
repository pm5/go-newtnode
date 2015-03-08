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

func TestParserCharParse(t *testing.T) {
	p := spc.NewParserChar(`"`)
	n, left, err := p.Parse(`"`)
	if n == nil || err != nil {
		t.Fatalf("Parser char parse failed")
	}
	if n.Type != "char" {
		t.Fatalf("Parser char parse failed. Expected `char`, got `%s`", n.Type)
	}
	if n.Content != `"` {
		t.Fatalf("Parser char parse failed. Expected `\"`, got `%s`", n.Content)
	}
	if n.Len != 1 {
		t.Fatalf("Parser char parse failed. Expected `1`, got `%d`", n.Len)
	}
	if n.Pos != 0 {
		t.Fatalf("Parser char parse failed. Expected `1`, got `%d`", n.Pos)
	}
	if left != "" {
		t.Fatalf("Parser char parse failed. Expected ``, got `%s`", left)
	}
}

func TestNewParserRegexp(t *testing.T) {
	p := spc.NewParserRegexp("[a-zA-Z0-9_]")
	if p == nil {
		t.Fatalf("Parser regexp creation failed")
	}
}

func TestParserRegexpParse(t *testing.T) {
	p := spc.NewParserRegexp(`[a-zA-Z\s]*`)
	expected := `just what do you think youre doing dave`
	n, _, err := p.Parse(expected)
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
		spc.NewParserRegexp(`[a-zA-Z\s\.]*`),
		spc.NewParserChar(`"`),
	)
	if p == nil {
		t.Fatalf("Parser tag creation failed")
	}
}

func TestParserTagParse(t *testing.T) {
	p := spc.NewParserTag("string",
		spc.NewParserChar(`"`),
		spc.NewParserRegexp(`[a-zA-Z\s\.!?]*`),
		spc.NewParserChar(`"`),
	)
	n, _, err := p.Parse("hello, world!")
	if err != nil || n == nil {
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
