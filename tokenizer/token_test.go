package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsClassVarDec(t *testing.T) {
	type test struct {
		token          Token
		expectedResult bool
	}
	tests := []test{
		{token: Token{TypeOf: TokenTypeKeyword, Value: "static"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "field"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "other"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIdentifier, Value: "static"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIntegerConstant, Value: "field"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeStringConstant, Value: "static"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "field"}, expectedResult: false},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expectedResult, tc.token.IsClassVarDec())
	}
}

func TestIsKeywordConstant(t *testing.T) {
	type test struct {
		token          Token
		expectedResult bool
	}
	tests := []test{
		{token: Token{TypeOf: TokenTypeKeyword, Value: "true"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "false"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "null"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "this"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "other"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIdentifier, Value: "true"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIntegerConstant, Value: "false"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeStringConstant, Value: "null"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "this"}, expectedResult: false},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expectedResult, tc.token.IsKeywordConstant())
	}
}

func TestIsOp(t *testing.T) {
	type test struct {
		token          Token
		expectedResult bool
	}
	tests := []test{
		{token: Token{TypeOf: TokenTypeSymbol, Value: "+"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "-"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "*"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "/"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "&amp;"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "|"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "&lt;"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "&gt;"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "="}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "other"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIdentifier, Value: "+"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIntegerConstant, Value: "-"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "*"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeStringConstant, Value: "/"}, expectedResult: false},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expectedResult, tc.token.IsOp())
	}
}

func TestIsSubroutineDec(t *testing.T) {
	type test struct {
		token          Token
		expectedResult bool
	}
	tests := []test{
		{token: Token{TypeOf: TokenTypeKeyword, Value: "constructor"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "function"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "method"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "other"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIdentifier, Value: "constructor"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIntegerConstant, Value: "function"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeStringConstant, Value: "method"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "constructor"}, expectedResult: false},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expectedResult, tc.token.IsSubroutineDec())
	}
}

func TestIsType(t *testing.T) {
	type test struct {
		token          Token
		expectedResult bool
	}
	tests := []test{
		{token: Token{TypeOf: TokenTypeKeyword, Value: "int"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "char"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "boolean"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeIdentifier, Value: "any"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "other"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIntegerConstant, Value: "int"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeStringConstant, Value: "char"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "boolean"}, expectedResult: false},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expectedResult, tc.token.IsType())
	}
}

func TestIsTypeOrVoid(t *testing.T) {
	type test struct {
		token          Token
		expectedResult bool
	}
	tests := []test{
		{token: Token{TypeOf: TokenTypeKeyword, Value: "void"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "int"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "char"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "boolean"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeIdentifier, Value: "any"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "other"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIntegerConstant, Value: "void"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeStringConstant, Value: "void"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "void"}, expectedResult: false},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expectedResult, tc.token.IsTypeOrVoid())
	}
}

func TestIsUnaryOp(t *testing.T) {
	type test struct {
		token          Token
		expectedResult bool
	}
	tests := []test{
		{token: Token{TypeOf: TokenTypeSymbol, Value: "-"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "~"}, expectedResult: true},
		{token: Token{TypeOf: TokenTypeSymbol, Value: "other"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIdentifier, Value: "-"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeIntegerConstant, Value: "~"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeKeyword, Value: "-"}, expectedResult: false},
		{token: Token{TypeOf: TokenTypeStringConstant, Value: "~"}, expectedResult: false},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expectedResult, tc.token.IsUnaryOp())
	}
}
