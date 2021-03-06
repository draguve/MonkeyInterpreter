package lexer

import (
	"MonkeyInterpreter/token"
	"testing"
)

//func TestNextToken(t *testing.T) {
//	input := `=+(){},;`
//	tests := []struct {
//		expectedType token.TType
//		expectedLiteral string
//	}{
//		{token.ASSIGN, "="},
//		{token.PLUS, "+"},
//		{token.LBRACK, "("},
//		{token.RBRACK, ")"},
//		{token.LBRACE, "{"},
//		{token.RBRACE, "}"},
//		{token.COMMA, ","},
//		{token.SEMICOLON, ";"},
//		{token.EOF, ""},
//	}
//	l := New(input)
//	for i, tt := range tests {
//		tok := l.NextToken()
//		if tok.Type != tt.expectedType {
//			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
//				i, tt.expectedType, tok.Type)
//		}
//		if tok.Literal != tt.expectedLiteral {
//			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
//				i, tt.expectedLiteral, tok.Literal)
//		}
//	}
//}

func TestNextToken(t *testing.T) {
	input := `let five int = 5;
let ten bool = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
!-5;
5 < 10 > 5;

if(5 < 10){
	return true;
} else {
	return false;
}
10 == 10;
10 != 9;
[1, 2];
for`
	tests := []struct {
		expectedType    token.TType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.INT_TYPE,"int"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.BOOL_TYPE,"bool"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LBRACK, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RBRACK, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LBRACK, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RBRACK, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LBRACK, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RBRACK, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.LSQRBRACK, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RSQRBRACK, "]"},
		{token.SEMICOLON, ";"},
		{token.FOR,"for"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
