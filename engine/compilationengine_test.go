package engine

import (
	"testing"

	"github.com/TheInvader360/jack-compiler/common"
	"github.com/TheInvader360/jack-compiler/tokenizer"
	"github.com/stretchr/testify/assert"
)

func TestNewCompilationEngine(t *testing.T) {
	tokens := []tokenizer.Token{}
	engine := NewCompilationEngine(tokens, false, false)
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
		engine := NewCompilationEngine(tc.tokens, true, true)
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
		{jack: common.MultiConditionIf_Main_Jack, expectedVM: common.MultiConditionIf_Main_VM},
		{jack: common.Seven_Main_Jack, expectedVM: common.Seven_Main_VM},
		{jack: common.SimpleIf_Main_Jack, expectedVM: common.SimpleIf_Main_VM},
		{jack: common.SimpleWhile_Main_Jack, expectedVM: common.SimpleWhile_Main_VM},
		{jack: common.ConvertToBin_Main_Jack, expectedVM: common.ConvertToBin_Main_VM},
		{jack: common.Square_Main_Jack, expectedVM: common.Square_Main_VM},
		{jack: common.Square_Square_Jack, expectedVM: common.Square_Square_VM},
		{jack: common.Square_SquareGame_Jack, expectedVM: common.Square_SquareGame_VM},
		{jack: common.Average_Main_Jack, expectedVM: common.Average_Main_VM},
		{jack: common.Pong_Ball_Jack, expectedVM: common.Pong_Ball_VM},
		{jack: common.Pong_Bat_Jack, expectedVM: common.Pong_Bat_VM},
		{jack: common.Pong_Main_Jack, expectedVM: common.Pong_Main_VM},
		{jack: common.Pong_PongGame_Jack, expectedVM: common.Pong_PongGame_VM},
		{jack: common.ComplexArrays_Main_Jack, expectedVM: common.ComplexArrays_Main_VM},
		{jack: common.Sokoban_Board_Jack, expectedVM: common.Sokoban_Board_VM},
		{jack: common.Sokoban_Cell_Jack, expectedVM: common.Sokoban_Cell_VM},
		{jack: common.Sokoban_CellType_Jack, expectedVM: common.Sokoban_CellType_VM},
		{jack: common.Sokoban_Controller_Jack, expectedVM: common.Sokoban_Controller_VM},
		{jack: common.Sokoban_Direction_Jack, expectedVM: common.Sokoban_Direction_VM},
		{jack: common.Sokoban_Level_Jack, expectedVM: common.Sokoban_Level_VM},
		{jack: common.Sokoban_LevelManager_Jack, expectedVM: common.Sokoban_LevelManager_VM},
		{jack: common.Sokoban_Main_Jack, expectedVM: common.Sokoban_Main_VM},
		{jack: common.Sokoban_Model_Jack, expectedVM: common.Sokoban_Model_VM},
		{jack: common.Sokoban_Player_Jack, expectedVM: common.Sokoban_Player_VM},
		{jack: common.Sokoban_Sprites_Jack, expectedVM: common.Sokoban_Sprites_VM},
		{jack: common.Sokoban_State_Jack, expectedVM: common.Sokoban_State_VM},
		{jack: common.Sokoban_Utils_Jack, expectedVM: common.Sokoban_Utils_VM},
		{jack: common.Sokoban_View_Jack, expectedVM: common.Sokoban_View_VM},
	}
	for _, tc := range tests {
		tokens := tokenizer.Tokenize([]byte(tc.jack))
		engine := NewCompilationEngine(tokens, false, false)
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
				{TypeOf: tokenizer.TokenTypeKeyword, Value: "int"},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "a"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ","},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "b"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ","},
				{TypeOf: tokenizer.TokenTypeIdentifier, Value: "c"},
				{TypeOf: tokenizer.TokenTypeSymbol, Value: ";"},
			},
			expectedTree: &Node{Name: "varDec", Value: "", Children: []Node{
				{Name: "keyword", Value: "var", Children: []Node{}},
				{Name: "keyword", Value: "int", Children: []Node{}},
				{Name: "identifier", Value: "a", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
				{Name: "symbol", Value: ",", Children: []Node{}},
				{Name: "identifier", Value: "b", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "local", IdentifierIndex: "TODO"},
				{Name: "symbol", Value: ",", Children: []Node{}},
				{Name: "identifier", Value: "c", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "local", IdentifierIndex: "TODO"},
				{Name: "symbol", Value: ";", Children: []Node{}},
			}},
		},
	}
	for _, tc := range tests {
		engine := NewCompilationEngine(tc.tokens, false, false)
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
					{Name: "identifier", Value: "a", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
				}},
				{Name: "symbol", Value: "+", Children: []Node{}},
				{Name: "term", Value: "", Children: []Node{
					{Name: "identifier", Value: "b", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
				}},
			}},
		},
	}
	for _, tc := range tests {
		engine := NewCompilationEngine(tc.tokens, false, false)
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
				{Name: "identifier", Value: "varName", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
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
				{Name: "identifier", Value: "arr", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
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
				{Name: "identifier", Value: "subroutineName", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
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
				{Name: "identifier", Value: "ClassName", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
				{Name: "symbol", Value: ".", Children: []Node{}},
				{Name: "identifier", Value: "subroutineName", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
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
							{Name: "identifier", Value: "x", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
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
					{Name: "identifier", Value: "x", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
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
					{Name: "identifier", Value: "x", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
				}},
			}},
		},
	}
	for _, tc := range tests {
		engine := NewCompilationEngine(tc.tokens, false, false)
		tree := engine.compileTerm()
		assert.Equal(t, tc.expectedTree, tree)
	}
}
