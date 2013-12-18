package operators

import (
	. "interpreter"
	. "interpreter/values"
)

func Join(clhs chan []Value, crhs chan []Value, p Expression, cout chan []Value) {
	matrhs := make([][]Value, 0)
	for in := range crhs {
		matrhs = append(matrhs, in)
	}

	for lhs := range clhs {
		for _, rhs := range matrhs {
			out := append(lhs, rhs...)
			v := p.Execute(out).(BoolValue)
			if v.Get() {
				cout <- out
			}
	    }
	}
	close(cout)
}