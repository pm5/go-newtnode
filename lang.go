package spc

func newLangParser() (lang *TagParser) {
	lang = NewTagParser("lang")
	return
}

func NewLang(name, grammar string) (*TagParser, error) {
	return NewTagParser(name), nil
}
