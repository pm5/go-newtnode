package spc

func newLangParser() (lang *TagParser) {
	lang = NewTagParser("lang")
	rule := NewTagParser("rule")
	name := NewRegexpParser(`[a-zA-Z\-_]+`)
	rule.Add(name, false)
	rule.Add(NewCharParser(`:`), false)
	expr := NewTagParser("expr")
	rule.Add(expr, false)
	rule.Add(NewCharParser(`;`), false)
	lang.Add(rule, true)
	return
}

func NewLang(name, grammar string) (*TagParser, error) {
	return NewTagParser(name), nil
}
