package parser

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/lexer"
	"MonkeyInterpreter/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token

	errors []string
}

func New(lex *lexer.Lexer) *Parser{
	parser := &Parser{
		l:lex,
		errors: []string{},
	}
	parser.nextToken()
	parser.nextToken()
	return parser
}

func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF{
		stmt := p.parseStatement()
		if stmt != nil{
			program.Statements = append(program.Statements,stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) Errors() []string{
	return p.errors
}

func (p *Parser) peekError(t token.TType){
	msg := fmt.Sprintf("expected next token to be %s , got %s instead",t,p.peekToken.Type)
	p.errors = append(p.errors,msg)
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parserReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT){
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken,Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN){
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(t token.TType) bool{
	if p.peekTokenIs(t){
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) curTokenIs(t token.TType) bool{
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TType) bool{
	return p.peekToken.Type == t
}

func (p *Parser) parserReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	}
	return stmt
}

