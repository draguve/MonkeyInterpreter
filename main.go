package main

import (
	"MonkeyInterpreter/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
