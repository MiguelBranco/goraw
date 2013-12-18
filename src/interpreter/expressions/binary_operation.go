package expressions

import (
	. "interpreter"
	. "interpreter/values"
)

type BinaryOperator int32

const (
	Eq BinaryOperator = iota + 1
	Neq
	Ge
	Gt
	Le
	Lt
	Add
	Sub
	Mult
	Div
)

type BinaryOperation struct {
	t ExpressionType
	op BinaryOperator
	e1 Expression
	e2 Expression
}

func (e *BinaryOperation) Execute(args []Value) Value {
	switch e.op {
	case Eq:
		return NewConcreteBoolValue(e.e1.Execute(args).Equal(e.e2.Execute(args)))
	case Neq:
		return NewConcreteBoolValue(!e.e1.Execute(args).Equal(e.e2.Execute(args)))
	case Ge, Gt, Le, Lt, Add, Sub, Mult, Div:
		switch e.t {
		case Float:
			e1 := e.e1.Execute(args).(FloatValue)
			e2 := e.e2.Execute(args).(FloatValue)
			switch e.op {
			case Ge: return NewConcreteBoolValue(e1.Get() >= e2.Get())
			case Gt: return NewConcreteBoolValue(e1.Get() > e2.Get())
			case Le: return NewConcreteBoolValue(e1.Get() <= e2.Get())
			case Lt: return NewConcreteBoolValue(e1.Get() < e2.Get())
			case Add: return NewConcreteFloatValue(e1.Get() + e2.Get())
			case Sub: return NewConcreteFloatValue(e1.Get() - e2.Get())
			case Mult: return NewConcreteFloatValue(e1.Get() * e2.Get())
			case Div: return NewConcreteFloatValue(e1.Get() / e2.Get())
			}			
		case Int:
			e1 := e.e1.Execute(args).(IntValue)
			e2 := e.e2.Execute(args).(IntValue)
			switch e.op {
			case Ge: return NewConcreteBoolValue(e1.Get() >= e2.Get())
			case Gt: return NewConcreteBoolValue(e1.Get() > e2.Get())
			case Le: return NewConcreteBoolValue(e1.Get() <= e2.Get())
			case Lt: return NewConcreteBoolValue(e1.Get() < e2.Get())
			case Add: return NewConcreteIntValue(e1.Get() + e2.Get())
			case Sub: return NewConcreteIntValue(e1.Get() - e2.Get())
			case Mult: return NewConcreteIntValue(e1.Get() * e2.Get())
			case Div: return NewConcreteIntValue(e1.Get() / e2.Get())
			}
		}
	}
	panic("invalid binary operation")
}

func (e *BinaryOperation) Type() ExpressionType {
	return e.t
}

func NewBinaryOperation(t ExpressionType, op BinaryOperator, e1 Expression, e2 Expression) *BinaryOperation {
	return &BinaryOperation{t, op, e1, e2}
}

// FIXME: It is rather silly to have a "BinaryOperation" which takes an ExpressionType as input,
//        when for some operations (e.g. Eq, Neq, Ge, ...) the return type is known to be Bool.
//		  Split this BinaryOperation construct into separate ones, e.g. Eq, Neq, Ge, Gt, ...