package spc

import (
	"errors"
	"fmt"
	"github.com/tj/go-debug"
	"regexp"
)

var d = debug.Debug("spc:parser")

type IndefiniteParser struct {
	Parser     *Parser
	Indefinite bool
}

type Parser struct {
	Type     string
	Content  string
	Pattern  *regexp.Regexp
	Name     string
	Children []*IndefiniteParser
}

func NewParserChar(content string) *Parser {
	return &Parser{Type: "char", Content: content}
}

func NewParserRegexp(pattern string) *Parser {
	return &Parser{Type: "regexp", Pattern: regexp.MustCompile(pattern)}
}

func NewParserTag(name string, children ...*Parser) *Parser {
	p := Parser{Type: "tag", Name: name}
	for _, child := range children {
		p.Add(child, false)
	}
	return &p
}

func (p *Parser) String() string {
	switch p.Type {
	case "char":
		return "char `" + p.Content + "`"
	case "regexp":
		return "regexp `" + p.Pattern.String() + "`"
	case "tag":
		return "tag `" + p.Name + "`"
	case "or":
		return "or"
	}
	return ""
}

func NewParserOr(a, b *Parser) *Parser {
	p := Parser{Type: "or"}
	p.Add(a, false)
	p.Add(b, false)
	return &p
}

func (p *Parser) Parse(content string, index int) (node *Node, err error) {
	d("Parse `%s` with %s at %d", content, p, index)
	if (p.Type == "char" && index >= len(content)) || index > len(content) {
		return nil, errors.New(fmt.Sprintf("Index exceeds parsed string length %d.", len(content)))
	}
	switch p.Type {
	case "char":
		if content[index] == p.Content[0] {
			return NewNodeChar(p.Content, index), nil
		}
		return nil, errors.New(fmt.Sprintf("Parser `%c` at index %d got `%c`.", p.Content[0], index, content[index]))
	case "regexp":
		r := p.Pattern.FindStringIndex(content[index:])
		if r == nil {
			return nil, errors.New(fmt.Sprintf("Parser `%s` at index %d does not match.", p.Pattern, index))
		}
		return NewNodeRegexp(content[index+r[0] : index+r[1]]), nil
	case "tag":
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
	case "or":
		n, err := p.Children[0].Parser.Parse(content, index)
		if err == nil {
			return n, nil
		}
		n, err = p.Children[1].Parser.Parse(content, index)
		if err != nil {
			return nil, err
		}
		return n, nil
	}
	return nil, errors.New(fmt.Sprintf("Unknown parser type `%s`.", p.Type))
}

func (p *Parser) Add(child *Parser, indefinite bool) error {
	if p.Type != "tag" && p.Type != "or" {
		return errors.New(fmt.Sprintf("Wrong type for Add(). Expected `tag`, got `%s`", p.Type))
	}
	p.Children = append(p.Children, &IndefiniteParser{Parser: child, Indefinite: indefinite})
	return nil
}
