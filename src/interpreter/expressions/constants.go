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

func (e *BoolConst) Type() ExpressionType {
	return Bool
}

func NewBoolConst(v bool) *BoolConst {
	return &BoolConst{v}
}

// FloatConst

type FloatConst struct {
	v float64
}

func (e *FloatConst) Execute(args []Value) Value {
	return NewConcreteFloatValue(e.v)
}

func (e *FloatConst) Type() ExpressionType {
	return Float
}

func NewFloatConst(v float64) *FloatConst {
	return &FloatConst{v}
}

// IntConst

type IntConst struct {
	v int
}

func (e *IntConst) Execute(args []Value) Value {
	return NewConcreteIntValue(e.v)
}

func (e *IntConst) Type() ExpressionType {
	return Int
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

func (e *StringConst) Type() ExpressionType {
	return String
}

func NewStringConst(v string) *StringConst {
	return &StringConst{v}
}
