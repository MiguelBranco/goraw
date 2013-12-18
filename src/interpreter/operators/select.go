package operators

import (
	. "interpreter"
	. "interpreter/values"
)

func Select(cin chan []Value, p Expression, cout chan []Value) {
	for in := range cin {
		switch v := p.Execute(in).(type) {
		case BoolValue:
			if v.Get() {
				cout <- in
			}
		default:
			panic("invalid type")
		}
	}
	close(cout)
}
