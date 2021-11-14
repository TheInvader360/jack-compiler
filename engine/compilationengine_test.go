package engine

import (
	"testing"

	"github.com/TheInvader360/jack-compiler/common"
	"github.com/TheInvader360/jack-compiler/tokenizer"

	"github.com/stretchr/testify/assert"
)

func TestNewCompilationEngine(t *testing.T) {
	tokens := []tokenizer.Token{}
	engine := NewCompilationEngine(tokens)
	assert.Equal(t, tokens, engine.tokens)
	assert.Equal(t, 0, engine.tokenIndex)
}

func TestCompileClass(t *testing.T) {
	type test struct {
		tokens      []tokenizer.Token
		expectedXML string
	}
	tests := []test{
		//TODO: {tokens: tokenizer.Array_Main_Tokens(t), expectedXML: common.Array_Main_XML},
		//TODO: {tokens: tokenizer.HelloWorld_Main_Tokens(t), expectedXML: common.HelloWorld_Main_XML},
		//TODO: {tokens: tokenizer.Square_Main_Tokens(t), expectedXML: common.Square_Main_XML},
		//TODO: {tokens: tokenizer.Square_Square_Tokens(t), expectedXML: common.Square_Square_XML},
		//TODO: {tokens: tokenizer.Square_SquareGame_Tokens(t), expectedXML: common.Square_SquareGame_XML},
		{tokens: tokenizer.SquareExpressionless_Main_Tokens(t), expectedXML: common.SquareExpressionless_Main_XML},
		{tokens: tokenizer.SquareExpressionless_Square_Tokens(t), expectedXML: common.SquareExpressionless_Square_XML},
		{tokens: tokenizer.SquareExpressionless_SquareGame_Tokens(t), expectedXML: common.SquareExpressionless_SquareGame_XML},
	}
	for _, tc := range tests {
		engine := NewCompilationEngine(tc.tokens)
		tree := engine.CompileClass()
		assert.Equal(t, tc.expectedXML, tree)
	}
}

func TestCompileVarDec(t *testing.T) {
	type test struct {
		tokens       []tokenizer.Token
		expectedTree *Node
	}
	tests := []test{
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeKeyword, Value: "var"},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "int"},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "a"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ","},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "b"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ","},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "c"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ";"},
			},
			expectedTree: &Node{Name: "varDec", Value: "", Children: []Node{
				{Name: "keyword", Value: "var", Children: []Node{}},
				{Name: "identifier", Value: "int", Children: []Node{}},
				{Name: "identifier", Value: "a", Children: []Node{}},
				{Name: "symbol", Value: ",", Children: []Node{}},
				{Name: "identifier", Value: "b", Children: []Node{}},
				{Name: "symbol", Value: ",", Children: []Node{}},
				{Name: "identifier", Value: "c", Children: []Node{}},
				{Name: "symbol", Value: ";", Children: []Node{}},
			}},
		},
	}
	for _, tc := range tests {
		engine := NewCompilationEngine(tc.tokens)
		tree := engine.compileVarDec()
		assert.Equal(t, tc.expectedTree, tree)
	}
}

func TestCompileExpression(t *testing.T) {
	type test struct {
		tokens       []tokenizer.Token
		expectedTree *Node
	}
	tests := []test{
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "a"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "+"},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "b"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ";"},
			},
			expectedTree: &Node{Name: "expression", Value: "", Children: []Node{
				{Name: "term", Value: "", Children: []Node{
					{Name: "identifier", Value: "a", Children: []Node{}},
				}},
				{Name: "symbol", Value: "+", Children: []Node{}},
				{Name: "term", Value: "", Children: []Node{
					{Name: "identifier", Value: "b", Children: []Node{}},
				}},
			}},
		},
	}
	for _, tc := range tests {
		engine := NewCompilationEngine(tc.tokens)
		tree := engine.compileExpression()
		assert.Equal(t, tc.expectedTree, tree)
	}
}
