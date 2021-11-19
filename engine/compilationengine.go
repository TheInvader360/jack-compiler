package engine

import (
	"fmt"
	"strconv"

	"github.com/TheInvader360/jack-compiler/handler"
	"github.com/TheInvader360/jack-compiler/symbols"
	"github.com/TheInvader360/jack-compiler/tokenizer"
	"github.com/TheInvader360/jack-compiler/writer"

	"github.com/pkg/errors"
)

type CompilationEngine struct {
	tokens     []tokenizer.Token
	tokenIndex int
	cst        *symbols.CombinedSymbolTable
	vmw        *writer.VMWriter
	className  string
}

// NewCompilationEngine - creates a compilation engine
func NewCompilationEngine(tokens []tokenizer.Token) *CompilationEngine {
	ce := CompilationEngine{
		tokens: tokens,
		cst:    symbols.NewCombinedSymbolTable(),
		vmw:    writer.NewVMWriter(),
	}

	return &ce
}

// CompileClass - compiles a class, returns in both xml and vm formats
func (ce *CompilationEngine) CompileClass() (string, string) {
	// Grammar: 'class' className '{' classVarDec* subroutineDec* '}'

	node := Node{Name: "class", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"class"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'class'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // className
	ce.className = ce.currentToken().Value
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '{'
	ce.advance()

	for { // classVarDec*...
		classVarDec := ce.compileClassVarDec()
		if classVarDec != nil {
			node.AddChild(*classVarDec) // classVarDec
		} else {
			break
		}
	}

	for { // subroutineDec*...
		subroutineDec := ce.compileSubroutineDec()
		if subroutineDec != nil {
			node.AddChild(*subroutineDec) // subroutineDec
		} else {
			break
		}
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '}'
	ce.advance()

	return node.AsXML(), ce.vmw.Code
}

func (ce *CompilationEngine) compileClassVarDec() *Node {
	// Grammar: ('static' | 'field') type varName (',' varName)* ';'

	if !ce.currentToken().IsClassVarDec() {
		return nil
	}

	node := Node{Name: "classVarDec", Value: "", Children: []Node{}}

	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ('static' | 'field')
	symbolKindOf := ce.currentToken().Value
	ce.advance()

	symbolTypeOf := ""
	if ce.currentToken().IsType() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // type
		symbolTypeOf = ce.currentToken().Value
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // varName
	symbolIdentifierName := ce.currentToken().Value
	ce.advance()

	ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKind(symbolKindOf)) // add to combined symbol table

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // (',' varName)*...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // varName
		symbolIdentifierName = ce.currentToken().Value
		ce.advance()

		ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKind(symbolKindOf)) // add to combined symbol table
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ';'
	ce.advance()

	//fmt.Println("compileClassVarDec()", ce.cst)

	return &node
}

