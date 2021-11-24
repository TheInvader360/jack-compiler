package symbols

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEverything(t *testing.T) {
	cst := NewCombinedSymbolTable()

	cst.Define("s1", "boolean", SymbolKindStatic)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindStatic))
	assert.Equal(t, "boolean", cst.GetTypeOf("s1"))
	assert.Equal(t, SymbolKindStatic, cst.GetKindOf("s1"))
	assert.Equal(t, 0, cst.GetIndexOf("s1"))

	cst.Define("s2", "int", SymbolKindStatic)
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindStatic))
	assert.Equal(t, "int", cst.GetTypeOf("s2"))
	assert.Equal(t, SymbolKindStatic, cst.GetKindOf("s2"))
	assert.Equal(t, 1, cst.GetIndexOf("s2"))

	cst.Define("f1", "char", SymbolKindField)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindField))
	assert.Equal(t, "char", cst.GetTypeOf("f1"))
	assert.Equal(t, SymbolKindField, cst.GetKindOf("f1"))
	assert.Equal(t, 0, cst.GetIndexOf("f1"))

	cst.Define("f2", "SomeClassName", SymbolKindField)
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindField))
	assert.Equal(t, "SomeClassName", cst.GetTypeOf("f2"))
	assert.Equal(t, SymbolKindField, cst.GetKindOf("f2"))
	assert.Equal(t, 1, cst.GetIndexOf("f2"))

	cst.Define("a1", "boolean", SymbolKindArgument)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindArgument))
	assert.Equal(t, "boolean", cst.GetTypeOf("a1"))
	assert.Equal(t, SymbolKindArgument, cst.GetKindOf("a1"))
	assert.Equal(t, 0, cst.GetIndexOf("a1"))

	cst.Define("a2", "int", SymbolKindArgument)
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindArgument))
	assert.Equal(t, "int", cst.GetTypeOf("a2"))
	assert.Equal(t, SymbolKindArgument, cst.GetKindOf("a2"))
	assert.Equal(t, 1, cst.GetIndexOf("a2"))

	cst.Define("v1", "char", SymbolKindVar)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindVar))
	assert.Equal(t, "char", cst.GetTypeOf("v1"))
	assert.Equal(t, SymbolKindVar, cst.GetKindOf("v1"))
	assert.Equal(t, 0, cst.GetIndexOf("v1"))

	cst.Define("v2", "SomeClassName", SymbolKindVar)
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindVar))
	assert.Equal(t, "SomeClassName", cst.GetTypeOf("v2"))
	assert.Equal(t, SymbolKindVar, cst.GetKindOf("v2"))
	assert.Equal(t, 1, cst.GetIndexOf("v2"))

	cst.StartSubroutine()
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindStatic))
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindField))
	assert.Equal(t, 0, cst.GetVarCount(SymbolKindArgument))
	assert.Equal(t, 0, cst.GetVarCount(SymbolKindVar))

	cst.Define("s3", "Anything", SymbolKindStatic)
	assert.Equal(t, 3, cst.GetVarCount(SymbolKindStatic))

	cst.Define("f3", "SomeClassName", SymbolKindField)
	assert.Equal(t, 3, cst.GetVarCount(SymbolKindField))

	cst.Define("a3", "SomeClassName", SymbolKindArgument)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindArgument))

	cst.Define("v3", "SomeClassName", SymbolKindVar)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindVar))

	assert.Equal(t, "", cst.GetTypeOf("Unknown"))
	assert.Equal(t, SymbolKindUnknown, cst.GetKindOf("Unknown"))
	assert.Equal(t, -1, cst.GetIndexOf("Unknown"))

	assert.Equal(t, "\nC: map[f1:{identifierName:f1 typeOf:char kindOf:field indexOf:0} f2:{identifierName:f2 typeOf:SomeClassName kindOf:field indexOf:1} f3:{identifierName:f3 typeOf:SomeClassName kindOf:field indexOf:2} s1:{identifierName:s1 typeOf:boolean kindOf:static indexOf:0} s2:{identifierName:s2 typeOf:int kindOf:static indexOf:1} s3:{identifierName:s3 typeOf:Anything kindOf:static indexOf:2}]\nS: map[a3:{identifierName:a3 typeOf:SomeClassName kindOf:argument indexOf:0} v3:{identifierName:v3 typeOf:SomeClassName kindOf:local indexOf:0}]\n", cst.String())
}
