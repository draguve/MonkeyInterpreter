package repl

import(
	"MonkeyInterpreter/lexer"
	"MonkeyInterpreter/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader,out io.Writer){
	scanner := bufio.NewScanner(in)

	for{
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned{
			return
		}
		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken();tok.Type != token.EOF;tok = lex.NextToken(){
			_, _ = fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}