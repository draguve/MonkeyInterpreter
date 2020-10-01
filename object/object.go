package object

import "fmt"

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
