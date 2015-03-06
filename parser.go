package newtnode

type Parser struct {
	Type    string
	Content string
}

func NewParserChar(content string) *Parser {
	return &Parser{Type: "char", Content: content}
}

func (p *Parser) Parse(content string) (*Node, error) {
	switch p.Type {
	case "char":
		if content == p.Content {
			return NewNodeChar(content, len(content), 0), nil
		}
	}
	return nil, nil
}
