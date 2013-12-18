package expressions

import (
	. "interpreter"
	. "interpreter/values"
)

type Not struct {
	e Expression
}

func (e *Not) Execute(args []Value) Value {
	switch v := e.e.Execute(args).(type) {
	case BoolValue:
		return NewConcreteBoolValue(!v.Get())
	default:
		panic("invalid type")
	}
}

func NewNot(e Expression) *Not {
	return &Not{e}
}
