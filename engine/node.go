package engine

import (
	"fmt"
)

type Node struct {
	Name     string
	Value    string
	Children []Node
}

// AddChild - adds a child node to the node
func (n *Node) AddChild(node Node) {
	n.Children = append(n.Children, node)
}

// AsXML - returns the node/tree as xml
func (n *Node) AsXML() string {
	xml := ""

	if n.Value != "" {
		xml += fmt.Sprintf("<%[1]s> %[2]s </%[1]s>\n", n.Name, n.Value)
	} else {
		xml += fmt.Sprintf("<%s>\n", n.Name)
		for _, child := range n.Children {
			xml += child.AsXML()
		}
		xml += fmt.Sprintf("</%s>\n", n.Name)
	}

	return xml
}
