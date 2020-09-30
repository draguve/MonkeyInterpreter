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
	checkParserErrors(t, parser)

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

func checkParserErrors(t *testing.T, parser *Parser) {
	errors := parser.Errors()
	if len(errors) == 0{
		return
	}
	t.Errorf("parser has %d errors",len(errors))
	for _,msg := range errors{
		t.Errorf("parser error: %q",msg)
	}
	t.FailNow()
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

func TestReturnStatements(t *testing.T){
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements got %d",len(program.Statements))
	}

	for _,stmt := range program.Statements{
		returned,ok := stmt.(*ast.ReturnStatement)
		if !ok{
			t.Errorf("return.TokenLiteral not return,got %q",returned.TokenLiteral())
			continue
		}
		if returned.TokenLiteral() != "return"{
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",returned.TokenLiteral())
		}
	}
}