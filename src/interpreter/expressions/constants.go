package expressions

import (
	. "interpreter"
	. "interpreter/values"
)

// BoolConst

type BoolConst struct {
	v bool
}

func (e *BoolConst) Execute(args []Value) Value {
	return NewConcreteBoolValue(e.v)
}

func NewBoolConst(v bool) *BoolConst {
	return &BoolConst{v}
}

// IntConst

type IntConst struct {
	v int
}

func (e *IntConst) Execute(args []Value) Value {
	return NewConcreteIntValue(e.v)
}

func NewIntConst(v int) *IntConst {
	return &IntConst{v}
}

// StringConst

type StringConst struct {
	v string
}

func (e *StringConst) Execute(args []Value) Value {
	return NewConcreteStringValue(e.v)
}

func NewStringConst(v string) *StringConst {
	return &StringConst{v}
}
