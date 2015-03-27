package test

import (
	".."
	"testing"
)

func TestNewLang(t *testing.T) {
	p, err := spc.NewLang("math", `
	expr    : <product> (('+' | '-') <product>)*;
	product : <value> (('*' | '/') <value>)*;
	value   : /[0-9]+/ | '(' <expr> ')';
	math	: /^/ <expr> /$/;
		`)
	if err != nil || p == nil {
		t.Fatalf("Lang parser creation failed: %s", err)
	}
}
