package compiler

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/code"
	"MonkeyInterpreter/object"
)

type Compiler struct {
	instructions code.Instructions
	constants []object.Object
}

func New() *Compiler{
	return &Compiler{
		instructions: code.Instructions{},
		constants: []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error{
	return nil
}

type ByteCode struct {
	Instructions code.Instructions
	Constants []object.Object
}

func (c *Compiler) Bytecode() *ByteCode{
	return &ByteCode{
		Instructions: c.instructions,
		Constants: c.constants,
	}
}