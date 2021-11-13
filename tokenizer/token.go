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
