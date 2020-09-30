package ast

import (
	"MonkeyInterpreter/token"
	"bytes"
)

type Node interface {
	 TokenLiteral() string
	 String() string
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
func (p *Program) String() string{
	var out bytes.Buffer
	for _,s := range p.Statements{
		out.WriteString(s.String())
	}
	return out.String()
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
func (l *LetStatement) String() string{
	var out bytes.Buffer
	out.WriteString(l.TokenLiteral() + " " + l.Name.String() + " = " + l.Value.String()+";")
	return out.String()
}


type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string // this will store the name of the variable/function
}
func (i *Identifier) expressionNode() {

}
func (i *Identifier) TokenLiteral() string{
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}
func (r *ReturnStatement) statementNode(){}
func (r *ReturnStatement) TokenLiteral() string{
	return r.Token.Literal
}
func (r *ReturnStatement) String() string{
	var out bytes.Buffer
	out.WriteString(r.TokenLiteral() + " " + r.Value.String()+";")
	return out.String()
}

type ExpressionStatement struct{
	Token token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode(){}
func (es *ExpressionStatement) TokenLiteral() string{
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string{
	if es.Expression != nil{
		return es.Expression.String()
	}
	return ""
}
