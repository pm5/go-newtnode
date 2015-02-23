package gopc

import (
	"fmt"
)

type ASTNode struct {
	Type     string // "regex", "char", "tag"
	Len      int
	Pos      int
	TagName  string
	Content  string
	Children []*ASTNode
}

func (a *ASTNode) String() (out string) {
	switch a.Type {
	case "regex":
		out = fmt.Sprintf("%s", a.Type)
		if len(a.Content) > 0 {
			out += fmt.Sprintf(" '%s'", a.Content)
		}
		break
	case "char":
		out = fmt.Sprintf("%s:%d:%d '%s'", a.Type, a.Len, a.Pos, a.Content)
		break
	case "tag":
		out = fmt.Sprintf("%s", a.TagName)
		for _, c := range a.Children {
			out += "\n  "
			out += fmt.Sprintf("%s", c)
		}
		break
	}
	return
}

func (a *ASTNode) Add(c *ASTNode) {
}
