package expressions

import (
	. "interpreter"
	. "interpreter/values"
)

type IfThenElse struct {
	e1 Expression
	e2 Expression
	e3 Expression
}

func (e *IfThenElse) Execute(args []Value) Value {
	switch e1 := e.e1.Execute(args).(type) {
	case BoolValue:
		if e1.Get() {
			return e.e2.Execute(args)
		} else {
			return e.e3.Execute(args)
		}
	default:
		panic("invalid type")
	}
}

func NewIfThenElse(e1 Expression, e2 Expression, e3 Expression) *IfThenElse {
	return &IfThenElse{e1, e2, e3}
}
