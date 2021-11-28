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

const (
	SubroutineTypeConstructor = "constructor"
	SubroutineTypeFunction    = "function"
	SubroutineTypeMethod      = "method"
)

type CompilationEngine struct {
	tokens      []tokenizer.Token
	debug       bool
	vmComments  bool
	tokenIndex  int
	cst         *symbols.CombinedSymbolTable
	vmw         *writer.VMWriter
	className   string
	labelCounts map[string]int
}

// NewCompilationEngine - creates a compilation engine
func NewCompilationEngine(tokens []tokenizer.Token, debug, vmComments bool) *CompilationEngine {
	ce := CompilationEngine{
		tokens:      tokens,
		debug:       debug,
		vmComments:  vmComments,
		cst:         symbols.NewCombinedSymbolTable(),
		vmw:         writer.NewVMWriter(debug),
		labelCounts: map[string]int{},
	}

	if ce.debug {
		fmt.Println("CompilationEngine.NewCompilationEngine()", ce.cst)
	}

	return &ce
}

// CompileClass - compiles a class, returns in both xml and vm formats
func (ce *CompilationEngine) CompileClass() (string, string, string) {
	// Grammar: 'class' className '{' classVarDec* subroutineDec* '}'

	node := Node{Name: "class", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"class"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'class'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})

	node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // className
	ce.className = ce.currentToken().Value
	if ce.vmComments {
		ce.vmw.WriteComment(fmt.Sprintf("Class: %s", ce.className))
	}
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

	if ce.debug {
		fmt.Println("CompilationEngine.CompileClass()", ce.cst)
	}

	return node.AsXML(), node.AsExtendedXML(), ce.vmw.Code
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
	symbolIdentifierName := ce.currentToken().Value
	ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKind(symbolKindOf))                                                                                                                                         // add to combined symbol table
	node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // varName
	ce.advance()

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // (',' varName)*...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
		symbolIdentifierName := ce.currentToken().Value
		ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKind(symbolKindOf))                                                                                                                                         // add to combined symbol table
		node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // varName
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ';'
	ce.advance()

	if ce.debug {
		fmt.Println("CompilationEngine.compileClassVarDec()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) compileSubroutineDec() *Node {
	// Grammar: ('constructor' | 'function' | 'method') ('void' | type) subroutineName '(' parameterList ')' subroutineBody

	if !ce.currentToken().IsSubroutineDec() {
		return nil
	}

	ce.cst.StartSubroutine()          // clears subroutine scope symbol table
	ce.labelCounts = map[string]int{} // clears label counts

	node := Node{Name: "subroutineDec", Value: "", Children: []Node{}}

	childNode := Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}} // ('constructor' | 'function' | 'method')
	subroutineType := ce.currentToken().Value
	if subroutineType == SubroutineTypeMethod {
		ce.cst.Define("this", ce.className, symbols.SymbolKindArgument) // add to combined symbol table ("this" is always the first argument of methods only)
	}
	node.AddChild(childNode)
	ce.advance()

	if ce.currentToken().IsTypeOrVoid() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ('void' | type)
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // subroutineName
	subroutineName := ce.currentToken().Value
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

	subroutineBody := ce.compileSubroutineBody(subroutineName, subroutineType)
	if subroutineBody != nil {
		node.AddChild(*subroutineBody) // subroutineBody
	}

	if ce.debug {
		fmt.Println("CompilationEngine.compileSubroutineDec()", ce.cst)
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

	if ce.debug {
		fmt.Println("CompilationEngine.compileParameterList()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) addParameter(node *Node) *Node {
	symbolTypeOf := ""
	if ce.currentToken().IsType() {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // type
		symbolTypeOf = ce.currentToken().Value
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	symbolIdentifierName := ce.currentToken().Value
	ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKindArgument)                                                                                                                                               // add to combined symbol table
	node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // varName
	ce.advance()

	if ce.debug {
		fmt.Println("CompilationEngine.addParameter()", ce.cst)
	}

	return node
}

func (ce *CompilationEngine) compileSubroutineBody(subroutineName, subroutineType string) *Node {
	// Grammar: '{' varDec* statements '}'

	if ce.vmComments {
		ce.vmw.WriteComment(fmt.Sprintf("SubroutineBody: %s", subroutineName))
	}

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

	ce.vmw.WriteFunction(fmt.Sprintf("%s.%s", ce.className, subroutineName), ce.cst.GetVarCount(symbols.SymbolKindVar)) // vm write: function thisClassName.subroutineName n

	if subroutineType == SubroutineTypeConstructor {
		ce.vmw.WritePush(writer.SegmentConstant, ce.cst.GetVarCount(symbols.SymbolKindField)) // vm write: push constant n
		ce.vmw.WriteCall("Memory.alloc", 1)                                                   // vm write: call Memory.alloc 1
		ce.vmw.WritePop(writer.SegmentPointer, 0)                                             // vm write: pop pointer 0
	}

	if subroutineType == SubroutineTypeMethod {
		ce.vmw.WritePush(writer.SegmentArgument, 0) // vm write: push argument 0
		ce.vmw.WritePop(writer.SegmentPointer, 0)   // vm write: pop pointer 0
	}

	statements := ce.compileStatements()
	if statements != nil {
		node.AddChild(*statements) // statements
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"}"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '}'
	ce.advance()

	if ce.debug {
		fmt.Println("CompilationEngine.compileSubroutineBody()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) compileVarDec() *Node {
	// Grammar: 'var' type varName (',' varName)* ';'

	node := Node{Name: "varDec", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"var"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'var'
	ce.advance()

	symbolTypeOf := ""
	if ce.currentToken().IsType() {
		if ce.currentToken().TypeOf == tokenizer.TokenTypeIdentifier {
			node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // type (identifier)
		} else {
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // type (other)
		}
		symbolTypeOf = ce.currentToken().Value
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	symbolIdentifierName := ce.currentToken().Value
	ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKindVar)                                                                                                                                                    // add to combined symbol table
	node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // varName
	ce.advance()

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // (',' varName)*...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
		symbolIdentifierName := ce.currentToken().Value
		ce.cst.Define(symbolIdentifierName, symbolTypeOf, symbols.SymbolKindVar)                                                                                                                                                    // add to combined symbol table
		node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // varName
		ce.advance()
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ';'
	ce.advance()

	if ce.debug {
		fmt.Println("CompilationEngine.compileVarDec()", ce.cst)
	}

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

	if ce.debug {
		fmt.Println("CompilationEngine.compileStatements()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) compileDoStatement() *Node {
	// Grammar: 'do' subroutineCall ';'
	// Where subroutineCall: subroutineName '(' expressionList ')' | (className | varName) '.' subroutineName '(' expressionList ')'

	if ce.vmComments {
		ce.vmw.WriteComment("Do")
	}

	node := Node{Name: "doStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"do"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'do'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // subroutineCall: subroutineName OR (className | varName)
	classOrVarName := ce.currentToken().Value
	classOrVarType := ce.cst.GetTypeOf(classOrVarName)
	vmClassName := classOrVarName
	vmSubroutineName := ""
	vmNumArgs := 0
	ce.advance()

	if ce.currentToken().Value == "." { // the previous token was actually (className | varName), and should be followed by '.' and then subroutineName before continuing...
		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"."})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '.'
		ce.advance()

		ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
		node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // subroutineName
		vmSubroutineName = ce.currentToken().Value
		ce.advance()
	} else { // the previous token was actually the subroutineName...
		vmSubroutineName = vmClassName
		vmClassName = ce.className
		vmNumArgs++
		ce.vmw.WritePush(writer.SegmentPointer, 0) // vm write: push pointer n
	}

	if classOrVarType != "" {
		kind := ce.cst.GetKindOf(vmClassName)
		index := ce.cst.GetIndexOf(vmClassName)
		if kind == symbols.SymbolKindField {
			ce.vmw.WritePush(writer.SegmentThis, index) // vm write: push this n
		} else {
			ce.vmw.WritePush(writer.Segment(kind), index) // vm write: push segment n
		}
		vmNumArgs++
		vmClassName = classOrVarType
	}

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
	ce.advance()

	expressionList, numExpressionListArgs := ce.compileExpressionList()
	if expressionList != nil {
		node.AddChild(*expressionList) // expressionList
	}
	vmNumArgs += numExpressionListArgs

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{";"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ';'
	ce.advance()

	ce.vmw.WriteCall(fmt.Sprintf("%s.%s", vmClassName, vmSubroutineName), vmNumArgs) // vm write: call className.subroutineName n
	ce.vmw.WritePop(writer.SegmentTemp, 0)                                           // vm write: pop temp 0

	if ce.debug {
		fmt.Println("CompilationEngine.compileDoStatement()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) compileIfStatement() *Node {
	// Grammar: 'if' '(' expression ')' '{' statements '}' ( 'else' '{' statements '}' )?

	if ce.vmComments {
		ce.vmw.WriteComment("If")
	}

	node := Node{Name: "ifStatement", Value: "", Children: []Node{}}

	l1 := fmt.Sprintf("IF_TRUE%d", ce.labelCounts["IF_TRUE"])
	ce.labelCounts["IF_TRUE"]++
	l2 := fmt.Sprintf("IF_FALSE%d", ce.labelCounts["IF_FALSE"])
	ce.labelCounts["IF_FALSE"]++
	l3 := fmt.Sprintf("IF_END%d", ce.labelCounts["IF_END"])
	ce.labelCounts["IF_END"]++

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

	ce.vmw.WriteIf(l1)    // vm write: if-goto L1
	ce.vmw.WriteGoto(l2)  // vm write: goto L2
	ce.vmw.WriteLabel(l1) // vm write: label L1

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
		ce.vmw.WriteGoto(l3)  // vm write: goto L3
		ce.vmw.WriteLabel(l2) // vm write: label L2

		ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"else"})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'else'
		ce.advance()

		if ce.vmComments {
			ce.vmw.WriteComment("Else")
		}

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

		ce.vmw.WriteLabel(l3) // vm write: label L3
	} else {
		ce.vmw.WriteLabel(l2) // vm write: label L2
	}

	if ce.debug {
		fmt.Println("CompilationEngine.compileIfStatement()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) compileLetStatement() *Node {
	// Grammar: 'let' varName ( '[' expression ']' )? '=' expression ';'

	if ce.vmComments {
		ce.vmw.WriteComment("Let")
	}

	node := Node{Name: "letStatement", Value: "", Children: []Node{}}

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"let"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'let'
	ce.advance()

	ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
	node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // varName
	vmName := ce.currentToken().Value
	vmSegment := ce.cst.GetKindOf(vmName)
	vmIndex := ce.cst.GetIndexOf(vmName)
	ce.advance()

	isArray := false
	if ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "[" { // ( '[' expression ']' )?...
		isArray = true

		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '['
		ce.advance()

		expression := ce.compileExpression()
		if expression != nil {
			node.AddChild(*expression) // expression
		}

		ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"]"})
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ']'
		ce.advance()

		ce.vmw.WritePush(writer.GetSegment(string(vmSegment)), vmIndex) // vm write: push segment n
		ce.vmw.WriteArithmetic("add")                                   // vm write: add
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

	if isArray {
		ce.vmw.WritePop(writer.SegmentTemp, 0)    // vm write: pop temp 0
		ce.vmw.WritePop(writer.SegmentPointer, 1) // vm write: pop pointer 1
		ce.vmw.WritePush(writer.SegmentTemp, 0)   // vm write: push temp 0
		ce.vmw.WritePop(writer.SegmentThat, 0)    // vm write: pop that 0
	} else {
		if vmSegment == symbols.SymbolKindField {
			ce.vmw.WritePop(writer.SegmentThis, vmIndex) // vm write: pop this n
		} else {
			ce.vmw.WritePop(writer.GetSegment(string(vmSegment)), vmIndex) // vm write: pop segment n
		}
	}

	if ce.debug {
		fmt.Println("CompilationEngine.compileLetStatement()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) compileReturnStatement() *Node {
	// Grammar: 'return' expression? ';'

	if ce.vmComments {
		ce.vmw.WriteComment("Return")
	}

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

	if ce.debug {
		fmt.Println("CompilationEngine.compileReturnStatement()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) compileWhileStatement() *Node {
	// Grammar: 'while' '(' expression ')' '{' statements '}'

	if ce.vmComments {
		ce.vmw.WriteComment("While")
	}

	node := Node{Name: "whileStatement", Value: "", Children: []Node{}}

	l1 := fmt.Sprintf("WHILE_EXP%d", ce.labelCounts["WHILE_EXP"])
	ce.labelCounts["WHILE_EXP"]++
	l2 := fmt.Sprintf("WHILE_END%d", ce.labelCounts["WHILE_END"])
	ce.labelCounts["WHILE_END"]++

	ce.validateCurrentToken(tokenizer.TokenTypeKeyword, []string{"while"})
	node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // 'while'
	ce.advance()

	ce.vmw.WriteLabel(l1) // vm write: label L1

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

	ce.vmw.WriteArithmetic("not") // vm write: not
	ce.vmw.WriteIf(l2)            // vm write: if-goto L2

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

	ce.vmw.WriteGoto(l1)  // vm write: goto L1
	ce.vmw.WriteLabel(l2) // vm write: label L2

	if ce.debug {
		fmt.Println("CompilationEngine.compileWhileStatement()", ce.cst)
	}

	return &node
}

func (ce *CompilationEngine) compileExpressionList() (*Node, int) {
	// Grammar: (expression (',' expression)* )?

	node := Node{Name: "expressionList", Value: "", Children: []Node{}}

	numArgs := 0

	if ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == ")" { // no expressions...
		return &node, 0
	} else { // first expression...
		expression := ce.compileExpression()
		if expression != nil {
			node.AddChild(*expression) // expression
		}
		numArgs++
	}

	for ce.currentToken().TypeOf == tokenizer.TokenTypeSymbol && ce.currentToken().Value == "," { // subsequent expression(s)...
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ','
		ce.advance()

		expression := ce.compileExpression()
		if expression != nil {
			node.AddChild(*expression) // expression
		}
		numArgs++
	}

	if ce.debug {
		fmt.Println("CompilationEngine.compileExpressionList()", ce.cst)
	}

	return &node, numArgs
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

	if ce.debug {
		fmt.Println("CompilationEngine.compileExpression()", ce.cst)
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
		ce.vmw.WritePush(writer.SegmentConstant, intValue) // vm write: push constant n
		ce.advance()
	} else if ce.currentToken().TypeOf == tokenizer.TokenTypeStringConstant {
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // stringConstant
		ce.vmw.WritePush(writer.SegmentConstant, len(ce.currentToken().Value))                                          // vm write: push constant n
		ce.vmw.WriteCall("String.new", 1)                                                                               // vm write: call String.new 1
		for _, char := range ce.currentToken().Value {
			ce.vmw.WritePush(writer.SegmentConstant, int(char)) // vm write: push constant n
			ce.vmw.WriteCall("String.appendChar", 2)            // vm write: call String.appendChar 2
		}
		ce.advance()
	} else if ce.currentToken().IsKeywordConstant() {
		if ce.currentToken().Value == "true" {
			ce.vmw.WritePush(writer.SegmentConstant, 0) // vm write: push constant 0
			ce.vmw.WriteArithmetic("not")               // vm write: not
		}
		if ce.currentToken().Value == "false" || ce.currentToken().Value == "null" {
			ce.vmw.WritePush(writer.SegmentConstant, 0) // vm write: push constant 0
		}
		if ce.currentToken().Value == "this" {
			ce.vmw.WritePush(writer.SegmentPointer, 0) // vm write: push pointer 0
		}
		node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // keywordConstant
		ce.advance()
	} else if ce.currentToken().TypeOf == tokenizer.TokenTypeIdentifier {
		if ce.nextToken().TypeOf == tokenizer.TokenTypeSymbol && ce.nextToken().Value == "(" { // subroutineCall: subroutineName '(' expressionList ')'...
			node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // subroutineName
			vmName := ce.currentToken().Value
			ce.advance()

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
			ce.advance()

			expressionList, numArgs := ce.compileExpressionList()
			if expressionList != nil {
				node.AddChild(*expressionList) // expressionList
			}

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
			ce.advance()

			ce.vmw.WriteCall(vmName, numArgs) // vm write: call className.subroutineName n
		} else if ce.nextToken().TypeOf == tokenizer.TokenTypeSymbol && ce.nextToken().Value == "." { // subroutineCall: (className | varName) '.' subroutineName '(' expressionList ')'...

			node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // (className | varName)
			classOrVarName := ce.currentToken().Value
			classOrVarType := ce.cst.GetTypeOf(classOrVarName)
			vmClassName := classOrVarName
			vmSubroutineName := ""
			vmNumArgs := 0
			ce.advance()

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"."})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '.'
			ce.advance()

			ce.validateCurrentToken(tokenizer.TokenTypeIdentifier, []string{})
			node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // subroutineName
			vmSubroutineName = ce.currentToken().Value
			ce.advance()

			if classOrVarType != "" {
				kind := ce.cst.GetKindOf(vmClassName)
				index := ce.cst.GetIndexOf(vmClassName)
				if kind == symbols.SymbolKindField {
					ce.vmw.WritePush(writer.SegmentThis, index) // vm write: push this n
				} else {
					ce.vmw.WritePush(writer.Segment(kind), index) // vm write: push segment n
				}
				vmNumArgs++
				vmClassName = classOrVarType
			}

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{"("})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // '('
			ce.advance()

			expressionList, numExpressionListArgs := ce.compileExpressionList()
			if expressionList != nil {
				node.AddChild(*expressionList) // expressionList
			}
			vmNumArgs += numExpressionListArgs

			ce.validateCurrentToken(tokenizer.TokenTypeSymbol, []string{")"})
			node.AddChild(Node{Name: string(ce.currentToken().TypeOf), Value: ce.currentToken().Value, Children: []Node{}}) // ')'
			ce.advance()

			ce.vmw.WriteCall(fmt.Sprintf("%s.%s", vmClassName, vmSubroutineName), vmNumArgs) // vm write: call className.subroutineName n
		} else if ce.nextToken().TypeOf == tokenizer.TokenTypeSymbol && ce.nextToken().Value == "[" { // varName '[' expression ']'...
			node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // varName
			vmName := ce.currentToken().Value
			vmSegment := ce.cst.GetKindOf(vmName)
			vmIndex := ce.cst.GetIndexOf(vmName)
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

			ce.vmw.WritePush(writer.GetSegment(string(vmSegment)), vmIndex) // vm write: push segment n
			ce.vmw.WriteArithmetic("add")                                   // vm write: add
			ce.vmw.WritePop(writer.SegmentPointer, 1)                       // vm write: pop pointer 1
			ce.vmw.WritePush(writer.SegmentThat, 0)                         // vm write: push that 0
		} else { // varName...
			node.AddChild(Node{Name: string(tokenizer.TokenTypeIdentifier), Value: ce.currentToken().Value, Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"}) // varName
			vmName := ce.currentToken().Value
			vmSegment := ce.cst.GetKindOf(vmName)
			vmIndex := ce.cst.GetIndexOf(vmName)
			ce.advance()

			if vmSegment == symbols.SymbolKindField {
				ce.vmw.WritePush(writer.SegmentThis, vmIndex) // vm write: push this n
			} else {
				ce.vmw.WritePush(writer.GetSegment(string(vmSegment)), vmIndex) // vm write: push segment n
			}
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
		unaryOp := ce.currentToken().Value
		ce.advance()

		term := ce.compileTerm()
		if term != nil {
			node.AddChild(*term) // term
		}

		if unaryOp == "-" {
			ce.vmw.WriteArithmetic("neg") // vm write: neg
		}
		if unaryOp == "~" {
			ce.vmw.WriteArithmetic("not") // vm write: not
		}
	}

	if ce.debug {
		fmt.Println("CompilationEngine.compileTerm()", ce.cst)
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
