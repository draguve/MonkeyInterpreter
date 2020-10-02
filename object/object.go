package object

import (
	"MonkeyInterpreter/ast"
	"bytes"
	"fmt"
	"strings"
)

type Type string

type Object interface {
	Type() Type
	Inspect() string
}

const (
	INTEGER = "INTEGER"
	BOOLEAN = "BOOLEAN"
	NULL    = "NULL"
	RETURN  = "RETURN"
	ERROR = "ERROR"
	FUNCTION = "FUNCTION"
	BUILTIN_OBJ = "BUILTIN"
	ARRAY = "ARRAY"
)

type Integer struct {
	Value int16
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() Type {
	return INTEGER
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
func (b *Boolean) Type() Type {
	return BOOLEAN
}

type Null struct{}

func (n *Null) Type() Type {
	return NULL
}
func (n *Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Type() Type {
	return RETURN
}
func (r *ReturnValue) Inspect() string {
	return r.Value.Inspect()
}

type Error struct {
	Message string
}
func (e *Error) Type() Type { return ERROR }
func (e *Error) Inspect() string { return "ERROR: " + e.Message}

type Function struct {
	Arguments []*ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}

func (f *Function) Type() Type { return FUNCTION }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	var params []string
	for _, p := range f.Arguments {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string {return "builtin function"}

type Array struct {
	Elements []Object
}

func (a *Array) Type() Type {return ARRAY}
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elems := []string{}
	for _,e := range a.Elements {
		elems = append(elems,e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elems,","))
	out.WriteString("]")

	return out.String()
}

