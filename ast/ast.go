package ast

import (
	"MonkeyInterpreter/token"
)

type Node interface {
	 TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}
func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token //the LET Token
	Name *Identifier //token for the name of the variable/function
	Value Expression // the other side of the let
}
func (l *LetStatement) TokenLiteral() string{
	return l.Token.Literal
}
func (l *LetStatement) statementNode(){}


type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string // this will store the name of the variable/function
}
func (i *Identifier) expressionNode() {

}
func (i *Identifier) TokenLiteral() string{
	return i.Token.Literal
}

