package operators

import (
	. "interpreter"
	. "interpreter/values"
)

func Select(cin chan []Value, p Expression, cout chan []Value) {
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			cout <- in
		}
	}
	close(cout)
}
