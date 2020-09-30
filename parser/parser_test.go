package parser

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T){
	input := `
let x = 5;
let y = 10;
let foobar = 87878;
`
	lex := lexer.New(input)
	parser := New(lex)

	program := parser.ParseProgram()
	if program == nil{
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3{
		t.Fatalf("program.Statements does not contain 3 statements got=%d",len(program.Statements))
	}
	tests := []struct{
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i,tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t,stmt,tt.expectedIdentifier){
			return
		}
	}

}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let"{
		t.Errorf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt,ok := stmt.(*ast.LetStatement)
	if !ok{
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != name{
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name{
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",name,letStmt.Name.TokenLiteral())
		return false
	}
	return true
}
