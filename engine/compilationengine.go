package engine

import (
	"fmt"

	"github.com/TheInvader360/jack-compiler/handler"
	"github.com/TheInvader360/jack-compiler/tokenizer"

	"github.com/pkg/errors"
)

var tokens = []tokenizer.Token{}
var tokenIndex = 0

func CompileClass(t []tokenizer.Token) string {
	// Grammar: 'class' className '{' classVarDec* subroutineDec* '}'

	tokens = t
	tokenIndex = 0

	root := Node{Name: "class", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeKeyword, []string{"class"}) // 'class'
	root.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // className
	root.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
	root.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	for { // classVarDec*
		classVarDec := compileClassVarDec()
		if classVarDec != nil {
			root.AddChild(*classVarDec)
		} else {
			break
		}
	}

	for { // subroutineDec*
		subroutineDec := compileSubroutineDec()
		if subroutineDec != nil {
			root.AddChild(*subroutineDec)
		} else {
			break
		}
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
	root.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	return root.AsXML()
}

func compileClassVarDec() *Node {
	// Grammar: ('static' | 'field') type varName (',' varName)* ';'
	// Where type: 'int' | 'char' | 'boolean' | className

	token := tokens[tokenIndex]

	if token.TypeOf == tokenizer.TokenTypeKeyword && (token.Value == "static" || token.Value == "field") {
		node := Node{Name: "classVarDec", Value: "", Children: []Node{}}

		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}}) // ('static' | 'field')
		tokenIndex++

		token = tokens[tokenIndex]
		if token.TypeOf == tokenizer.TokenTypeKeyword {
			requireToken(token, tokenizer.TokenTypeKeyword, []string{"int", "char", "boolean"}) // type - 'int' | 'char' | 'boolean'
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		} else if token.TypeOf == tokenizer.TokenTypeIdentifier {
			requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // type - className
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		}

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // varName
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		token = tokens[tokenIndex]
		for token.TypeOf == tokenizer.TokenTypeSymbol && token.Value == "," { // (',' varName)*
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}}) // ','
			tokenIndex++

			token = tokens[tokenIndex]
			requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // varName
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		}

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeSymbol, []string{";"}) // ';'
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		return &node
	} else {
		return nil
	}
}

func compileSubroutineDec() *Node {
	// Grammar: ('constructor' | 'function' | 'method') ('void' | type) subroutineName '(' parameterList ')' subroutineBody
	// Where type: 'int' | 'char' | 'boolean' | className

	token := tokens[tokenIndex]

	if token.TypeOf == tokenizer.TokenTypeKeyword && (token.Value == "constructor" || token.Value == "function" || token.Value == "method") {
		node := Node{Name: "subroutineDec", Value: "", Children: []Node{}}

		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}}) // ('constructor' | 'function' | 'method')
		tokenIndex++

		token = tokens[tokenIndex]
		if token.TypeOf == tokenizer.TokenTypeKeyword {
			requireToken(token, tokenizer.TokenTypeKeyword, []string{"void", "int", "char", "boolean"}) // ('void' | 'type') - 'int' | 'char' | 'boolean'
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		}
		if token.TypeOf == tokenizer.TokenTypeIdentifier {
			requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // type - className
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		}
		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // subroutineName
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeSymbol, []string{"("}) // '('
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		parameterList := compileParameterList() // parameterList
		if parameterList != nil {
			node.AddChild(*parameterList)
		}

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeSymbol, []string{")"}) // ')'
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		subroutineBody := compileSubroutineBody() // subroutineBody
		if subroutineBody != nil {
			node.AddChild(*subroutineBody)
		}

		return &node
	} else {
		return nil
	}
}

func compileParameterList() *Node {
	// Grammar: ((type varName) (',' type varName)*)?

	node := Node{Name: "parameterList", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	if token.TypeOf == tokenizer.TokenTypeSymbol && token.Value == ")" { // no parameters...
		return &node
	} else { // first parameter
		token = tokens[tokenIndex]
		if token.TypeOf == tokenizer.TokenTypeKeyword {
			requireToken(token, tokenizer.TokenTypeKeyword, []string{"int", "char", "boolean"}) // type - 'int' | 'char' | 'boolean'
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		} else if token.TypeOf == tokenizer.TokenTypeIdentifier {
			requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // type - className
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		}

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // varName
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++
	}

	token = tokens[tokenIndex]
	for token.TypeOf == tokenizer.TokenTypeSymbol && token.Value == "," { // (',' type varName)*
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}}) // ','
		tokenIndex++

		token = tokens[tokenIndex]
		if token.TypeOf == tokenizer.TokenTypeKeyword {
			requireToken(token, tokenizer.TokenTypeKeyword, []string{"int", "char", "boolean"}) // type - 'int' | 'char' | 'boolean'
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		} else if token.TypeOf == tokenizer.TokenTypeIdentifier {
			requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // type - className
			node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
			tokenIndex++
		}

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // varName
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		token = tokens[tokenIndex]
	}

	return &node
}

