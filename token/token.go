package token

type TType string

var keywords = map[string]TType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

type Token struct {
	Type    TType  //this will hold what type the token it is
	Literal string //this will hold what was in the source code
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // names of variables or functions
	INT   = "INT"   // 4,48 etc

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	COMMA     = ","
	SEMICOLON = ";"

	LBRACK = "("
	RBRACK = ")"
	LBRACE = "{"
	RBRACE = "}"
	LSQRBRACK = "["
	RSQRBRACK = "]"


	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

//TokenType Returns the type of Token when ident is supplied with a string
func WhichTokenType(ident string) TType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
