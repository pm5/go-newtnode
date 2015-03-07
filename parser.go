package spc

import (
	"errors"
	"regexp"
)

type Parser struct {
	Type    string
	Content string
	Pattern *regexp.Regexp
}

func NewParserChar(content string) *Parser {
	return &Parser{Type: "char", Content: content}
}

func NewParserRegexp(pattern string) *Parser {
	return &Parser{Type: "regexp", Pattern: regexp.MustCompile(pattern)}
}

func (p *Parser) Parse(content string) (*Node, error) {
	switch p.Type {
	case "char":
		if content == p.Content {
			return NewNodeChar(content, len(content), 0), nil
		}
	case "regexp":
		r := p.Pattern.FindStringIndex(content)
		if r == nil {
			return nil, errors.New("Parsing failed")
		}
		return NewNodeRegexp(content[r[0]:r[1]]), nil
	}
	return nil, nil
}
