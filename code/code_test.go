package code

import (
	"testing"
)


func TestMake(t *testing.T) {
	tests := []struct {
		op Opcode
		operand uint16
		expected []uint16
	}{
		{LDA, 0x23, []uint16{0x00,0x23}},
	}
	for _, tt := range tests {
		instruction := Make(tt.op, tt.operand)

		if len(instruction) != len(tt.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d",
				len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d",
					i, b, instruction[i])
			}
		}
	}
}

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{
		Make(LDA, 1),
		Make(LDA, 2),
		Make(LDA, 65535),
	}
	var expected = `0000 LDA 1
0002 LDA 2
0004 LDA 65535
`
	concatted := Instructions{}
	for _, ins := range instructions {
		concatted = append(concatted, ins...)
	}
	if concatted.String() != expected {
		t.Errorf("instructions wrongly formatted.\nwant=%q\ngot=%q",
			expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op Opcode
		operand uint16
		operandRequired bool
	}{
		{LDA, 0, true},
	}
	for _, tt := range tests {
		instruction := Make(tt.op, tt.operand)
		def, err := Lookup(byte(tt.op))
		if err != nil {
			t.Fatalf("definition not found: %q\n", err)
		}
		operand,operandRequired := ReadOperand(def, instruction[1:])
		if operandRequired != tt.operandRequired{
			t.Fatalf("intruction format wrong ")
		}
		if operandRequired {
			if operand != tt.operand {
				t.Errorf("operand wrong. want=%d, got=%d", tt.operand, operand)
			}
		}
	}
}