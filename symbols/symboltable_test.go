package symbols

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEverything(t *testing.T) {
	cst := NewCombinedSymbolTable()

	cst.Define("s1", "boolean", SymbolKindStatic)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindStatic))
	assert.Equal(t, "boolean", cst.getTypeOf("s1"))
	assert.Equal(t, SymbolKindStatic, cst.getKindOf("s1"))
	assert.Equal(t, 0, cst.getIndexOf("s1"))

	cst.Define("s2", "int", SymbolKindStatic)
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindStatic))
	assert.Equal(t, "int", cst.getTypeOf("s2"))
	assert.Equal(t, SymbolKindStatic, cst.getKindOf("s2"))
	assert.Equal(t, 1, cst.getIndexOf("s2"))

	cst.Define("f1", "char", SymbolKindField)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindField))
	assert.Equal(t, "char", cst.getTypeOf("f1"))
	assert.Equal(t, SymbolKindField, cst.getKindOf("f1"))
	assert.Equal(t, 0, cst.getIndexOf("f1"))

	cst.Define("f2", "SomeClassName", SymbolKindField)
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindField))
	assert.Equal(t, "SomeClassName", cst.getTypeOf("f2"))
	assert.Equal(t, SymbolKindField, cst.getKindOf("f2"))
	assert.Equal(t, 1, cst.getIndexOf("f2"))

	cst.Define("a1", "boolean", SymbolKindArgument)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindArgument))
	assert.Equal(t, "boolean", cst.getTypeOf("a1"))
	assert.Equal(t, SymbolKindArgument, cst.getKindOf("a1"))
	assert.Equal(t, 0, cst.getIndexOf("a1"))

	cst.Define("a2", "int", SymbolKindArgument)
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindArgument))
	assert.Equal(t, "int", cst.getTypeOf("a2"))
	assert.Equal(t, SymbolKindArgument, cst.getKindOf("a2"))
	assert.Equal(t, 1, cst.getIndexOf("a2"))

	cst.Define("v1", "char", SymbolKindVar)
	assert.Equal(t, 1, cst.GetVarCount(SymbolKindVar))
	assert.Equal(t, "char", cst.getTypeOf("v1"))
	assert.Equal(t, SymbolKindVar, cst.getKindOf("v1"))
	assert.Equal(t, 0, cst.getIndexOf("v1"))

	cst.Define("v2", "SomeClassName", SymbolKindVar)
	assert.Equal(t, 2, cst.GetVarCount(SymbolKindVar))
	assert.Equal(t, "SomeClassName", cst.getTypeOf("v2"))
	assert.Equal(t, SymbolKindVar, cst.getKindOf("v2"))
	assert.Equal(t, 1, cst.getIndexOf("v2"))

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

	assert.Equal(t, "", cst.getTypeOf("Unknown"))
	assert.Equal(t, SymbolKindUnknown, cst.getKindOf("Unknown"))
	assert.Equal(t, -1, cst.getIndexOf("Unknown"))
}
