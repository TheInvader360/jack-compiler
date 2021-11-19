package symbols

type SymbolKind string

const (
	SymbolKindVar      SymbolKind = "local"
	SymbolKindArgument SymbolKind = "argument"
	SymbolKindStatic   SymbolKind = "static"
	SymbolKindField    SymbolKind = "field"
	SymbolKindUnknown  SymbolKind = ""
)

type Symbol struct {
	identifierName string
	typeOf         string
	kindOf         SymbolKind
	indexOf        int
}

type SymbolTable map[string]Symbol

type CombinedSymbolTable struct {
	classScopeSymbolTable      SymbolTable
	subroutineScopeSymbolTable SymbolTable
}

// NewCombinedSymbolTable - creates a combined symbol table (contains both an empty class scope symbol table and an empty subroutine scope symbol table)
func NewCombinedSymbolTable() *CombinedSymbolTable {
	st := CombinedSymbolTable{
		classScopeSymbolTable:      SymbolTable{},
		subroutineScopeSymbolTable: SymbolTable{},
	}

	return &st
}

// StartSubroutine - start a new subroutine (resets the subroutine symbol table)
func (cst *CombinedSymbolTable) StartSubroutine() {
	cst.subroutineScopeSymbolTable = SymbolTable{}
}

// Define - defines a new identifier of a given name, type, and kind, and assigns it a running index
func (cst *CombinedSymbolTable) Define(identifierName, typeOf string, kind SymbolKind) {
	symbol := Symbol{identifierName: identifierName, typeOf: typeOf, kindOf: SymbolKind(kind)}
	if symbol.kindOf == SymbolKindStatic || symbol.kindOf == SymbolKindField {
		// static and field identifiers have a class scope
		symbol.indexOf = cst.classScopeSymbolTable.getVarCount(symbol.kindOf)
		cst.classScopeSymbolTable[identifierName] = symbol
	}
	if symbol.kindOf == SymbolKindArgument || symbol.kindOf == SymbolKindVar {
		// argument and var identifiers have a subroutine scope
		symbol.indexOf = cst.subroutineScopeSymbolTable.getVarCount(symbol.kindOf)
		cst.subroutineScopeSymbolTable[identifierName] = symbol
	}
}

// GetVarCount - returns the number of variables of the given kind already defined in the relevant symbol table
func (cst *CombinedSymbolTable) GetVarCount(kind SymbolKind) int {
	if kind == SymbolKindArgument || kind == SymbolKindVar {
		return cst.subroutineScopeSymbolTable.getVarCount(kind)
	}
	return cst.classScopeSymbolTable.getVarCount(kind)
}

// getVarCount - returns the number of variables of the given kind already defined in the given symbol table
func (st *SymbolTable) getVarCount(kind SymbolKind) int {
	varCount := 0
	for _, symbol := range *st {
		if symbol.kindOf == kind {
			varCount++
		}
	}

	return varCount
}

// getTypeOf - Returns the type of the given identifier
func (cst *CombinedSymbolTable) getTypeOf(identifierName string) string {
	if symbol, ok := cst.classScopeSymbolTable[identifierName]; ok {
		return symbol.typeOf
	}
	if symbol, ok := cst.subroutineScopeSymbolTable[identifierName]; ok {
		return symbol.typeOf
	}

	return ""
}

// getKindOf - Returns the kind of the given identifier
func (cst *CombinedSymbolTable) getKindOf(identifierName string) SymbolKind {
	if symbol, ok := cst.classScopeSymbolTable[identifierName]; ok {
		return symbol.kindOf
	}
	if symbol, ok := cst.subroutineScopeSymbolTable[identifierName]; ok {
		return symbol.kindOf
	}

	return SymbolKindUnknown
}

// getIndexOf - Returns the index of the given identifier
func (cst *CombinedSymbolTable) getIndexOf(identifierName string) int {
	if symbol, ok := cst.classScopeSymbolTable[identifierName]; ok {
		return symbol.indexOf
	}
	if symbol, ok := cst.subroutineScopeSymbolTable[identifierName]; ok {
		return symbol.indexOf
	}

	return -1
}
