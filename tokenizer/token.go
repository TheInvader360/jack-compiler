package tokenizer

type TokenType string

const (
	TokenTypeKeyword         TokenType = "keyword"
	TokenTypeSymbol          TokenType = "symbol"
	TokenTypeIntegerConstant TokenType = "integerConstant"
	TokenTypeStringConstant  TokenType = "stringConstant"
	TokenTypeIdentifier      TokenType = "identifier"
)

type Token struct {
	TypeOf TokenType
	Value  string
}

func (t *Token) IsClassVarDec() bool {
	if t.TypeOf != TokenTypeKeyword {
		return false
	}
	if t.Value != "static" && t.Value != "field" {
		return false
	}
	return true
}

func (t *Token) IsOp() bool {
	if t.TypeOf != TokenTypeSymbol {
		return false
	}
	if t.Value != "+" && t.Value != "-" && t.Value != "*" && t.Value != "/" && t.Value != "&amp;" && t.Value != "|" && t.Value != "&lt;" && t.Value != "&gt;" && t.Value != "=" {
		return false
	}
	return true
}

func (t *Token) IsSubroutineDec() bool {
	if t.TypeOf != TokenTypeKeyword {
		return false
	}
	if t.Value != "constructor" && t.Value != "function" && t.Value != "method" {
		return false
	}
	return true
}

func (t *Token) IsType() bool {
	if t.TypeOf == TokenTypeKeyword && (t.Value == "int" || t.Value == "char" || t.Value == "boolean") {
		return true
	}
	if t.TypeOf == TokenTypeIdentifier { // can't know for sure that this is a className 'type' as there is no linking across files...
		return true
	}
	return false
}

func (t *Token) IsTypeOrVoid() bool {
	return t.IsType() || (t.TypeOf == TokenTypeKeyword && t.Value == "void")
}
