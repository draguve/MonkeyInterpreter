package compiler

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/code"
	"MonkeyInterpreter/lexer"
	"MonkeyInterpreter/object"
	"MonkeyInterpreter/parser"
	"fmt"
	"testing"
)

type compilerTestCase struct {
	input string
	expectedConstants []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T){
	tests := []compilerTestCase{
		{
			input: "1 + 2",
			expectedConstants: []interface{}{1,2},
			expectedInstructions:[]code.Instructions{
				code.Make(code.LDA,0),
				code.Make(code.LDA,1),
			},
		},
	}
	runCompilerTests(t,tests)
}

func parse(input string) *ast.Program{
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func runCompilerTests(t *testing.T, tests []compilerTestCase){
	t.Helper()

	for _,tt := range tests{
		program := parse(tt.input)

		compiler := New()
		err := compiler.Compile(program)
		if err != nil{
			t.Fatalf("compiler error: %s",err)
		}

		byteCode := compiler.Bytecode()

		err = testInstructions(tt.expectedInstructions,byteCode.Instructions)
		if err != nil {
			t.Fatalf("testInstructions failed: %s",err)
		}

		err = testConstants(t,tt.expectedConstants,byteCode.Constants)
		if err != nil{
			t.Fatalf("testConstants failed: %s",err)
		}
	}
}

func testInstructions(
	expected []code.Instructions,
	actual code.Instructions,
) error {
	concatted := concatInstructions(expected)
	if len(actual) != len(concatted) {
		return fmt.Errorf("wrong instructions length.\nwant=%q\ngot =%q",
			concatted, actual)
	}
	for i, ins := range concatted {
		if actual[i] != ins {
			return fmt.Errorf("wrong instruction at %d.\nwant=%q\ngot =%q",
				i, concatted, actual)
		}
	}
	return nil
}

func concatInstructions(s []code.Instructions) code.Instructions {
	out := code.Instructions{}
	for _, ins := range s {
		out = append(out, ins...)
	}
	return out
}

func testConstants(
	t *testing.T,
	expected []interface{},
	actual []object.Object,
) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("wrong number of constants. got=%d, want=%d",
			len(actual), len(expected))
	}
	for i, constant := range expected {
		switch constant := constant.(type) {
		case uint16:
			err := testIntegerObject(uint16(constant), actual[i])
			if err != nil {
				return fmt.Errorf("constant %d - testIntegerObject failed: %s",
					i, err)
			}
		}
	}
	return nil
}

func testIntegerObject(expected uint16, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)",
			actual, actual)
	}
	if uint16(result.Value) != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d",result.Value, expected)
	}
	return nil
}