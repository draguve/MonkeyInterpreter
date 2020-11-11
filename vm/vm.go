package vm

import (
	"MonkeyInterpreter/code"
	"MonkeyInterpreter/compiler"
	"MonkeyInterpreter/object"
	"errors"
	"fmt"
)

const StackSize = 2048

type VM struct{
	constants []object.Object
	instructions code.Instructions

	stack []object.Object
	sp uint16
	registerA uint16
	registerB uint16
	registerX uint16
}

func New(bytecode *compiler.ByteCode) *VM{
	return &VM{
		instructions: bytecode.Instructions,
		constants: bytecode.Constants,

		stack: make([] object.Object,StackSize),
	}
}

func (vm *VM) StackTop() object.Object{
	if vm.sp == 0{
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error{
	for ip:=0;ip<len(vm.instructions);ip++{
		op := code.Opcode(vm.instructions[ip])

		switch op{
		case code.LDA:
			constIndex := vm.instructions[ip+1]
			ip += 1
			v, ok := vm.constants[constIndex].(*object.Integer)
			if !ok{
				return errors.New("The constant to be loaded was not a valid constant address")
			}
			vm.registerA = uint16(v.Value)
		}
	}
	return nil
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize{
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}
