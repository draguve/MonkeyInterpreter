package parser

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/lexer"
	"MonkeyInterpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(lex *lexer.Lexer) *Parser{
	parser := &Parser{l:lex}
	parser.nextToken()
	parser.nextToken()
	return parser
}

func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

//Implement this later
func (p *Parser) ParseProgram() *ast.Program{
	return nil
}
