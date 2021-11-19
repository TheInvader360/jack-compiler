package tokenizer

import (
	"testing"

	"github.com/TheInvader360/jack-compiler/common"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	type test struct {
		jackSource     []byte
		expectedTokens []Token
	}
	tests := []test{
		{jackSource: []byte(common.Array_Main_Jack), expectedTokens: Array_Main_Tokens(t)},
		{jackSource: []byte(common.HelloWorld_Main_Jack), expectedTokens: HelloWorld_Main_Tokens(t)},
		{jackSource: []byte(common.Square_Main_Jack), expectedTokens: Square_Main_Tokens(t)},
		{jackSource: []byte(common.Square_Square_Jack), expectedTokens: Square_Square_Tokens(t)},
		{jackSource: []byte(common.Square_SquareGame_Jack), expectedTokens: Square_SquareGame_Tokens(t)},
		{jackSource: []byte(common.SquareExpressionless_Main_Jack), expectedTokens: SquareExpressionless_Main_Tokens(t)},
		{jackSource: []byte(common.SquareExpressionless_Square_Jack), expectedTokens: SquareExpressionless_Square_Tokens(t)},
		{jackSource: []byte(common.SquareExpressionless_SquareGame_Jack), expectedTokens: SquareExpressionless_SquareGame_Tokens(t)},
		{jackSource: []byte(`"semicolon;test"`), expectedTokens: []Token{{TypeOf: TokenTypeStringConstant, Value: "semicolon;test"}}}, // string constants containing semicolons were registering as symbols (logic bug fixed!)
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
		{tokens: Array_Main_Tokens(t), expectedXML: common.Array_MainT_XML},
		{tokens: HelloWorld_Main_Tokens(t), expectedXML: common.HelloWorld_MainT_XML},
		{tokens: Square_Main_Tokens(t), expectedXML: common.Square_MainT_XML},
		{tokens: Square_Square_Tokens(t), expectedXML: common.Square_SquareT_XML},
		{tokens: Square_SquareGame_Tokens(t), expectedXML: common.Square_SquareGameT_XML},
		{tokens: SquareExpressionless_Main_Tokens(t), expectedXML: common.SquareExpressionless_MainT_XML},
		{tokens: SquareExpressionless_Square_Tokens(t), expectedXML: common.SquareExpressionless_SquareT_XML},
		{tokens: SquareExpressionless_SquareGame_Tokens(t), expectedXML: common.SquareExpressionless_SquareGameT_XML},
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
		{jackSource: []byte(common.Array_Main_Jack), expectedXML: common.Array_MainT_XML},
		{jackSource: []byte(common.HelloWorld_Main_Jack), expectedXML: common.HelloWorld_MainT_XML},
		{jackSource: []byte(common.Square_Main_Jack), expectedXML: common.Square_MainT_XML},
		{jackSource: []byte(common.Square_Square_Jack), expectedXML: common.Square_SquareT_XML},
		{jackSource: []byte(common.Square_SquareGame_Jack), expectedXML: common.Square_SquareGameT_XML},
		{jackSource: []byte(common.SquareExpressionless_Main_Jack), expectedXML: common.SquareExpressionless_MainT_XML},
		{jackSource: []byte(common.SquareExpressionless_Square_Jack), expectedXML: common.SquareExpressionless_SquareT_XML},
		{jackSource: []byte(common.SquareExpressionless_SquareGame_Jack), expectedXML: common.SquareExpressionless_SquareGameT_XML},
	}
	for _, tc := range tests {
		tokens := Tokenize(tc.jackSource)
		xml := AsXML(tokens)
		assert.Equal(t, tc.expectedXML, xml)
	}
}
