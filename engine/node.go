package engine

import (
	"fmt"
)

type Node struct {
	Name               string
	Value              string
	Children           []Node
	IdentifierCategory string
	IdentifierAction   string
	IdentifierKind     string
	IdentifierIndex    string
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

// AsExtendedXML - returns the node/tree as xml (with the extra symbol table info added)
func (n *Node) AsExtendedXML() string {
	xml := ""

	if n.Value != "" {
		if n.Name == "identifier" {
			xml += fmt.Sprintf("<%[1]s category=\"%[3]s\" action=\"%[4]s\" kind=\"%[5]s\" index=\"%[6]s\"> %[2]s </%[1]s>\n", n.Name, n.Value, n.IdentifierCategory, n.IdentifierAction, n.IdentifierKind, n.IdentifierIndex)
		} else {
			xml += fmt.Sprintf("<%[1]s> %[2]s </%[1]s>\n", n.Name, n.Value)
		}
	} else {
		xml += fmt.Sprintf("<%s>\n", n.Name)
		for _, child := range n.Children {
			xml += child.AsExtendedXML()
		}
		xml += fmt.Sprintf("</%s>\n", n.Name)
	}

	return xml
}
