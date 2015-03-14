package spc

import (
	"errors"
	"fmt"
	"github.com/tj/go-debug"
	"regexp"
)

var d = debug.Debug("spc:parser")

type Parser interface {
	Parse(content string, index int) (node *Node, err error)
	String() string
}

type IndefiniteParser struct {
	Parser     Parser
	Indefinite bool
}

type CharParser struct {
	Content string
}

type RegexpParser struct {
	Pattern *regexp.Regexp
}

type TagParser struct {
	Name     string
	Children []*IndefiniteParser
}

type OrParser struct {
	Children []*IndefiniteParser
	a, b     Parser
}

func NewParserChar(content string) *CharParser {
	return &CharParser{Content: content}
}

func NewParserRegexp(pattern string) *RegexpParser {
	return &RegexpParser{Pattern: regexp.MustCompile("^" + pattern)}
}

func NewParserTag(name string, children ...Parser) *TagParser {
	p := TagParser{Name: name}
	for _, child := range children {
		p.Add(child, false)
	}
	return &p
}

func NewParserOr(a, b Parser) *OrParser {
	return &OrParser{a: a, b: b}
}

func (p CharParser) String() string {
	return "char `" + p.Content + "`"
}

func (p RegexpParser) String() string {
	return "regexp `" + p.Pattern.String()[1:] + "`"
}

func (p TagParser) String() string {
	return "tag `" + p.Name + "`"
}

func (p OrParser) String() string {
	return "or"
}

func (p CharParser) Parse(content string, index int) (node *Node, err error) {
	d("Parse `%s` with %s at %d", content, p, index)
	if index >= len(content) {
		return nil, errors.New(fmt.Sprintf("Index exceeds parsed string length %d.", len(content)))
	}
	if content[index] == p.Content[0] {
		return NewNodeChar(p.Content, index), nil
	}
	return nil, errors.New(fmt.Sprintf("Parser `%c` at index %d got `%c`.", p.Content[0], index, content[index]))
}

func (p RegexpParser) Parse(content string, index int) (node *Node, err error) {
	d("Parse `%s` with %s at %d", content, p, index)
	if index > len(content) {
		return nil, errors.New(fmt.Sprintf("Index exceeds parsed string length %d.", len(content)))
	}
	r := p.Pattern.FindStringIndex(content[index:])
	if r == nil {
		return nil, errors.New(fmt.Sprintf("Parser `%s` at index %d does not match.", p.Pattern, index))
	}
	return NewNodeRegexp(content[index+r[0] : index+r[1]]), nil
}

func (p TagParser) Parse(content string, index int) (node *Node, err error) {
	d("Parse `%s` with %s at %d", content, p, index)
	if index > len(content) {
		return nil, errors.New(fmt.Sprintf("Index exceeds parsed string length %d.", len(content)))
	}
	n := NewNodeTag(p.Name)
	i := index
	for _, child := range p.Children {
		if child.Indefinite {
			childNode, err := child.Parser.Parse(content, i)
			for err == nil {
				n.Add(childNode)
				i += childNode.Len()
				childNode, err = child.Parser.Parse(content, i)
			}
		} else {
			childNode, err := child.Parser.Parse(content, i)
			if err != nil {
				return nil, err
			}
			n.Add(childNode)
			i += childNode.Len()
		}
	}
	return n, nil
}

func (p OrParser) Parse(content string, index int) (node *Node, err error) {
	d("Parse `%s` with %s at %d", content, p, index)
	if index > len(content) {
		return nil, errors.New(fmt.Sprintf("Index exceeds parsed string length %d.", len(content)))
	}
	n, err := p.a.Parse(content, index)
	if err == nil {
		return n, nil
	}
	n, err = p.b.Parse(content, index)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (p *TagParser) Add(child Parser, indefinite bool) error {
	p.Children = append(p.Children, &IndefiniteParser{Parser: child, Indefinite: indefinite})
	return nil
}
