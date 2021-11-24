package engine

import (
	"testing"
)

func Array_Main_Tree(t *testing.T) Node {
	return Node{Name: "class", Value: "", Children: []Node{
		{Name: "keyword", Value: "class", Children: []Node{}},
		{Name: "identifier", Value: "Main", Children: []Node{}},
		{Name: "symbol", Value: "{", Children: []Node{}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "function", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "main", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "identifier", Value: "Array", Children: []Node{}},
					{Name: "identifier", Value: "a", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "keyword", Value: "int", Children: []Node{}},
					{Name: "identifier", Value: "length", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "keyword", Value: "int", Children: []Node{}},
					{Name: "identifier", Value: "i", Children: []Node{}},
					{Name: "symbol", Value: ",", Children: []Node{}},
					{Name: "identifier", Value: "sum", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "length", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Keyboard", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "readInt", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "stringConstant", Value: "HOW MANY NUMBERS? ", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "a", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Array", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "new", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "length", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "i", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "0", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "whileStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "while", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "i", Children: []Node{}},
							}},
							{Name: "symbol", Value: "&lt;", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "length", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "a", Children: []Node{}},
								{Name: "symbol", Value: "[", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "i", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "]", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "Keyboard", Children: []Node{}},
										{Name: "symbol", Value: ".", Children: []Node{}},
										{Name: "identifier", Value: "readInt", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{
											{Name: "expression", Value: "", Children: []Node{
												{Name: "term", Value: "", Children: []Node{
													{Name: "stringConstant", Value: "ENTER THE NEXT NUMBER: ", Children: []Node{}},
												}},
											}},
										}},
										{Name: "symbol", Value: ")", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "i", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "i", Children: []Node{}},
									}},
									{Name: "symbol", Value: "+", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "1", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "i", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "0", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "sum", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "0", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "whileStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "while", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "i", Children: []Node{}},
							}},
							{Name: "symbol", Value: "&lt;", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "length", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "sum", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "sum", Children: []Node{}},
									}},
									{Name: "symbol", Value: "+", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "a", Children: []Node{}},
										{Name: "symbol", Value: "[", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "i", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: "]", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "i", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "i", Children: []Node{}},
									}},
									{Name: "symbol", Value: "+", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "1", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Output", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "printString", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "stringConstant", Value: "THE AVERAGE IS: ", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Output", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "printInt", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "sum", Children: []Node{}},
								}},
								{Name: "symbol", Value: "/", Children: []Node{}},
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "length", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Output", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "println", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "symbol", Value: "}", Children: []Node{}},
	}}
}

func HelloWorld_Main_Tree(t *testing.T) Node {
	return Node{Name: "class", Value: "", Children: []Node{
		{Name: "keyword", Value: "class", Children: []Node{}},
		{Name: "identifier", Value: "Main", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
		{Name: "symbol", Value: "{", Children: []Node{}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "function", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "main", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Output", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "printString", Children: []Node{}, IdentifierCategory: "TODO", IdentifierAction: "TODO", IdentifierKind: "TODO", IdentifierIndex: "TODO"},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "stringConstant", Value: "Hello World", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "symbol", Value: "}", Children: []Node{}},
	}}
}

func Square_Main_Tree(t *testing.T) Node {
	return Node{Name: "class", Value: "", Children: []Node{
		{Name: "keyword", Value: "class", Children: []Node{}},
		{Name: "identifier", Value: "Main", Children: []Node{}},
		{Name: "symbol", Value: "{", Children: []Node{}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "static", Children: []Node{}},
			{Name: "keyword", Value: "boolean", Children: []Node{}},
			{Name: "identifier", Value: "test", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "function", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "main", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "identifier", Value: "SquareGame", Children: []Node{}},
					{Name: "identifier", Value: "game", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "game", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "SquareGame", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "new", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "game", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "run", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "game", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "dispose", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "function", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "more", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "keyword", Value: "int", Children: []Node{}},
					{Name: "identifier", Value: "i", Children: []Node{}},
					{Name: "symbol", Value: ",", Children: []Node{}},
					{Name: "identifier", Value: "j", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "identifier", Value: "String", Children: []Node{}},
					{Name: "identifier", Value: "s", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "identifier", Value: "Array", Children: []Node{}},
					{Name: "identifier", Value: "a", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "keyword", Value: "false", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "s", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "stringConstant", Value: "string constant", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "s", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "keyword", Value: "null", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "a", Children: []Node{}},
								{Name: "symbol", Value: "[", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "1", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "]", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "a", Children: []Node{}},
										{Name: "symbol", Value: "[", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "integerConstant", Value: "2", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: "]", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
						{Name: "keyword", Value: "else", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "i", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "i", Children: []Node{}},
									}},
									{Name: "symbol", Value: "*", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "symbol", Value: "-", Children: []Node{}},
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "j", Children: []Node{}},
												}},
											}},
										}},
										{Name: "symbol", Value: ")", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "j", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "j", Children: []Node{}},
									}},
									{Name: "symbol", Value: "/", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "symbol", Value: "-", Children: []Node{}},
												{Name: "term", Value: "", Children: []Node{
													{Name: "integerConstant", Value: "2", Children: []Node{}},
												}},
											}},
										}},
										{Name: "symbol", Value: ")", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "i", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "i", Children: []Node{}},
									}},
									{Name: "symbol", Value: "|", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "j", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "symbol", Value: "}", Children: []Node{}},
	}}
}

