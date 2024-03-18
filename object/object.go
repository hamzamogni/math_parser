package object

import (
	"fmt"
)

type ObjectType string

const (
	ERROR_OBJ   = "ERROR"
	NUMBER_OBJ  = "NUMBER"
	BUILTIN_OBJ = "BUILTIN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Number struct {
	Value float64
}

func (n *Number) Type() ObjectType { return NUMBER_OBJ }
func (n *Number) Inspect() string  { return fmt.Sprintf("%.2f", n.Value) }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string {
	return "Builtin Function"
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
