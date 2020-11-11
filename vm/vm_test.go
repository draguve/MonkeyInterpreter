package vm

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/compiler"
	"MonkeyInterpreter/lexer"
	"MonkeyInterpreter/object"
	"MonkeyInterpreter/parser"
	"fmt"
	"testing"
)

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testIntegerObject(expected int16, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)",
			actual, actual)
	}
	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
	}
	return nil
}

type vmTestCase struct {
	input string
	expected interface{}
}

func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()
	for _, tt := range tests {
		program := parse(tt.input)
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}
		vm := New(comp.Bytecode())
		err = vm.Run()
		if err != nil {
			t.Fatalf("vm error: %s", err)
		}
		//stackElem := vm.StackTop()
		//testExpectedObject(t, tt.expected, stackElem)

		registerA := vm.registerA
		testRegisterAValue(t, tt.expected.(uint16),registerA)

	}
}

func testRegisterAValue(
	t *testing.T,
	expected uint16,
	actual uint16,
) {
	t.Helper()
	if expected != actual {
		t.Errorf("registerA value failed expected = %x , got = %x",expected,actual)
	}
}

func testExpectedObject(
	t *testing.T,
	expected interface{},
	actual object.Object,
) {
	t.Helper()
	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(int16(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", uint16(1)},
		{"2", uint16(2)},
		{"1 + 2", uint16(2)}, // FIXME
	}
	runVmTests(t, tests)
}