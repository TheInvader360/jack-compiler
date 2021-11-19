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

func TestCompileClassXML(t *testing.T) {
	type test struct {
		tokens      []tokenizer.Token
		expectedXML string
	}
	tests := []test{
		{tokens: tokenizer.Array_Main_Tokens(t), expectedXML: common.Array_Main_XML},
		{tokens: tokenizer.HelloWorld_Main_Tokens(t), expectedXML: common.HelloWorld_Main_XML},
		{tokens: tokenizer.Square_Main_Tokens(t), expectedXML: common.Square_Main_XML},
		{tokens: tokenizer.Square_Square_Tokens(t), expectedXML: common.Square_Square_XML},
		{tokens: tokenizer.Square_SquareGame_Tokens(t), expectedXML: common.Square_SquareGame_XML},
		{tokens: tokenizer.SquareExpressionless_Main_Tokens(t), expectedXML: common.SquareExpressionless_Main_XML},
		{tokens: tokenizer.SquareExpressionless_Square_Tokens(t), expectedXML: common.SquareExpressionless_Square_XML},
		{tokens: tokenizer.SquareExpressionless_SquareGame_Tokens(t), expectedXML: common.SquareExpressionless_SquareGame_XML},
	}
	for _, tc := range tests {
		engine := NewCompilationEngine(tc.tokens)
		tree, _ := engine.CompileClass()
		assert.Equal(t, tc.expectedXML, tree)
	}
}

func TestCompileClassVM(t *testing.T) {
	type test struct {
		jack       string
		expectedVM string
	}
	tests := []test{
		{jack: common.Seven_Main_Jack, expectedVM: common.Seven_Main_VM},
		//TODO: {jack: common.ConvertToBin_Main_Jack, expectedVM: common.ConvertToBin_Main_VM},
		//TODO: {jack: common.Square_Main_Jack, expectedVM: common.Square_Main_VM},
		//TODO: {jack: common.Square_Square_Jack, expectedVM: common.Square_Square_VM},
		//TODO: {jack: common.Square_SquareGame_Jack, expectedVM: common.Square_SquareGame_VM},
		//TODO: {jack: common.Average_Main_Jack, expectedVM: common.Average_Main_VM},
		//TODO: {jack: common.Pong_Ball_Jack, expectedVM: common.Pong_Ball_VM},
		//TODO: {jack: common.Pong_Bat_Jack, expectedVM: common.Pong_Bat_VM},
		//TODO: {jack: common.Pong_Main_Jack, expectedVM: common.Pong_Main_VM},
		//TODO: {jack: common.Pong_PongGame_Jack, expectedVM: common.Pong_PongGame_VM},
		//TODO: {jack: common.ComplexArrays_Main_Jack, expectedVM: common.ComplexArrays_Main_VM},
	}
	for _, tc := range tests {
		tokens := tokenizer.Tokenize([]byte(tc.jack))
		engine := NewCompilationEngine(tokens)
		_, vmCode := engine.CompileClass()
		assert.Equal(t, tc.expectedVM, vmCode)
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

func TestCompileTerm(t *testing.T) {
	type test struct {
		tokens       []tokenizer.Token
		expectedTree *Node
	}
	tests := []test{
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeIntegerConstant, Value: "0"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "integerConstant", Value: "0", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeStringConstant, Value: "abc"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "stringConstant", Value: "abc", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeKeyword, Value: "true"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "keyword", Value: "true", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeKeyword, Value: "false"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "keyword", Value: "false", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeKeyword, Value: "null"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "keyword", Value: "null", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeKeyword, Value: "this"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "keyword", Value: "this", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "varName"},
				{TypeOf: tokenizer.TokenTypeStringConstant, Value: "anything..."}, // there should always be more tokens in a valid jack class (no need guard against out of range in code)
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "identifier", Value: "varName", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "arr"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "["},
				{TypeOf: tokenizer.TokenTypeIntegerConstant, Value: "0"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "]"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "identifier", Value: "arr", Children: []Node{}},
				{Name: "symbol", Value: "[", Children: []Node{}},
				{Name: "expression", Value: "", Children: []Node{
					{Name: "term", Value: "", Children: []Node{
						{Name: "integerConstant", Value: "0", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "]", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "subroutineName"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "("},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ")"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "identifier", Value: "subroutineName", Children: []Node{}},
				{Name: "symbol", Value: "(", Children: []Node{}},
				{Name: "expressionList", Value: "", Children: []Node{}},
				{Name: "symbol", Value: ")", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "ClassName"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "."},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "subroutineName"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "("},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ")"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "identifier", Value: "ClassName", Children: []Node{}},
				{Name: "symbol", Value: ".", Children: []Node{}},
				{Name: "identifier", Value: "subroutineName", Children: []Node{}},
				{Name: "symbol", Value: "(", Children: []Node{}},
				{Name: "expressionList", Value: "", Children: []Node{}},
				{Name: "symbol", Value: ")", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "("},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "-"},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "x"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ")"},
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "symbol", Value: "(", Children: []Node{}},
				{Name: "expression", Value: "", Children: []Node{
					{Name: "term", Value: "", Children: []Node{
						{Name: "symbol", Value: "-", Children: []Node{}},
						{Name: "term", Value: "", Children: []Node{
							{Name: "identifier", Value: "x", Children: []Node{}},
						}},
					}},
				}},
				{Name: "symbol", Value: ")", Children: []Node{}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "-"},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "x"},
				{TypeOf: tokenizer.TokenTypeStringConstant, Value: "anything..."}, // there should always be more tokens in a valid jack class (no need guard against out of range in code)
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "symbol", Value: "-", Children: []Node{}},
				{Name: "term", Value: "", Children: []Node{
					{Name: "identifier", Value: "x", Children: []Node{}},
				}},
			}},
		},
		{
			tokens: []tokenizer.Token{
				{TypeOf: tokenizer.TokenTypeSymbol, Value: "~"},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "x"},
				{TypeOf: tokenizer.TokenTypeStringConstant, Value: "anything..."}, // there should always be more tokens in a valid jack class (no need guard against out of range in code)
			},
			expectedTree: &Node{Name: "term", Value: "", Children: []Node{
				{Name: "symbol", Value: "~", Children: []Node{}},
				{Name: "term", Value: "", Children: []Node{
					{Name: "identifier", Value: "x", Children: []Node{}},
				}},
			}},
		},
	}
	for _, tc := range tests {
		engine := NewCompilationEngine(tc.tokens)
		tree := engine.compileTerm()
		assert.Equal(t, tc.expectedTree, tree)
	}
}
