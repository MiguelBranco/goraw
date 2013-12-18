package expressions

import (
	. "interpreter"
	. "interpreter/values"
)

type RecordProjection struct {
	t ExpressionType
	e Expression
	n string
}

func (e *RecordProjection) Execute(args []Value) Value {
	v := e.e.Execute(args).(RecordValue)
	return v.GetValueByName(e.n)
}

func (e *RecordProjection) Type() ExpressionType {
	return e.t
}

func NewRecordProjection(t ExpressionType, e Expression, n string) *RecordProjection {
	return &RecordProjection{t, e, n}
}