func Square_Square_Tree(t *testing.T) Node {
	return Node{Name: "class", Value: "", Children: []Node{
		{Name: "keyword", Value: "class", Children: []Node{}},
		{Name: "identifier", Value: "Square", Children: []Node{}},
		{Name: "symbol", Value: "{", Children: []Node{}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "field", Children: []Node{}},
			{Name: "keyword", Value: "int", Children: []Node{}},
			{Name: "identifier", Value: "x", Children: []Node{}},
			{Name: "symbol", Value: ",", Children: []Node{}},
			{Name: "identifier", Value: "y", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "field", Children: []Node{}},
			{Name: "keyword", Value: "int", Children: []Node{}},
			{Name: "identifier", Value: "size", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "constructor", Children: []Node{}},
			{Name: "identifier", Value: "Square", Children: []Node{}},
			{Name: "identifier", Value: "new", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{
				{Name: "keyword", Value: "int", Children: []Node{}},
				{Name: "identifier", Value: "Ax", Children: []Node{}},
				{Name: "symbol", Value: ",", Children: []Node{}},
				{Name: "keyword", Value: "int", Children: []Node{}},
				{Name: "identifier", Value: "Ay", Children: []Node{}},
				{Name: "symbol", Value: ",", Children: []Node{}},
				{Name: "keyword", Value: "int", Children: []Node{}},
				{Name: "identifier", Value: "Asize", Children: []Node{}},
			}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "x", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Ax", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "y", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Ay", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "size", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Asize", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "draw", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "keyword", Value: "this", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "dispose", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Memory", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "deAlloc", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "keyword", Value: "this", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "draw", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Screen", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "setColor", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "keyword", Value: "true", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Screen", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "y", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
								{Name: "symbol", Value: "+", Children: []Node{}},
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "size", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "y", Children: []Node{}},
								}},
								{Name: "symbol", Value: "+", Children: []Node{}},
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "size", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "erase", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Screen", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "setColor", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "keyword", Value: "false", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Screen", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "y", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
								{Name: "symbol", Value: "+", Children: []Node{}},
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "size", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "y", Children: []Node{}},
								}},
								{Name: "symbol", Value: "+", Children: []Node{}},
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "size", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "incSize", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "y", Children: []Node{}},
											}},
											{Name: "symbol", Value: "+", Children: []Node{}},
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "size", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ")", Children: []Node{}},
									}},
									{Name: "symbol", Value: "&lt;", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "254", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
							}},
							{Name: "symbol", Value: "&amp;", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "x", Children: []Node{}},
											}},
											{Name: "symbol", Value: "+", Children: []Node{}},
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "size", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ")", Children: []Node{}},
									}},
									{Name: "symbol", Value: "&lt;", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "510", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "erase", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "size", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "size", Children: []Node{}},
									}},
									{Name: "symbol", Value: "+", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "2", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "draw", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "decSize", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "size", Children: []Node{}},
							}},
							{Name: "symbol", Value: "&gt;", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "2", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "erase", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "size", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "size", Children: []Node{}},
									}},
									{Name: "symbol", Value: "-", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "2", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "draw", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveUp", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "y", Children: []Node{}},
							}},
							{Name: "symbol", Value: "&gt;", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "1", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "keyword", Value: "false", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "symbol", Value: "(", Children: []Node{}},
											{Name: "expression", Value: "", Children: []Node{
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "y", Children: []Node{}},
												}},
												{Name: "symbol", Value: "+", Children: []Node{}},
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "size", Children: []Node{}},
												}},
											}},
											{Name: "symbol", Value: ")", Children: []Node{}},
										}},
										{Name: "symbol", Value: "-", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "1", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "y", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "y", Children: []Node{}},
									}},
									{Name: "symbol", Value: "-", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "2", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "keyword", Value: "true", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "1", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveDown", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "y", Children: []Node{}},
									}},
									{Name: "symbol", Value: "+", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "size", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
							}},
							{Name: "symbol", Value: "&lt;", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "254", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "keyword", Value: "false", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "1", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "y", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "y", Children: []Node{}},
									}},
									{Name: "symbol", Value: "+", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "2", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "keyword", Value: "true", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "symbol", Value: "(", Children: []Node{}},
											{Name: "expression", Value: "", Children: []Node{
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "y", Children: []Node{}},
												}},
												{Name: "symbol", Value: "+", Children: []Node{}},
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "size", Children: []Node{}},
												}},
											}},
											{Name: "symbol", Value: ")", Children: []Node{}},
										}},
										{Name: "symbol", Value: "-", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "1", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveLeft", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "x", Children: []Node{}},
							}},
							{Name: "symbol", Value: "&gt;", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "1", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "keyword", Value: "false", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "symbol", Value: "(", Children: []Node{}},
											{Name: "expression", Value: "", Children: []Node{
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "x", Children: []Node{}},
												}},
												{Name: "symbol", Value: "+", Children: []Node{}},
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "size", Children: []Node{}},
												}},
											}},
											{Name: "symbol", Value: ")", Children: []Node{}},
										}},
										{Name: "symbol", Value: "-", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "1", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "x", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "x", Children: []Node{}},
									}},
									{Name: "symbol", Value: "-", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "2", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "keyword", Value: "true", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "1", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveRight", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "x", Children: []Node{}},
									}},
									{Name: "symbol", Value: "+", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "size", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
							}},
							{Name: "symbol", Value: "&lt;", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "510", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "keyword", Value: "false", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "1", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "x", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "x", Children: []Node{}},
									}},
									{Name: "symbol", Value: "+", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "2", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "keyword", Value: "true", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "symbol", Value: "(", Children: []Node{}},
											{Name: "expression", Value: "", Children: []Node{
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "x", Children: []Node{}},
												}},
												{Name: "symbol", Value: "+", Children: []Node{}},
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "size", Children: []Node{}},
												}},
											}},
											{Name: "symbol", Value: ")", Children: []Node{}},
										}},
										{Name: "symbol", Value: "-", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "1", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
										{Name: "symbol", Value: "+", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "size", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "symbol", Value: "}", Children: []Node{}},
	}}
}

