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
	v := e.e1.Execute(args).(BoolValue)
	if v.Get() {
		return e.e2.Execute(args)
	} else {
		return e.e3.Execute(args)
	}
}

func (e *IfThenElse) Type() ExpressionType {
	return e.e2.Type()
}

func NewIfThenElse(e1 Expression, e2 Expression, e3 Expression) *IfThenElse {
	return &IfThenElse{e1, e2, e3}
}
