package values

import . "interpreter"

// BoolValue

type BoolValue interface {
	Get() bool
}

// ConcreteBoolValue

type ConcreteBoolValue struct {
	v bool
}

func (this *ConcreteBoolValue) Equal(that Value) bool {
	switch that := that.(type) {
	case BoolValue:
		return this.v == that.Get()
	default:
		return false
	}
}

func (b *ConcreteBoolValue) Get() bool {
	return b.v
}

func NewConcreteBoolValue(v bool) *ConcreteBoolValue {
	return &ConcreteBoolValue{v}
}

// ReferenceBoolValue

// FIXME: ... implement ...