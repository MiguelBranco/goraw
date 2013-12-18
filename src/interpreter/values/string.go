package values

import . "interpreter"

// StringValue

type StringValue interface {
	Get() string
}

// ConcreteStringValue

type ConcreteStringValue struct {
	v string
}

func (this *ConcreteStringValue) Equal(that Value) bool {
	switch that := that.(type) {
	case StringValue:
		return this.v == that.Get()
	default:
		return false
	}
}

func (s *ConcreteStringValue) Get() string {
	return s.v
}

func NewConcreteStringValue(v string) *ConcreteStringValue {
	return &ConcreteStringValue{v}
}