func (ce *CompilationEngine) compileSubroutineDec() *Node {
	// Grammar: ('constructor' | 'function' | 'method') ('void' | type) subroutineName '(' parameterList ')' subroutineBody

	if !ce.currentToken().IsSubroutineDec() {
		return nil
	}

	ce.cst.StartSubroutine()

	node := Node{Name: "subroutineDec", Value: "", Children: []Node{}}

	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ('constructor' | 'function' | 'method')
	ce.advance()

	if ce.currentToken().IsTypeOrVoid() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ('void' | type)
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // subroutineName

	ce.vmw.WriteFunction(fmt.Sprintf("%s.%s", ce.className, ce.currentToken().Value), ce.cst.GetVarCount(symbols.SymbolKindVar)) // vm write: function thisClassName.subroutineName 'n'

	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
	ce.advance()

	parameterList := ce.compileParameterList()
	if parameterList != nil {
		node.AddChild(*parameterList) // parameterList
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
	ce.advance()

	subroutineBody := ce.compileSubroutineBody()
	if subroutineBody != nil {
		node.AddChild(*subroutineBody) // subroutineBody
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

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // varName
	ce.advance()

	return node
}

func (ce *CompilationEngine) compileSubroutineBody() *Node {
	// Grammar: '{' varDec* statements '}'

	node := Node{Name: "subroutineBody", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '{'
	ce.advance()

	for { // varDec*...
		if ce.currentToken().Value == "var" {
			varDec := ce.compileVarDec()
			if varDec != nil {
				node.AddChild(*varDec) // varDec
			}
		} else {
			break
		}
	}

	statements := ce.compileStatements()
	if statements != nil {
		node.AddChild(*statements) // statements
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '}'
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileVarDec() *Node {
	// Grammar: 'var' type varName (',' varName)* ';'

	node := Node{Name: "varDec", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"var"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'var'
	symbolKindOf := ce.currentToken().Value
	ce.advance()

	symbolTypeOf := ""
	if ce.currentToken().IsType() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // type
		symbolTypeOf = ce.currentToken().Value
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // varName
	symbolIdentifierName := ce.currentToken().Value
	ce.advance()

	ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKind(symbolKindOf)) // add to combined symbol table

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // (',' varName)*...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // varName
		symbolIdentifierName = ce.currentToken().Value
		ce.advance()

		ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKind(symbolKindOf)) // add to combined symbol table
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ';'
	ce.advance()

	//fmt.Println("compileVarDec()", ce.cst)

	return &node
}

func (ce *CompilationEngine) compileStatements() *Node {
	// Grammar: statement*
	// Where statement: doStatement | ifStatement | letStatement | returnStatement | whileStatement

	node := Node{Name: "statements", Value: "", Children: []Node{}}

	func() {
		for { // statement*...
			switch ce.currentToken().Value {
			case "do":
				doStatement := ce.compileDoStatement()
				if doStatement != nil {
					node.AddChild(*doStatement) // doStatement
				}
			case "if":
				ifStatement := ce.compileIfStatement()
				if ifStatement != nil {
					node.AddChild(*ifStatement) // ifStatement
				}
			case "let":
				letStatement := ce.compileLetStatement()
				if letStatement != nil {
					node.AddChild(*letStatement) // letStatement
				}
			case "return":
				returnStatement := ce.compileReturnStatement()
				if returnStatement != nil {
					node.AddChild(*returnStatement) // returnStatement
				}
			case "while":
				whileStatement := ce.compileWhileStatement()
				if whileStatement != nil {
					node.AddChild(*whileStatement) // whileStatement
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

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"do"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'do'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // subroutineCall: subroutineName OR (className | varName)
	vmName := ce.currentToken().Value
	ce.advance()

	if ce.currentToken().Value == "." { // the previous token was actually (className | varName), and should be followed by '.' and then subroutineName before continuing...
		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"."})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '.'
		vmName += "."
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // subroutineName
		vmName += ce.currentToken().Value
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
	ce.advance()

	expressionList := ce.compileExpressionList()
	if expressionList != nil {
		node.AddChild(*expressionList) // expressionList
	}
	expressionCount := len(expressionList.Children)

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ';'
	ce.advance()

	ce.vmw.WriteCall(vmName, expressionCount) // vm write: call 'vmName' 'n'
	ce.vmw.WritePop(writer.SegmentTemp, 0)    // vm write: pop temp 0

	return &node
}

func (ce *CompilationEngine) compileIfStatement() *Node {
	// Grammar: 'if' '(' expression ')' '{' statements '}' ( 'else' '{' statements '}' )?

	node := Node{Name: "ifStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"if"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'if'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
	ce.advance()

	expression := ce.compileExpression()
	if expression != nil {
		node.AddChild(*expression) // expression
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '{'
	ce.advance()

	statements := ce.compileStatements()
	if statements != nil {
		node.AddChild(*statements) // statements
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '}'
	ce.advance()

	if ce.currentToken().TypeOf == tokenizer.TokenTypeKeyword && ce.currentToken().Value == "else" { // there is an else...
		ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"else"})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'else'
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '{'
		ce.advance()

		statements = ce.compileStatements()
		if statements != nil {
			node.AddChild(*statements) // statements
		}

		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '}'
		ce.advance()
	}

	return &node
}

func (ce *CompilationEngine) compileLetStatement() *Node {
	// Grammar: 'let' varName ( '[' expression ']' )? '=' expression ';'

	node := Node{Name: "letStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"let"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'let'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // varName
	ce.advance()

	if ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "[" { // ( '[' expression ']' )?...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '['
		ce.advance()

		expression := ce.compileExpression()
		if expression != nil {
			node.AddChild(*expression) // expression
		}

		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"]"})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ']'
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"="})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '='
	ce.advance()

	expression := ce.compileExpression()
	if expression != nil {
		node.AddChild(*expression) // expression
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ';'
	ce.advance()

	return &node
}

func (ce *CompilationEngine) compileReturnStatement() *Node {
	// Grammar: 'return' expression? ';'

	node := Node{Name: "returnStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"return"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'return'
	ce.advance()

	if ce.currentToken().TypeOf != tokenizer.TokenTypeSymbol || ce.currentToken().Value != ";" { // expression?...
		expression := ce.compileExpression()
		if expression != nil {
			node.AddChild(*expression) // expression
		}
	} else {
		ce.vmw.WritePush(writer.SegmentConstant, 0) // vm write: push constant 0 (handle void - callee returns 0, caller discards the 0 value...)
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ';'
	ce.advance()

	ce.vmw.WriteReturn() // vm write: return

	return &node
}

func (ce *CompilationEngine) compileWhileStatement() *Node {
	// Grammar: 'while' '(' expression ')' '{' statements '}'

	node := Node{Name: "whileStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"while"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'while'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
	ce.advance()

	expression := ce.compileExpression()
	if expression != nil {
		node.AddChild(*expression) // expression
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"{"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '{'
	ce.advance()

	statements := ce.compileStatements()
	if statements != nil {
		node.AddChild(*statements) // statements
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '}'
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

		expression := ce.compileExpression()
		if expression != nil {
			node.AddChild(*expression) // expression
		}
	}

	return &node
}

func (ce *CompilationEngine) compileExpression() *Node {
	// Grammar: term (op term)*

	node := Node{Name: "expression", Value: "", Children: []Node{}}

	term := ce.compileTerm()
	if term != nil {
		node.AddChild(*term) // term
	}

	ops := []string{}
	for ce.currentToken().IsOp() { // (op term)*...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // op
		ops = append(ops, ce.currentToken().Value)
		ce.advance()

		term := ce.compileTerm()
		if term != nil {
			node.AddChild(*term) // term
		}
	}

	for _, op := range ops {
		ce.vmw.WriteArithmetic(op) // vm write: (arithmetic command)
	}

	return &node
}

func (ce *CompilationEngine) compileTerm() *Node {
	// Grammar: integerConstant | stringConstant | keywordConstant | subroutineCall | varName '[' expression ']' | varName | '(' expression ')' | unaryOp term
	// Where keywordConstant: 'true' | 'false' | 'null' | 'this'
	// Where subroutineCall: subroutineName '(' expressionList ')' | (className | varName) '.' subroutineName '(' expressionList ')'
	// Where unaryOp: '-' | ' ~ '

	node := Node{Name: "term", Value: "", Children: []Node{}}

	if ce.currentToken().TypeOf == tokenizer.TokenTypeIntegerConstant {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // integerConstant
		intValue, _ := strconv.Atoi(ce.currentToken().Value)
		ce.vmw.WritePush(writer.SegmentConstant, intValue) // vm write: push constant 'n'
		ce.advance()
	} else if ce.currentToken().TypeOf == tokenizer.TokenTypeStringConstant {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // stringConstant
		ce.advance()
	} else if ce.currentToken().IsKeywordConstant() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // keywordConstant
		ce.advance()
	} else if ce.currentToken().TypeOf == tokenizer.TokenTypeIdentifier {
		if ce.nextToken().TypeOf == tokenizer.TokenTypeSymbol && ce.nextToken().Value == "(" { // subroutineCall: subroutineName '(' expressionList ')'...
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // subroutineName
			ce.advance()

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
			ce.advance()

			expressionList := ce.compileExpressionList()
			if expressionList != nil {
				node.AddChild(*expressionList) // expressionList
			}

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
			ce.advance()
		} else if ce.nextToken().TypeOf == tokenizer.TokenTypeSymbol && ce.nextToken().Value == "." { // subroutineCall: (className | varName) '.' subroutineName '(' expressionList ')'...
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // (className | varName)
			ce.advance()

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"."})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '.'
			ce.advance()

			ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // subroutineName
			ce.advance()

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
			ce.advance()

			expressionList := ce.compileExpressionList()
			if expressionList != nil {
				node.AddChild(*expressionList) // expressionList
			}

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
			ce.advance()
		} else if ce.nextToken().TypeOf == tokenizer.TokenTypeSymbol && ce.nextToken().Value == "[" { // varName '[' expression ']'...
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // varName
			ce.advance()

			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // [
			ce.advance()

			expression := ce.compileExpression()
			if expression != nil {
				node.AddChild(*expression) // expression
			}

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"]"})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ']'
			ce.advance()
		} else { // varName...
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // varName
			ce.advance()
		}
	} else if ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "(" { // '(' expression ')'...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
		ce.advance()

		expression := ce.compileExpression()
		if expression != nil {
			node.AddChild(*expression) // expression
		}

		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
		ce.advance()
	} else if ce.currentToken().IsUnaryOp() { // unaryOp term...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // unaryOp
		ce.advance()

		term := ce.compileTerm()
		if term != nil {
			node.AddChild(*term) // term
		}
	}

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

func (ce *CompilationEngine) nextToken() *tokenizer.Token {
	return &ce.tokens[ce.tokenIndex+1]
}

func (ce *CompilationEngine) advance() {
	ce.tokenIndex++
}
