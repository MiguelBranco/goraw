package values

import . "interpreter"

// IntValue

type IntValue interface {
	Get() int
}

// ConcreteIntValue

type ConcreteIntValue struct {
	v int
}

func (this *ConcreteIntValue) Equal(that Value) bool {
	switch that := that.(type) {
	case IntValue:
		return this.v == that.Get()
	default:
		return false
	}
}

func (i *ConcreteIntValue) Get() int {
	return i.v
}

func NewConcreteIntValue(v int) *ConcreteIntValue {
	return &ConcreteIntValue{v}
}