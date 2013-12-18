package values

import . "interpreter"

type NullValue struct{}

func (n NullValue) Equal(that Value) bool {
	return false
}

func NewNullValue() *NullValue {
	return &NullValue{}
}