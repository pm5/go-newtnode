package spc

import (
	"errors"
	"regexp"
)

type Parser struct {
	Type     string
	Content  string
	Pattern  *regexp.Regexp
	Name     string
	Children []*Parser
}

func NewParserChar(content string) *Parser {
	return &Parser{Type: "char", Content: content}
}

func NewParserRegexp(pattern string) *Parser {
	return &Parser{Type: "regexp", Pattern: regexp.MustCompile(pattern)}
}

func NewParserTag(name string, children ...*Parser) *Parser {
	return &Parser{Type: "tag", Name: name, Children: children}
}

func (p *Parser) Parse(content string, index int) (node *Node, err error) {
	switch p.Type {
	case "char":
		if content[index] == p.Content[0] {
			return NewNodeChar(p.Content, len(p.Content), 0), nil
		}
	case "regexp":
		r := p.Pattern.FindStringIndex(content)
		if r == nil {
			return nil, errors.New("Parsing failed")
		}
		return NewNodeRegexp(content[r[0]:r[1]]), nil
	case "tag":
		n := NewNodeTag(p.Name)
		return n, nil
	}
	return nil, nil
}