func compileSubroutineBody() *Node {
	// Grammar: '{' varDec* statements '}'

	node := Node{Name: "subroutineBody", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	for { // varDec*
		token = tokens[tokenIndex]
		if token.Value == "var" {
			varDec := compileVarDec()
			if varDec != nil {
				node.AddChild(*varDec)
			}
		} else {
			break
		}
	}

	statements := compileStatements() // statements
	if statements != nil {
		node.AddChild(*statements)
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	return &node
}

func compileVarDec() *Node {
	// Grammar: 'var' type varName (',' varName)* ';'
	// Where type: 'int' | 'char' | 'boolean' | className

	node := Node{Name: "varDec", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeKeyword, []string{"var"}) // 'var'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	if token.TypeOf == tokenizer.TokenTypeKeyword {
		requireToken(token, tokenizer.TokenTypeKeyword, []string{"int", "char", "boolean"}) // type - 'int' | 'char' | 'boolean'
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++
	}
	if token.TypeOf == tokenizer.TokenTypeIdentifier {
		requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // type - className
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // varName
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	for token.TypeOf == tokenizer.TokenTypeSymbol && token.Value == "," { // (',' varName)*
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}}) // ','
		tokenIndex++

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // varName
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	return &node
}

func compileStatements() *Node {
	// Grammar: statement*
	// Where statement: doStatement | ifStatement | letStatement | returnStatement | whileStatement

	node := Node{Name: "statements", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]

	func() {
		for { // statement*
			switch token.Value {
			case "do":
				doStatement := compileDoStatement() // doStatement
				if doStatement != nil {
					node.AddChild(*doStatement)
				}
				token = tokens[tokenIndex]
			case "if":
				ifStatement := compileIfStatement() // ifStatement
				if ifStatement != nil {
					node.AddChild(*ifStatement)
				}
				token = tokens[tokenIndex]
			case "let":
				letStatement := compileLetStatement() // letStatement
				if letStatement != nil {
					node.AddChild(*letStatement)
				}
				token = tokens[tokenIndex]
			case "return":
				returnStatement := compileReturnStatement() // returnStatement
				if returnStatement != nil {
					node.AddChild(*returnStatement)
				}
				token = tokens[tokenIndex]
			case "while":
				whileStatement := compileWhileStatement() // whileStatement
				if whileStatement != nil {
					node.AddChild(*whileStatement)
				}
				token = tokens[tokenIndex]
			default:
				return // return from anonymous self-invoked function when there are no more statements
			}
		}
	}()

	return &node
}

func compileDoStatement() *Node {
	// Grammar: 'do' subroutineCall ';'
	// Where subroutineCall: subroutineName '(' expressionList ')' | (className | varName) '.' subroutineName '(' expressionList ')'

	node := Node{Name: "doStatement", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeKeyword, []string{"do"}) // 'do'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // subroutineName OR (className | varName)
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	if token.Value == "." { // the previous token was actually (className | varName), and should be followed by '.' subroutineName before continuing...
		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeSymbol, []string{"."}) // '.'
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // subroutineName
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"("}) // '('
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	expressionList := compileExpressionList() // expressionList
	if expressionList != nil {
		node.AddChild(*expressionList)
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{")"}) // ')'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	return &node
}

func compileIfStatement() *Node {
	// Grammar: 'if' '(' expression ')' '{' statements '}' ( 'else' '{' statements '}' )?

	node := Node{Name: "ifStatement", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeKeyword, []string{"if"}) // 'if'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"("}) // '('
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	expression := compileExpression() // expression
	if expression != nil {
		node.AddChild(*expression)
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{")"}) // ')'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	statements := compileStatements() // statements
	if statements != nil {
		node.AddChild(*statements)
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	if token.TypeOf == tokenizer.TokenTypeKeyword && token.Value == "else" { // there is an else...
		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeKeyword, []string{"else"}) // 'else'
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++

		statements = compileStatements() // statements
		if statements != nil {
			node.AddChild(*statements)
		}

		token = tokens[tokenIndex]
		requireToken(token, tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++
	}

	return &node
}

func compileLetStatement() *Node {
	// Grammar: 'let' varName ( '[' expression ']' )? '=' expression ';'

	node := Node{Name: "letStatement", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeKeyword, []string{"let"}) // 'let'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeIdentifier, []string{}) // varName
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	//TODO: ( '[' expression ']' )?

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"="}) // '='
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	expression := compileExpression() // expression
	if expression != nil {
		node.AddChild(*expression)
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	return &node
}

