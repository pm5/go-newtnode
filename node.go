package spc

import (
	"fmt"
)

type Node struct {
	Type     string // "regexp", "char", "tag"
	Pos      int
	TagName  string
	Content  string
	Children []*Node
}

func NewNodeTag(name string) *Node {
	n := &Node{Type: "tag", TagName: name}
	n.Children = make([]*Node, 0)
	return n
}

func NewNodeRegexp(content string) *Node {
	return &Node{Type: "regexp", Content: content}
}

func NewNodeChar(content string, position int) *Node {
	return &Node{Type: "char", Pos: position, Content: content}
}

func (a *Node) String() (out string) {
	switch a.Type {
	case "regexp":
		out = fmt.Sprintf("%s", a.Type)
		if len(a.Content) > 0 {
			out += fmt.Sprintf(" '%s'", a.Content)
		}
		break
	case "char":
		out = fmt.Sprintf("%s:%d:%d '%s'", a.Type, a.Len(), a.Pos, a.Content)
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

func (a *Node) Add(c *Node) {
	a.Children = append(a.Children, c)
}

func (a *Node) Delete(i int) {
	a.Children = append(a.Children[:i], a.Children[i+1:len(a.Children)]...)
}

func (a *Node) Length() int {
	return len(a.Children)
}

func (n *Node) Len() int {
	if n.Type == "char" || n.Type == "regexp" {
		return len(n.Content)
	}
	var sum int
	for _, childNode := range n.Children {
		sum += childNode.Len()
	}
	return sum
}
