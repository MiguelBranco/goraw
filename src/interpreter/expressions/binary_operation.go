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
		switch e1 := e.e1.Execute(args).(type) {
		case IntValue:
			switch e2 := e.e2.Execute(args).(type) {
			case IntValue:
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
			case FloatValue:
				switch e.op {
				case Ge: return NewConcreteBoolValue(float64(e1.Get()) >= e2.Get())
				case Gt: return NewConcreteBoolValue(float64(e1.Get()) > e2.Get())
				case Le: return NewConcreteBoolValue(float64(e1.Get()) <= e2.Get())
				case Lt: return NewConcreteBoolValue(float64(e1.Get()) < e2.Get())
				case Add: return NewConcreteFloatValue(float64(e1.Get()) + e2.Get())
				case Sub: return NewConcreteFloatValue(float64(e1.Get()) - e2.Get())
				case Mult: return NewConcreteFloatValue(float64(e1.Get()) * e2.Get())
				case Div: return NewConcreteFloatValue(float64(e1.Get()) / e2.Get())
				}
			default:
				panic("invalid type")
			}
		case FloatValue:
			switch e2 := e.e2.Execute(args).(type) {
			case IntValue:
				switch e.op {
				case Ge: return NewConcreteBoolValue(e1.Get() >= float64(e2.Get()))
				case Gt: return NewConcreteBoolValue(e1.Get() > float64(e2.Get()))
				case Le: return NewConcreteBoolValue(e1.Get() <= float64(e2.Get()))
				case Lt: return NewConcreteBoolValue(e1.Get() < float64(e2.Get()))
				case Add: return NewConcreteFloatValue(e1.Get() + float64(e2.Get()))
				case Sub: return NewConcreteFloatValue(e1.Get() - float64(e2.Get()))
				case Mult: return NewConcreteFloatValue(e1.Get() * float64(e2.Get()))
				case Div: return NewConcreteFloatValue(e1.Get() / float64(e2.Get()))
				}
			case FloatValue:
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
			default:
				panic("invalid type")
			}
		}
	}
	panic("invalid binary operator")
}

func NewBinaryOperation(op BinaryOperator, e1 Expression, e2 Expression) *BinaryOperation {
	return &BinaryOperation{op, e1, e2}
}
