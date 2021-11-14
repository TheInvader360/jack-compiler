package engine

import (
	"fmt"

	"github.com/TheInvader360/jack-compiler/handler"
	"github.com/TheInvader360/jack-compiler/tokenizer"

	"github.com/pkg/errors"
)

type CompilationEngine struct {
	tokens     []tokenizer.Token
	tokenIndex int
}

// NewCompilationEngine - Creates a compilation engine
func NewCompilationEngine(tokens []tokenizer.Token) *CompilationEngine {
	ce := CompilationEngine{
		tokens: tokens,
	}

	return &ce
}

func (ce *CompilationEngine) CompileClass() string {
	// Grammar: 'class' className '{' classVarDec* subroutineDec* '}'

	node := Node{Name: "class", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"class"}) // 'class'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // className
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	for { // classVarDec*
		classVarDec := ce.compileClassVarDec()
		if classVarDec != nil {
			node.AddChild(*classVarDec)
		} else {
			break
		}
	}

	for { // subroutineDec*
		subroutineDec := ce.compileSubroutineDec()
		if subroutineDec != nil {
			node.AddChild(*subroutineDec)
		} else {
			break
		}
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return node.AsXML()
}

func (ce *CompilationEngine) compileClassVarDec() *Node {
	// Grammar: ('static' | 'field') type varName (',' varName)* ';'

	if !ce.currentToken().IsClassVarDec() {
		return nil
	}

	node := Node{Name: "classVarDec", Value: "", Children: []Node{}}

	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ('static' | 'field')
	ce.advance()

	if ce.currentToken().IsType() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // type
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // varName
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // (',' varName)*
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // varName
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileSubroutineDec() *Node {
	// Grammar: ('constructor' | 'function' | 'method') ('void' | type) subroutineName '(' parameterList ')' subroutineBody

	if !ce.currentToken().IsSubroutineDec() {
		return nil
	}

	node := Node{Name: "subroutineDec", Value: "", Children: []Node{}}

	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ('constructor' | 'function' | 'method')
	ce.advance()

	if ce.currentToken().IsTypeOrVoid() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ('void' | type)
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // subroutineName
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("}) // '('
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	parameterList := ce.compileParameterList() // parameterList
	if parameterList != nil {
		node.AddChild(*parameterList)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"}) // ')'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	subroutineBody := ce.compileSubroutineBody() // subroutineBody
	if subroutineBody != nil {
		node.AddChild(*subroutineBody)
	}

	return &node
}

func (ce *CompilationEngine) compileParameterList() *Node {
	// Grammar: ((type varName) (',' type varName)*)?

	node := Node{Name: "parameterList", Value: "", Children: []Node{}}

	if ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == ")" { // no parameters...
		return &node
	} else { // first parameter...
		ce.addParameter(&node) // (type varName)
	}

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // subsequent parameter(s)...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		ce.addParameter(&node) // (type varName)
	}

	return &node
}

func (ce *CompilationEngine) addParameter(node *Node) *Node {
	if ce.currentToken().IsType() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // type
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // varName
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return node
}

func (ce *CompilationEngine) compileSubroutineBody() *Node {
	// Grammar: '{' varDec* statements '}'

	node := Node{Name: "subroutineBody", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	for { // varDec*
		if ce.currentToken().Value == "var" {
			varDec := ce.compileVarDec()
			if varDec != nil {
				node.AddChild(*varDec)
			}
		} else {
			break
		}
	}

	statements := ce.compileStatements() // statements
	if statements != nil {
		node.AddChild(*statements)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileVarDec() *Node {
	// Grammar: 'var' type varName (',' varName)* ';'

	node := Node{Name: "varDec", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"var"}) // 'var'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	if ce.currentToken().IsType() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // type
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // varName
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // (',' varName)*
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // varName
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileStatements() *Node {
	// Grammar: statement*
	// Where statement: doStatement | ifStatement | letStatement | returnStatement | whileStatement

	node := Node{Name: "statements", Value: "", Children: []Node{}}

	func() {
		for { // statement*
			switch ce.currentToken().Value {
			case "do":
				doStatement := ce.compileDoStatement() // doStatement
				if doStatement != nil {
					node.AddChild(*doStatement)
				}
			case "if":
				ifStatement := ce.compileIfStatement() // ifStatement
				if ifStatement != nil {
					node.AddChild(*ifStatement)
				}
			case "let":
				letStatement := ce.compileLetStatement() // letStatement
				if letStatement != nil {
					node.AddChild(*letStatement)
				}
			case "return":
				returnStatement := ce.compileReturnStatement() // returnStatement
				if returnStatement != nil {
					node.AddChild(*returnStatement)
				}
			case "while":
				whileStatement := ce.compileWhileStatement() // whileStatement
				if whileStatement != nil {
					node.AddChild(*whileStatement)
				}
			default:
				return // return from anonymous self-invoked function when there are no more statements
			}
		}
	}()

	return &node
}

func (ce *CompilationEngine) compileDoStatement() *Node {
	// Grammar: 'do' subroutineCall ';'
	// Where subroutineCall: subroutineName '(' expressionList ')' | (className | varName) '.' subroutineName '(' expressionList ')'

	node := Node{Name: "doStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"do"}) // 'do'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // subroutineName OR (className | varName)
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	if ce.currentToken().Value == "." { // the previous token was actually (className | varName), and should be followed by '.' subroutineName before continuing...
		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"."}) // '.'
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // subroutineName
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("}) // '('
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	expressionList := ce.compileExpressionList() // expressionList
	if expressionList != nil {
		node.AddChild(*expressionList)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"}) // ')'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileIfStatement() *Node {
	// Grammar: 'if' '(' expression ')' '{' statements '}' ( 'else' '{' statements '}' )?

	node := Node{Name: "ifStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"if"}) // 'if'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("}) // '('
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	expression := ce.compileExpression() // expression
	if expression != nil {
		node.AddChild(*expression)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"}) // ')'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	statements := ce.compileStatements() // statements
	if statements != nil {
		node.AddChild(*statements)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	if ce.currentToken().TypeOf == tokenizer.TokenTypeKeyword && ce.currentToken().Value == "else" { // there is an else...
		ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"else"}) // 'else'
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
		ce.advance()

		statements = ce.compileStatements() // statements
		if statements != nil {
			node.AddChild(*statements)
		}

		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
		ce.advance()
	}

	return &node
}

func (ce *CompilationEngine) compileLetStatement() *Node {
	// Grammar: 'let' varName ( '[' expression ']' )? '=' expression ';'

	node := Node{Name: "letStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"let"}) // 'let'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{}) // varName
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	//TODO: ( '[' expression ']' )?

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"="}) // '='
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	expression := ce.compileExpression() // expression
	if expression != nil {
		node.AddChild(*expression)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileReturnStatement() *Node {
	// Grammar: 'return' expression? ';'

	node := Node{Name: "returnStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"return"}) // 'return'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	if ce.currentToken().TypeOf != tokenizer.TokenTypeSymbol || ce.currentToken().Value != ";" {
		expression := ce.compileExpression() // expression
		if expression != nil {
			node.AddChild(*expression)
		}
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"}) // ';'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileWhileStatement() *Node {
	// Grammar: 'while' '(' expression ')' '{' statements '}'

	node := Node{Name: "whileStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"while"}) // 'while'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("}) // '('
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	expression := ce.compileExpression() // expression
	if expression != nil {
		node.AddChild(*expression)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"}) // ')'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"}) // '{'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	statements := ce.compileStatements() // statements
	if statements != nil {
		node.AddChild(*statements)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"}) // '}'
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileExpressionList() *Node {
	// Grammar: (expression (',' expression)* )?

	node := Node{Name: "expressionList", Value: "", Children: []Node{}}

	if ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == ")" { // no expressions...
		return &node
	} else { // first expression...
		expression := ce.compileExpression()
		if expression != nil {
			node.AddChild(*expression) // expression
		}
	}

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // subsequent expression(s)...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		expression := ce.compileExpression() // expression
		if expression != nil {
			node.AddChild(*expression)
		}
	}

	return &node
}

func (ce *CompilationEngine) compileExpression() *Node {
	// Grammar: term (op term)*

	node := Node{Name: "expression", Value: "", Children: []Node{}}

	term := ce.compileTerm() // term
	if term != nil {
		node.AddChild(*term)
	}

	for ce.currentToken().IsOp() { // (op term)*
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // op
		ce.advance()

		term := ce.compileTerm() // term
		if term != nil {
			node.AddChild(*term)
		}
	}

	return &node
}

func (ce *CompilationEngine) compileTerm() *Node {
	// Grammar: integerConstant | stringConstant | keywordConstant | varName | varName '[' expression ']' | subroutineCall | '(' expression ')' | unaryOp term

	node := Node{Name: "term", Value: "", Children: []Node{}}

	if ce.currentToken().TypeOf == tokenizer.TokenTypeIdentifier || ce.currentToken().TypeOf == tokenizer.TokenTypeKeyword { // identifier or keyword (temporary expressionless workaround)
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}})
		ce.advance()
	}
	//TODO: full implementation...

	return &node
}

func (ce *CompilationEngine) validateCurrentToken(allowableTypeOf tokenizer.TokenType, allowableValues []string) {
	if ce.currentToken().TypeOf != allowableTypeOf {
		handler.FatalError(errors.New(fmt.Sprintf("Unexpected token type: \"%s\" - allowable type: \"%s\"", ce.currentToken().TypeOf, allowableTypeOf)))
	}
	if len(allowableValues) > 0 {
		matched := false
		for _, allowableValue := range allowableValues {
			if ce.currentToken().Value == allowableValue {
				matched = true
			}
		}
		if !matched {
			handler.FatalError(errors.New(fmt.Sprintf("Unexpected token value: \"%s\" - allowable values: \"%s\"", ce.currentToken().Value, allowableValues)))
		}
	}
}

func (ce *CompilationEngine) currentToken() *tokenizer.Token {
	return &ce.tokens[ce.tokenIndex]
}

func (ce *CompilationEngine) advance() {
	ce.tokenIndex++
}
