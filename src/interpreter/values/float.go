package values

import . "interpreter"

// FloatValue

type FloatValue interface {
	Get() float64
}

// ConcreteFloatValue

type ConcreteFloatValue struct {
	v float64
}

func (this *ConcreteFloatValue) Equal(that Value) bool {
	switch that := that.(type) {
	case FloatValue:
		return this.v == that.Get()
	default:
		return false
	}
}

func (f *ConcreteFloatValue) Get() float64 {
	return f.v
}

func NewConcreteFloatValue(v float64) *ConcreteFloatValue {
	return &ConcreteFloatValue{v}
}