package engine

import (
	"testing"

	"github.com/TheInvader360/jack-compiler/common"
	"github.com/TheInvader360/jack-compiler/tokenizer"

	"github.com/stretchr/testify/assert"
)

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
		xml := CompileClass(tc.tokens)
		assert.Equal(t, tc.expectedXML, xml)
	}
}
