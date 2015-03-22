package test

import (
	".."
	"testing"
)

func TestParseLang(t *testing.T) {
	_ = spc.NewLang("math", `
	expr    : <product> (('+' | '-') <product>)*;
	product : <value> (('*' | '/') <value)*;
	value   : /[0-9]+/ | '(' <expr> ')';
	math	: /^/ <expr> /$/;
		`)
}
