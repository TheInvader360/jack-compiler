package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	comment         = `(?m)(\/\/.*)|(?s)(\/\*.*?\*\/)`
	keyword         = `(?m)(^class$|^constructor$|^function$|^method$|^field$|^static$|^var$|^int$|^char$|^boolean$|^void$|^true$|^false$|^null$|^this$|^let$|^do$|^if$|^else$|^while$|^return$)`
	symbol          = `(?m)([{}()[\].,;+\-*/&|<>=~])`
	integerConstant = `(?m)(\d+)`
	stringConstant  = `(?m)(\"[^\n]*\")`
	identifier      = `(?m)([A-Za-z_]\w*)`
)

// Tokenize - tokenize the given jack source
func Tokenize(jackSource []byte) []Token {
	commentRegexp := regexp.MustCompile(comment)
	keywordRegexp := regexp.MustCompile(keyword)
	symbolRegexp := regexp.MustCompile(symbol)
	integerConstantRegexp := regexp.MustCompile(integerConstant)
	stringConstantRegexp := regexp.MustCompile(stringConstant)
	identifierRegexp := regexp.MustCompile(identifier)
	tokenValuesRegexp := regexp.MustCompile(keyword + `|` + symbol + `|` + integerConstant + `|` + stringConstant + `|` + identifier)

	jackSource = commentRegexp.ReplaceAll(jackSource, []byte{}) // remove comments

	tokenValues := tokenValuesRegexp.FindAllString(string(jackSource), -1) // find all the individual token values

	tokens := []Token{}
	for _, tokenValue := range tokenValues {
		switch {
		case keywordRegexp.MatchString(tokenValue):
			tokens = append(tokens, Token{TypeOf: TokenTypeKeyword, Value: tokenValue})
		case integerConstantRegexp.MatchString(tokenValue):
			tokens = append(tokens, Token{TypeOf: TokenTypeIntegerConstant, Value: tokenValue})
		case stringConstantRegexp.MatchString(tokenValue):
			tokens = append(tokens, Token{TypeOf: TokenTypeStringConstant, Value: strings.ReplaceAll(tokenValue, "\"", "")})
		case identifierRegexp.MatchString(tokenValue):
			tokens = append(tokens, Token{TypeOf: TokenTypeIdentifier, Value: tokenValue})
		case symbolRegexp.MatchString(tokenValue):
			if tokenValue == "<" {
				tokenValue = "&lt;"
			}
			if tokenValue == ">" {
				tokenValue = "&gt;"
			}
			if tokenValue == "&" {
				tokenValue = "&amp;"
			}
			tokens = append(tokens, Token{TypeOf: TokenTypeSymbol, Value: tokenValue})
		}
	}

	return tokens
}

// AsXML - returns the given list of tokens as xml
func AsXML(tokens []Token) string {
	xml := "<tokens>\n"
	for _, token := range tokens {
		xml += fmt.Sprintf("<%[1]s> %[2]s </%[1]s>\n", token.TypeOf, token.Value)
	}
	xml += "</tokens>\n"

	return xml
}
