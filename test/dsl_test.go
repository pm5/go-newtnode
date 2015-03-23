package test

import (
	".."
	"testing"
)

func TestParseLang(t *testing.T) {
	p := spc.NewLang("math", `
	expr    : <product> (('+' | '-') <product>)*;
	product : <value> (('*' | '/') <value)*;
	value   : /[0-9]+/ | '(' <expr> ')';
	math	: /^/ <expr> /$/;
		`)
	if err != nil {
		t.Fatalf("Math parsed failed: %s", err)
	}
	if n.Len() != 12 {
		t.Fatalf("Math parsed failed. Expected `(4*2*11+2)-5`, got `%s`.", n)
	}
}
