package engine

import (
	"fmt"
)

type Node struct {
	Name     string
	Value    string
	Children []Node
}

func (n *Node) AddChild(node Node) {
	//fmt.Println("AddChild", node.Name, node.Value)
	n.Children = append(n.Children, node)
}

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