func compileReturnStatement() *Node {
	// Grammar: 'return' expression? ';'

	node := Node{Name: "returnStatement", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeKeyword, []string{"return"}) // 'return'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	if token.TypeOf != tokenizer.TokenTypeSymbol || token.Value != ";" { // no expression...
		expression := compileExpression() // expression
		if expression != nil {
			node.AddChild(*expression)
		}
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	return &node
}

func compileWhileStatement() *Node {
	// Grammar: 'while' '(' expression ')' '{' statements '}'

	node := Node{Name: "whileStatement", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeKeyword, []string{"while"}) // 'while'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"("}) // '('
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	expression := compileExpression() // expression
	if expression != nil {
		node.AddChild(*expression)
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{")"}) // ')'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	statements := compileStatements() // statements
	if statements != nil {
		node.AddChild(*statements)
	}

	token = tokens[tokenIndex]
	requireToken(token, tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
	node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
	tokenIndex++

	return &node
}

func compileExpressionList() *Node {
	// Grammar: (expression (',' expression)* )?

	node := Node{Name: "expressionList", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	if token.TypeOf == tokenizer.TokenTypeSymbol && token.Value == ")" { // no expressions...
		return &node
	} else { // first expression
		expression := compileExpression()
		if expression != nil {
			node.AddChild(*expression)
		}
	}

	token = tokens[tokenIndex]
	for token.TypeOf == tokenizer.TokenTypeSymbol && token.Value == "," { // (',' expression)*
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}}) // ','
		tokenIndex++

		token = tokens[tokenIndex]
		expression := compileExpression() // expression
		if expression != nil {
			node.AddChild(*expression)
		}

		token = tokens[tokenIndex]
	}

	return &node
}

func compileExpression() *Node {
	// Grammar: term (op term)*

	node := Node{Name: "expression", Value: "", Children: []Node{}}

	term := compileTerm() // term
	if term != nil {
		node.AddChild(*term)
	}

	token := tokens[tokenIndex]
	for isOp(token) { // (op term)*
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}}) // op
		tokenIndex++

		token = tokens[tokenIndex]
		term := compileTerm() // term
		if term != nil {
			node.AddChild(*term)
		}
	}

	return &node
}

func isOp(token tokenizer.Token) bool {
	if token.TypeOf != tokenizer.TokenTypeSymbol {
		return false
	}
	if token.Value != "+" && token.Value != "-" && token.Value != "*" && token.Value != "/" && token.Value != "&amp;" && token.Value != "|" && token.Value != "&lt;" && token.Value != "&gt;" && token.Value != "=" {
		return false
	}
	return true
}

func compileTerm() *Node {
	// Grammar: integerConstant | stringConstant | keywordConstant | varName | varName '[' expression ']' | subroutineCall | '(' expression ')' | unaryOp term

	node := Node{Name: "term", Value: "", Children: []Node{}}

	token := tokens[tokenIndex]
	if token.TypeOf == tokenizer.TokenTypeIdentifier || token.TypeOf == tokenizer.TokenTypeKeyword { // identifier or keyword (temporary expressionless workaround)
		node.AddChild(Node{Name: string(token.TypeOf), Value: token.Value, Children: []Node{}})
		tokenIndex++
	}

	return &node
}

func requireToken(token tokenizer.Token, allowableTypeOf tokenizer.TokenType, allowableValues []string) {
	if token.TypeOf != allowableTypeOf {
		handler.FatalError(errors.New(fmt.Sprintf("Unexpected token type: \"%s\" - allowable type: \"%s\"", token.TypeOf, allowableTypeOf)))
	}
	if len(allowableValues) > 0 {
		matched := false
		for _, allowableValue := range allowableValues {
			if token.Value == allowableValue {
				matched = true
			}
		}
		if !matched {
			handler.FatalError(errors.New(fmt.Sprintf("Unexpected token value: \"%s\" - allowable values: \"%s\"", token.Value, allowableValues)))
		}
	}
}
