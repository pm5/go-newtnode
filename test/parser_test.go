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
	n, err := p.Parse(`"`)
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
	n, err := p.Parse(expected)
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
