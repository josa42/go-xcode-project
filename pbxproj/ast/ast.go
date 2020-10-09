package ast

type NodeType string

const (
	DICT   = NodeType("Dict")
	LIST   = NodeType("LIST")
	STRING = NodeType("String")
)

////////////////////////////////////////////////////////////////////////////////

type Node interface {
	Type() NodeType
}

var _ Node = &DictNode{}

type DictNode struct {
	Values map[string]Node
}

func (d *DictNode) Type() NodeType {
	return NodeType(DICT)
}

func (d *DictNode) Set(key string, value Node) {
	if d.Values == nil {
		d.Values = map[string]Node{}
	}
	d.Values[key] = value
}

////////////////////////////////////////////////////////////////////////////////

var _ Node = &ListNode{}

type ListNode struct {
	Values []Node
}

func (d *ListNode) Type() NodeType {
	return NodeType(LIST)
}

func (d *ListNode) Append(n Node) {
	d.Values = append(d.Values, n)
}

////////////////////////////////////////////////////////////////////////////////

var _ Node = StringNode{}

type StringNode struct {
	Value string
}

func (d StringNode) Type() NodeType {
	return NodeType(STRING)
}
