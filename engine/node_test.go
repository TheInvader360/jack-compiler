package engine

import (
	"testing"

	"github.com/TheInvader360/jack-compiler/common"
	"github.com/stretchr/testify/assert"
)

func TestAddChild(t *testing.T) {
	root := Node{Name: "root", Value: "", Children: []Node{}}
	child := Node{Name: "child", Value: "", Children: []Node{}}
	root.AddChild(child)
	expected := Node{Name: "root", Value: "", Children: []Node{{Name: "child", Value: "", Children: []Node{}}}}
	assert.Equal(t, expected, root)
}

func TestAsXML(t *testing.T) {
	type test struct {
		tree        Node
		expectedXML string
	}
	tests := []test{
		{tree: Array_Main_Tree(t), expectedXML: common.Array_Main_XML},
		{tree: HelloWorld_Main_Tree(t), expectedXML: common.HelloWorld_Main_XML},
		{tree: Square_Main_Tree(t), expectedXML: common.Square_Main_XML},
		{tree: Square_Square_Tree(t), expectedXML: common.Square_Square_XML},
		{tree: Square_SquareGame_Tree(t), expectedXML: common.Square_SquareGame_XML},
		{tree: SquareExpressionless_Main_Tree(t), expectedXML: common.SquareExpressionless_Main_XML},
		{tree: SquareExpressionless_Square_Tree(t), expectedXML: common.SquareExpressionless_Square_XML},
		{tree: SquareExpressionless_SquareGame_Tree(t), expectedXML: common.SquareExpressionless_SquareGame_XML},
	}
	for _, tc := range tests {
		xml := tc.tree.AsXML()
		assert.Equal(t, tc.expectedXML, xml)
	}
}

func TestAsExtendedXML(t *testing.T) {
	type test struct {
		tree        Node
		expectedXML string
	}
	tests := []test{
		{tree: HelloWorld_Main_Tree(t), expectedXML: common.HelloWorld_MainE_XML},
		//TODO: more test cases
	}
	for _, tc := range tests {
		xml := tc.tree.AsExtendedXML()
		assert.Equal(t, tc.expectedXML, xml)
	}
}
