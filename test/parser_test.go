package test

import (
	".."
	"testing"
)

func TestParserString(t *testing.T) {
	var p spc.Parser
	var expected string

	p = spc.NewCharParser(".")
	expected = "char `.`"
	if p.String() != expected {
		t.Fatalf("Incorrect parser string. Expected `%s`, got `%s`.", expected, p)
	}

	p = spc.NewRegexpParser(`[0-9]+`)
	expected = "regexp `[0-9]+`"
	if p.String() != expected {
		t.Fatalf("Incorrect parser string. Expected `%s`, got `%s`.", expected, p)
	}

	p = spc.NewTagParser("number")
	expected = "tag `number`"
	if p.String() != expected {
		t.Fatalf("Incorrect parser string. Expected `%s`, got `%s`.", expected, p)
	}

	p = spc.NewOrParser(spc.NewCharParser("+"), spc.NewCharParser("-"))
	expected = "or"
	if p.String() != expected {
		t.Fatalf("Incorrect parser string. Expected `%s`, got `%s`.", expected, p)
	}
}

func TestNewCharParser(t *testing.T) {
	p := spc.NewCharParser(`"`)
	if p == nil {
		t.Fatalf("Parser creation failed")
	}
}

func TestCharParse(t *testing.T) {
	p := spc.NewCharParser(`o`)
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

func TestNewRegexpParser(t *testing.T) {
	p := spc.NewRegexpParser("[a-zA-Z0-9_]")
	if p == nil {
		t.Fatalf("Parser regexp creation failed")
	}
}

func TestRegexpParse(t *testing.T) {
	p := spc.NewRegexpParser(`[a-zA-Z\s]*`)
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

func TestRegexpParseFromStart(t *testing.T) {
	p := spc.NewRegexpParser(`[0-9]+`)
	sample := `the number starts from the 28th byte`
	n, err := p.Parse(sample, 0)
	if err == nil {
		t.Fatalf("Parser regexp parse failed. Expected failure, got `%s`.", n.Content)
	}
}

func TestNewTagParser(t *testing.T) {
	p := spc.NewTagParser("string",
		spc.NewCharParser(`"`),
		spc.NewRegexpParser(`[a-zA-Z\s\.]+`),
		spc.NewCharParser(`"`),
	)
	if p == nil {
		t.Fatalf("Parser tag creation failed")
	}
}

func TestTagParse(t *testing.T) {
	p := spc.NewTagParser("string",
		spc.NewCharParser(`"`),
		spc.NewRegexpParser(`[a-zA-Z\s\.,!?]+`),
		spc.NewCharParser(`"`),
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
	w := spc.NewRegexpParser(`[a-zA-Z]+`)
	s := spc.NewTagParser("sentence")
	s.Add(w, false)
	s.Add(spc.NewTagParser("more-words", spc.NewCharParser(` `), w), true)
	s.Add(spc.NewRegexpParser(`[\.!?]`), false)
	n, err := s.Parse("Do you like green eggs and ham?", 0)
	if err != nil {
		t.Fatalf("Parser add failed: %s", err)
	}
	if n.Len() != 31 {
		t.Fatalf("Parse sentence failed. Expected 31, got %d.", n.Len())
	}
}

func TestNewOrParser(t *testing.T) {
	p := spc.NewOrParser(
		spc.NewCharParser(`+`),
		spc.NewCharParser(`-`),
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

/**
 * Test the following language
 *
 * 	expression : <product> (('+' | '-') <product>)*;
 * 	product    : <value>   (('*' | '/')   <value>)*;
 * 	value      : /[0-9]+/ | '(' <expression> ')';
 * 	maths      : /^/ <expression> /$/;
 *
 */
func TestParseMath(t *testing.T) {
	expression := spc.NewTagParser("expression")
	product := spc.NewTagParser("product")
	value := spc.NewTagParser("value")
	value.Add(spc.NewOrParser(
		spc.NewRegexpParser(`[0-9]+`),
		spc.NewTagParser("parenthesized-expression",
			spc.NewCharParser(`(`),
			expression,
			spc.NewCharParser(`)`),
		),
	), false)
	product.Add(value, false)
	product.Add(spc.NewTagParser("sub-product",
		spc.NewOrParser(
			spc.NewCharParser(`*`),
			spc.NewCharParser(`/`),
		),
		value,
	), true)
	expression.Add(product, false)
	expression.Add(spc.NewTagParser("sub-expression",
		spc.NewOrParser(
			spc.NewCharParser(`+`),
			spc.NewCharParser(`-`),
		),
		product,
	), true)
	maths := spc.NewTagParser("maths",
		spc.NewRegexpParser(`^`),
		expression,
		spc.NewRegexpParser(`$`),
	)
	n, err := maths.Parse("(4*2*11+2)-5", 0)
	if err != nil {
		t.Fatalf("Math parsed failed: %s", err)
	}
	if n.Len() != 12 {
		t.Fatalf("Math parsed failed. Expected `(4*2*11+2)-5`, got `%s`.", n)
	}
}
