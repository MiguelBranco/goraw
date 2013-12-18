package expressions

import (
	. "interpreter"
	. "interpreter/values"
)

type RecordProjection struct {
	e Expression
	n string
}

func (e *RecordProjection) Execute(args []Value) Value {
	switch v := e.e.Execute(args).(type) {
	case RecordValue:
		return v.GetValueByName(e.n)
	default:
		panic("invalid type")
	}
}

func NewRecordProjection(e Expression, n string) *RecordProjection {
	return &RecordProjection{e, n}
}