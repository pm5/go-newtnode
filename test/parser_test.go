package test

import (
	".."
	"testing"
)

func TestNewParserChar(t *testing.T) {
	p := newtnode.NewParserChar(`"`)
	if p == nil {
		t.Fatalf("Parser creation failed")
	}
}

func TestNewParserCharParse(t *testing.T) {
	p := newtnode.NewParserChar(`"`)
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
