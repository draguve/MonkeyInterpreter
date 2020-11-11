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
	switch node := node.(type) {
		case *ast.Program:
			for _,s := range node.Statements{
				err := c.Compile(s)
				if err != nil{
					return err
				}
			}
		case *ast.ExpressionStatement:
			err := c.Compile(node.Expression)
			if err != nil {
				return err
			}
		case *ast.InfixExpression:
			err := c.Compile(node.Left)
			if err != nil {
				return err
			}

			err = c.Compile(node.Right)
			if err != nil{
				return err
			}
		case *ast.IntegerLiteral:
			integer := &object.Integer{Value: node.Value}
			//TODO : be careful have to make changes to reflect the fact that we have the ALU between the a & b registers instead of a stack machine
			//TODO : check if there's an overflow on the location for the constant later on,hence we cant handle more that 2^16 constants
			c.emit(code.LDA,uint16(c.addConstant(integer)))
	}
	return nil
}

func (c *Compiler) emit(op code.Opcode,operand uint16) int{
	ins := code.Make(op,operand)
	pos := c.addInstructions(ins)
	return pos
}

func (c *Compiler) addConstant(obj object.Object) int{
	c.constants = append(c.constants,obj)
	return len(c.constants) - 1
}

func (c *Compiler) addInstructions(ins []uint16) int{
	posNewInstructions := len(c.instructions)
	c.instructions = append(c.instructions,ins...)
	return posNewInstructions
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