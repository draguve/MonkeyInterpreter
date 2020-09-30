package parser

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/lexer"
	"fmt"
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

func TestIdentifierExpression(t *testing.T){
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t,p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",len(program.Statements))
	}
	stmt,ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not converted to identifier. got=%T",stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T){
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t,p)

	if len(program.Statements) != 1{
		t.Fatalf("program has not enough statement. got=%d",len(program.Statements))
	}
	stmt,ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok{
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal,ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok{
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T){
	prefixTests := []struct{
		input string
		operator string
		intValue int16
	}{
		{"!5;","!",5},
		{"-15","-",15},
	}
	for _,tt := range prefixTests{
		l := lexer.New(tt.input);
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t,p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt,ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok{
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator{
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t,exp.Right,tt.intValue){
			return
		}

	}
}

func testIntegerLiteral(t *testing.T, intValue ast.Expression, value int16) bool {
	integer,ok := intValue.(*ast.IntegerLiteral)
	if !ok{
		t.Errorf("il not *ast.IntegerLiteral. got=%T", intValue)
		return false
	}
	if integer.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integer.Value)
		return false
	}
	if integer.TokenLiteral() != fmt.Sprintf("%d",value){
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integer.TokenLiteral())
		return false
	}
	return true
}

func TestParsingInfixExpressions(t *testing.T){
	infixTests := []struct {
		input      string
		leftValue  int16
		operator   string
		rightValue int16
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _,tt := range infixTests{
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		fmt.Printf("%+v",program.Statements)
		checkParserErrors(t,p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}