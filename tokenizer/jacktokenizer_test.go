package tokenizer

import (
	"testing"

	"github.com/TheInvader360/jack-compiler/fixtures"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	type test struct {
		jackSource     []byte
		expectedTokens []Token
	}
	tests := []test{
		{
			jackSource: []byte(fixtures.HelloWorld_Main_Jack),
			expectedTokens: []Token{
				{TypeOf: TokenTypeKeyword, Value: "class"},
				{TypeOf: TokenTypeIdentifier, Value: "Main"},
				{TypeOf: TokenTypeSymbol, Value: "{"},
				{TypeOf: TokenTypeKeyword, Value: "function"},
				{TypeOf: TokenTypeKeyword, Value: "void"},
				{TypeOf: TokenTypeIdentifier, Value: "main"},
				{TypeOf: TokenTypeSymbol, Value: "("},
				{TypeOf: TokenTypeSymbol, Value: ")"},
				{TypeOf: TokenTypeSymbol, Value: "{"},
				{TypeOf: TokenTypeKeyword, Value: "do"},
				{TypeOf: TokenTypeIdentifier, Value: "Output"},
				{TypeOf: TokenTypeSymbol, Value: "."},
				{TypeOf: TokenTypeIdentifier, Value: "printString"},
				{TypeOf: TokenTypeSymbol, Value: "("},
				{TypeOf: TokenTypeStringConstant, Value: "Hello World"},
				{TypeOf: TokenTypeSymbol, Value: ")"},
				{TypeOf: TokenTypeSymbol, Value: ";"},
				{TypeOf: TokenTypeKeyword, Value: "return"},
				{TypeOf: TokenTypeSymbol, Value: ";"},
				{TypeOf: TokenTypeSymbol, Value: "}"},
				{TypeOf: TokenTypeSymbol, Value: "}"},
			},
		},
	}
	for _, tc := range tests {
		tokens := Tokenize(tc.jackSource)
		assert.Equal(t, tc.expectedTokens, tokens)
	}
}

func TestAsXML(t *testing.T) {
	type test struct {
		tokens      []Token
		expectedXML string
	}
	tests := []test{
		{
			tokens: []Token{
				{TypeOf: TokenTypeKeyword, Value: "class"},
				{TypeOf: TokenTypeIdentifier, Value: "Main"},
				{TypeOf: TokenTypeSymbol, Value: "{"},
				{TypeOf: TokenTypeKeyword, Value: "function"},
				{TypeOf: TokenTypeKeyword, Value: "void"},
				{TypeOf: TokenTypeIdentifier, Value: "main"},
				{TypeOf: TokenTypeSymbol, Value: "("},
				{TypeOf: TokenTypeSymbol, Value: ")"},
				{TypeOf: TokenTypeSymbol, Value: "{"},
				{TypeOf: TokenTypeKeyword, Value: "do"},
				{TypeOf: TokenTypeIdentifier, Value: "Output"},
				{TypeOf: TokenTypeSymbol, Value: "."},
				{TypeOf: TokenTypeIdentifier, Value: "printString"},
				{TypeOf: TokenTypeSymbol, Value: "("},
				{TypeOf: TokenTypeStringConstant, Value: "Hello World"},
				{TypeOf: TokenTypeSymbol, Value: ")"},
				{TypeOf: TokenTypeSymbol, Value: ";"},
				{TypeOf: TokenTypeKeyword, Value: "return"},
				{TypeOf: TokenTypeSymbol, Value: ";"},
				{TypeOf: TokenTypeSymbol, Value: "}"},
				{TypeOf: TokenTypeSymbol, Value: "}"},
			},
			expectedXML: fixtures.HelloWorld_MainT_XML,
		},
	}
	for _, tc := range tests {
		xml := AsXML(tc.tokens)
		assert.Equal(t, tc.expectedXML, xml)
	}
}

func TestTokenizeAsXML(t *testing.T) {
	type test struct {
		jackSource  []byte
		expectedXML string
	}
	tests := []test{
		{
			jackSource:  []byte(fixtures.HelloWorld_Main_Jack),
			expectedXML: fixtures.HelloWorld_MainT_XML,
		},
		{
			jackSource:  []byte(fixtures.Array_Main_Jack),
			expectedXML: fixtures.Array_MainT_XML,
		},
		{
			jackSource:  []byte(fixtures.Square_Main_Jack),
			expectedXML: fixtures.Square_MainT_XML,
		},
		{
			jackSource:  []byte(fixtures.Square_Square_Jack),
			expectedXML: fixtures.Square_SquareT_XML,
		},
		{
			jackSource:  []byte(fixtures.Square_SquareGame_Jack),
			expectedXML: fixtures.Square_SquareGameT_XML,
		},
	}
	for _, tc := range tests {
		tokens := Tokenize(tc.jackSource)
		xml := AsXML(tokens)
		assert.Equal(t, tc.expectedXML, xml)
	}
}
