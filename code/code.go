package code
//Akshat is the best
//follow Akshat @gognaakshat
import (
	"bytes"
	"encoding/binary"
	"fmt"
)


type Instructions []uint16

type Opcode byte

const (
	LDA Opcode = iota //this instruction will load the address into the MAR
	LDB
)

//const AddressInstructionMask byte = 0xc0

type Definition struct {
	Name string
	DoubleLength bool
}

var definitions = map[Opcode]*Definition{
	LDA: {"LDA",true},//take data from memory address X and load it into A register
	LDB: {"LDB",true},
}

func (ins Instructions) String() string{
	var out bytes.Buffer

	i := 0
	for i < len(ins){
		b := make([]byte, 2)

		binary.LittleEndian.PutUint16(b,ins[i])
		def,err := Lookup(b[0])

		if err != nil{
			fmt.Fprintf(&out,"ERROR: %s\n",err)
			continue
		}
		operand,operandRequired := ReadOperand(def,ins[i:i+2])
		_, _ = fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operand, operandRequired))
		if operandRequired {
			i += 2
		}else{
			i += 1
		}

	}

	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operand uint16,operandRequired bool) string {
	if operandRequired != def.DoubleLength {
		return fmt.Sprintf("ERROR: operand len %t does not match defined %t\n",
			operandRequired, def.DoubleLength)
	}
	if operandRequired{
		return fmt.Sprintf("%s %d", def.Name, operand)
	}else{
		return fmt.Sprintf("%s", def.Name)
	}
}

func Lookup(op byte)(*Definition,error){
	def,ok := definitions[Opcode(op)]
	if !ok{
		return nil,fmt.Errorf("opcode %d undefined", op)
	}
	return def,nil
}

func Make(op Opcode, operand uint16) []uint16 {
	def,ok := definitions[op]
	if !ok {
		return []uint16{}
	}
	var instruction []uint16
	if def.DoubleLength{
		instruction = make([]uint16, 2)
		instruction[0] = binary.BigEndian.Uint16([]byte{byte(op),0x0})
		instruction[1] = operand
	}else{
		instruction = make([]uint16, 1)
		instruction[0] = binary.BigEndian.Uint16([]byte{byte(op),byte(operand)})
	}
	return instruction
}

func ReadOperand(def *Definition, ins Instructions) (uint16,bool) {
	if !def.DoubleLength{
		//implement the bit stuff here later
		return 0,false
	}
	if len(ins)!=2 {
		return 0,true
	}
	return ins[1],true
}