func Square_SquareGame_Tree(t *testing.T) Node {
	return Node{Name: "class", Value: "", Children: []Node{
		{Name: "keyword", Value: "class", Children: []Node{}},
		{Name: "identifier", Value: "SquareGame", Children: []Node{}},
		{Name: "symbol", Value: "{", Children: []Node{}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "field", Children: []Node{}},
			{Name: "identifier", Value: "Square", Children: []Node{}},
			{Name: "identifier", Value: "square", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "field", Children: []Node{}},
			{Name: "keyword", Value: "int", Children: []Node{}},
			{Name: "identifier", Value: "direction", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "constructor", Children: []Node{}},
			{Name: "identifier", Value: "SquareGame", Children: []Node{}},
			{Name: "identifier", Value: "new", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "square", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "new", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "0", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "0", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "integerConstant", Value: "30", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "direction", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "0", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "keyword", Value: "this", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "dispose", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "square", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "dispose", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Memory", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "deAlloc", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "keyword", Value: "this", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveSquare", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
							{Name: "symbol", Value: "=", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "1", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "moveUp", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
							{Name: "symbol", Value: "=", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "2", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "moveDown", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
							{Name: "symbol", Value: "=", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "3", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "moveLeft", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
							{Name: "symbol", Value: "=", Children: []Node{}},
							{Name: "term", Value: "", Children: []Node{
								{Name: "integerConstant", Value: "4", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "moveRight", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Sys", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "wait", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "integerConstant", Value: "5", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "run", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "keyword", Value: "char", Children: []Node{}},
					{Name: "identifier", Value: "key", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "keyword", Value: "boolean", Children: []Node{}},
					{Name: "identifier", Value: "exit", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "exit", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "keyword", Value: "false", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "whileStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "while", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "symbol", Value: "~", Children: []Node{}},
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "exit", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "whileStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "while", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
									{Name: "symbol", Value: "=", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "0", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "key", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "Keyboard", Children: []Node{}},
												{Name: "symbol", Value: ".", Children: []Node{}},
												{Name: "identifier", Value: "keyPressed", Children: []Node{}},
												{Name: "symbol", Value: "(", Children: []Node{}},
												{Name: "expressionList", Value: "", Children: []Node{}},
												{Name: "symbol", Value: ")", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
									{Name: "doStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "do", Children: []Node{}},
										{Name: "identifier", Value: "moveSquare", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{}},
										{Name: "symbol", Value: ")", Children: []Node{}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
									{Name: "symbol", Value: "=", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "81", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "exit", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "keyword", Value: "true", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
									{Name: "symbol", Value: "=", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "90", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "doStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "do", Children: []Node{}},
										{Name: "identifier", Value: "square", Children: []Node{}},
										{Name: "symbol", Value: ".", Children: []Node{}},
										{Name: "identifier", Value: "decSize", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{}},
										{Name: "symbol", Value: ")", Children: []Node{}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
									{Name: "symbol", Value: "=", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "88", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "doStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "do", Children: []Node{}},
										{Name: "identifier", Value: "square", Children: []Node{}},
										{Name: "symbol", Value: ".", Children: []Node{}},
										{Name: "identifier", Value: "incSize", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{}},
										{Name: "symbol", Value: ")", Children: []Node{}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
									{Name: "symbol", Value: "=", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "131", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "direction", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "integerConstant", Value: "1", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
									{Name: "symbol", Value: "=", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "133", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "direction", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "integerConstant", Value: "2", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
									{Name: "symbol", Value: "=", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "130", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "direction", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "integerConstant", Value: "3", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
									{Name: "symbol", Value: "=", Children: []Node{}},
									{Name: "term", Value: "", Children: []Node{
										{Name: "integerConstant", Value: "132", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "direction", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "integerConstant", Value: "4", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "whileStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "while", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "symbol", Value: "~", Children: []Node{}},
										{Name: "term", Value: "", Children: []Node{
											{Name: "symbol", Value: "(", Children: []Node{}},
											{Name: "expression", Value: "", Children: []Node{
												{Name: "term", Value: "", Children: []Node{
													{Name: "identifier", Value: "key", Children: []Node{}},
												}},
												{Name: "symbol", Value: "=", Children: []Node{}},
												{Name: "term", Value: "", Children: []Node{
													{Name: "integerConstant", Value: "0", Children: []Node{}},
												}},
											}},
											{Name: "symbol", Value: ")", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "key", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "Keyboard", Children: []Node{}},
												{Name: "symbol", Value: ".", Children: []Node{}},
												{Name: "identifier", Value: "keyPressed", Children: []Node{}},
												{Name: "symbol", Value: "(", Children: []Node{}},
												{Name: "expressionList", Value: "", Children: []Node{}},
												{Name: "symbol", Value: ")", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
									{Name: "doStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "do", Children: []Node{}},
										{Name: "identifier", Value: "moveSquare", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{}},
										{Name: "symbol", Value: ")", Children: []Node{}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "symbol", Value: "}", Children: []Node{}},
	}}
}

func SquareExpressionless_Main_Tree(t *testing.T) Node {
	return Node{Name: "class", Value: "", Children: []Node{
		{Name: "keyword", Value: "class", Children: []Node{}},
		{Name: "identifier", Value: "Main", Children: []Node{}},
		{Name: "symbol", Value: "{", Children: []Node{}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "static", Children: []Node{}},
			{Name: "keyword", Value: "boolean", Children: []Node{}},
			{Name: "identifier", Value: "test", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "function", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "main", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "identifier", Value: "SquareGame", Children: []Node{}},
					{Name: "identifier", Value: "game", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "game", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "game", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "game", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "run", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "game", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "dispose", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "function", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "more", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "keyword", Value: "boolean", Children: []Node{}},
					{Name: "identifier", Value: "b", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "b", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{}},
						{Name: "symbol", Value: "}", Children: []Node{}},
						{Name: "keyword", Value: "else", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "symbol", Value: "}", Children: []Node{}},
	}}
}

func SquareExpressionless_Square_Tree(t *testing.T) Node {
	return Node{Name: "class", Value: "", Children: []Node{
		{Name: "keyword", Value: "class", Children: []Node{}},
		{Name: "identifier", Value: "Square", Children: []Node{}},
		{Name: "symbol", Value: "{", Children: []Node{}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "field", Children: []Node{}},
			{Name: "keyword", Value: "int", Children: []Node{}},
			{Name: "identifier", Value: "x", Children: []Node{}},
			{Name: "symbol", Value: ",", Children: []Node{}},
			{Name: "identifier", Value: "y", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "field", Children: []Node{}},
			{Name: "keyword", Value: "int", Children: []Node{}},
			{Name: "identifier", Value: "size", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "constructor", Children: []Node{}},
			{Name: "identifier", Value: "Square", Children: []Node{}},
			{Name: "identifier", Value: "new", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{
				{Name: "keyword", Value: "int", Children: []Node{}},
				{Name: "identifier", Value: "Ax", Children: []Node{}},
				{Name: "symbol", Value: ",", Children: []Node{}},
				{Name: "keyword", Value: "int", Children: []Node{}},
				{Name: "identifier", Value: "Ay", Children: []Node{}},
				{Name: "symbol", Value: ",", Children: []Node{}},
				{Name: "keyword", Value: "int", Children: []Node{}},
				{Name: "identifier", Value: "Asize", Children: []Node{}},
			}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "x", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Ax", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "y", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Ay", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "size", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "Asize", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "draw", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "x", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "dispose", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Memory", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "deAlloc", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "keyword", Value: "this", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "draw", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Screen", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "setColor", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Screen", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "y", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "y", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "erase", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Screen", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "setColor", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Screen", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "y", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "x", Children: []Node{}},
								}},
							}},
							{Name: "symbol", Value: ",", Children: []Node{}},
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "y", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "incSize", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "x", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "erase", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "size", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "size", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "draw", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "decSize", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "size", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "erase", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "size", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "size", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "draw", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveUp", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "y", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "y", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "y", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveDown", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "y", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "y", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "y", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveLeft", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "x", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "x", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "x", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveRight", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "x", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "letStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "let", Children: []Node{}},
								{Name: "identifier", Value: "x", Children: []Node{}},
								{Name: "symbol", Value: "=", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "x", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "setColor", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "Screen", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "drawRectangle", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "x", Children: []Node{}},
										}},
									}},
									{Name: "symbol", Value: ",", Children: []Node{}},
									{Name: "expression", Value: "", Children: []Node{
										{Name: "term", Value: "", Children: []Node{
											{Name: "identifier", Value: "y", Children: []Node{}},
										}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "symbol", Value: "}", Children: []Node{}},
	}}
}

func SquareExpressionless_SquareGame_Tree(t *testing.T) Node {
	return Node{Name: "class", Value: "", Children: []Node{
		{Name: "keyword", Value: "class", Children: []Node{}},
		{Name: "identifier", Value: "SquareGame", Children: []Node{}},
		{Name: "symbol", Value: "{", Children: []Node{}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "field", Children: []Node{}},
			{Name: "identifier", Value: "Square", Children: []Node{}},
			{Name: "identifier", Value: "square", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "classVarDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "field", Children: []Node{}},
			{Name: "keyword", Value: "int", Children: []Node{}},
			{Name: "identifier", Value: "direction", Children: []Node{}},
			{Name: "symbol", Value: ";", Children: []Node{}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "constructor", Children: []Node{}},
			{Name: "identifier", Value: "SquareGame", Children: []Node{}},
			{Name: "identifier", Value: "new", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "square", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "square", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "direction", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "square", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "dispose", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "square", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "dispose", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Memory", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "deAlloc", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "square", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "moveSquare", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "moveUp", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "moveDown", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "moveLeft", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "ifStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "if", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "direction", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "doStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "do", Children: []Node{}},
								{Name: "identifier", Value: "square", Children: []Node{}},
								{Name: "symbol", Value: ".", Children: []Node{}},
								{Name: "identifier", Value: "moveRight", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expressionList", Value: "", Children: []Node{}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: ";", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "doStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "do", Children: []Node{}},
						{Name: "identifier", Value: "Sys", Children: []Node{}},
						{Name: "symbol", Value: ".", Children: []Node{}},
						{Name: "identifier", Value: "wait", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expressionList", Value: "", Children: []Node{
							{Name: "expression", Value: "", Children: []Node{
								{Name: "term", Value: "", Children: []Node{
									{Name: "identifier", Value: "direction", Children: []Node{}},
								}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "subroutineDec", Value: "", Children: []Node{
			{Name: "keyword", Value: "method", Children: []Node{}},
			{Name: "keyword", Value: "void", Children: []Node{}},
			{Name: "identifier", Value: "run", Children: []Node{}},
			{Name: "symbol", Value: "(", Children: []Node{}},
			{Name: "parameterList", Value: "", Children: []Node{}},
			{Name: "symbol", Value: ")", Children: []Node{}},
			{Name: "subroutineBody", Value: "", Children: []Node{
				{Name: "symbol", Value: "{", Children: []Node{}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "keyword", Value: "char", Children: []Node{}},
					{Name: "identifier", Value: "key", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "varDec", Value: "", Children: []Node{
					{Name: "keyword", Value: "var", Children: []Node{}},
					{Name: "keyword", Value: "boolean", Children: []Node{}},
					{Name: "identifier", Value: "exit", Children: []Node{}},
					{Name: "symbol", Value: ";", Children: []Node{}},
				}},
				{Name: "statements", Value: "", Children: []Node{
					{Name: "letStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "let", Children: []Node{}},
						{Name: "identifier", Value: "exit", Children: []Node{}},
						{Name: "symbol", Value: "=", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "key", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
					{Name: "whileStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "while", Children: []Node{}},
						{Name: "symbol", Value: "(", Children: []Node{}},
						{Name: "expression", Value: "", Children: []Node{
							{Name: "term", Value: "", Children: []Node{
								{Name: "identifier", Value: "exit", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: ")", Children: []Node{}},
						{Name: "symbol", Value: "{", Children: []Node{}},
						{Name: "statements", Value: "", Children: []Node{
							{Name: "whileStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "while", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "key", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "key", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
									{Name: "doStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "do", Children: []Node{}},
										{Name: "identifier", Value: "moveSquare", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{}},
										{Name: "symbol", Value: ")", Children: []Node{}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "exit", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "exit", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "doStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "do", Children: []Node{}},
										{Name: "identifier", Value: "square", Children: []Node{}},
										{Name: "symbol", Value: ".", Children: []Node{}},
										{Name: "identifier", Value: "decSize", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{}},
										{Name: "symbol", Value: ")", Children: []Node{}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "doStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "do", Children: []Node{}},
										{Name: "identifier", Value: "square", Children: []Node{}},
										{Name: "symbol", Value: ".", Children: []Node{}},
										{Name: "identifier", Value: "incSize", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{}},
										{Name: "symbol", Value: ")", Children: []Node{}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "direction", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "exit", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "direction", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "key", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "direction", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "square", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "ifStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "if", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "direction", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "direction", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
							{Name: "whileStatement", Value: "", Children: []Node{
								{Name: "keyword", Value: "while", Children: []Node{}},
								{Name: "symbol", Value: "(", Children: []Node{}},
								{Name: "expression", Value: "", Children: []Node{
									{Name: "term", Value: "", Children: []Node{
										{Name: "identifier", Value: "key", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: ")", Children: []Node{}},
								{Name: "symbol", Value: "{", Children: []Node{}},
								{Name: "statements", Value: "", Children: []Node{
									{Name: "letStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "let", Children: []Node{}},
										{Name: "identifier", Value: "key", Children: []Node{}},
										{Name: "symbol", Value: "=", Children: []Node{}},
										{Name: "expression", Value: "", Children: []Node{
											{Name: "term", Value: "", Children: []Node{
												{Name: "identifier", Value: "key", Children: []Node{}},
											}},
										}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
									{Name: "doStatement", Value: "", Children: []Node{
										{Name: "keyword", Value: "do", Children: []Node{}},
										{Name: "identifier", Value: "moveSquare", Children: []Node{}},
										{Name: "symbol", Value: "(", Children: []Node{}},
										{Name: "expressionList", Value: "", Children: []Node{}},
										{Name: "symbol", Value: ")", Children: []Node{}},
										{Name: "symbol", Value: ";", Children: []Node{}},
									}},
								}},
								{Name: "symbol", Value: "}", Children: []Node{}},
							}},
						}},
						{Name: "symbol", Value: "}", Children: []Node{}},
					}},
					{Name: "returnStatement", Value: "", Children: []Node{
						{Name: "keyword", Value: "return", Children: []Node{}},
						{Name: "symbol", Value: ";", Children: []Node{}},
					}},
				}},
				{Name: "symbol", Value: "}", Children: []Node{}},
			}},
		}},
		{Name: "symbol", Value: "}", Children: []Node{}},
	}}
}
