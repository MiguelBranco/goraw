package expressions

import (
	. "interpreter"
	. "interpreter/values"
)

type AttributeConstruction struct {
	N string
	E Expression
}

type RecordConstruction struct {
	atts []AttributeConstruction
}

func (e *RecordConstruction) Execute(args []Value) Value {
	atts := make([]ConcreteRecordAttribute, len(e.atts))
	for i, att := range e.atts {
		atts[i].N = att.N
		atts[i].V = att.E.Execute(args)
	}
	return NewConcreteRecordValue(atts)
}

func (e *RecordConstruction) Type() ExpressionType {
	return Record
}

func NewRecordConstruction(atts []AttributeConstruction) *RecordConstruction {
	return &RecordConstruction{atts}
}