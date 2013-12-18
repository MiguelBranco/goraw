package expressions

import (
	. "interpreter"
	. "interpreter/values"
)

type Not struct {
	e Expression
}

func (e *Not) Execute(args []Value) Value {
	v := e.e.Execute(args).(BoolValue)
	return NewConcreteBoolValue(!v.Get())
}

func (e *Not) Type() ExpressionType {
	return Bool
}

func NewNot(e Expression) *Not {
	return &Not{e}
}
