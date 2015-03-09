package spc

import (
	"errors"
	"fmt"
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
			return NewNodeChar(p.Content, index), nil
		}
		return nil, errors.New(fmt.Sprintf("Parser `%c` at index %d got `%c`.", p.Content[0], index, content[index]))
	case "regexp":
		r := p.Pattern.FindStringIndex(content)
		if r == nil {
			return nil, errors.New(fmt.Sprintf("Parser `%s` at index %d does not match.", p.Pattern, index))
		}
		return NewNodeRegexp(content[r[0]:r[1]]), nil
	case "tag":
		n := NewNodeTag(p.Name)
		i := index
		for _, parser := range p.Children {
			child, err := parser.Parse(content, i)
			if err != nil {
				return nil, err
			}
			n.Add(child)
			i += child.Len()
		}
		return n, nil
	}
	return nil, nil
}
