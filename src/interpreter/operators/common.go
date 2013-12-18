package operators

import . "interpreter"

// Auxiliary method to create a slice of values with a single value
func WrapValue(v Value) []Value {
	out := make([]Value, 1)
	out[0] = v
	return out 	
}