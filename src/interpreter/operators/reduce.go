package operators

import (
	. "interpreter"
	. "interpreter/values"
)

func sumInt(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	acc := 0
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in).(IntValue)
			acc += nv.Get()
		}
	}
	cout <- WrapValue(NewConcreteIntValue(acc))
	close(cout)
}

func sumFloat(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	acc := 0.0
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in).(FloatValue)
			acc += nv.Get()
		}
	}
	cout <- WrapValue(NewConcreteFloatValue(acc))
	close(cout)
}

func multInt(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	acc := 1
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in).(IntValue)
			acc *= nv.Get()
		}
	}
	cout <- WrapValue(NewConcreteIntValue(acc))
	close(cout)
}

func multFloat(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	acc := 1.0
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in).(FloatValue)
			acc *= nv.Get()
		}
	}
	cout <- WrapValue(NewConcreteFloatValue(acc))
	close(cout)
}

func maxInt(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	first := true
	acc := 0
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in).(IntValue)
			if first || nv.Get() > acc {
				acc = nv.Get()
				first = false
			}
		}
	}
	cout <- WrapValue(NewConcreteIntValue(acc))
	close(cout)
}

func maxFloat(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	first := true
	acc := 0.0
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in).(FloatValue)
			if first || nv.Get() > acc {
				acc = nv.Get()
				first = false
			}
		}
	}
	cout <- WrapValue(NewConcreteFloatValue(acc))
	close(cout)
}

func or(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	acc := false
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in).(BoolValue)
			acc = acc || nv.Get()
		}
	}
	cout <- WrapValue(NewConcreteBoolValue(acc))
	close(cout)
}

func and(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	acc := true
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in).(BoolValue)
			acc = acc && nv.Get()
		}
	}
	cout <- WrapValue(NewConcreteBoolValue(acc))
	close(cout)
}

func unionSet(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	var set map[Value]bool

	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in)
			set[nv] = true
		}
	}

	for k := range set {
		cout <- WrapValue(k)
	}
	close(cout)
}

func unionBag(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	var bag map[Value]int

	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in)
			if n, ok := bag[nv]; ok {
				bag[nv] = n + 1
			} else {
				bag[nv] = 1
			}
		}
	}

	for k, v := range bag {
		for i := 0; i < v; i++ {
			cout <- WrapValue(k)
		}
	}
	close(cout)
}

func appendList(cin chan []Value, e Expression, p Expression, cout chan []Value) {
	for in := range cin {
		v := p.Execute(in).(BoolValue)
		if v.Get() {
			nv := e.Execute(in)
			cout <- WrapValue(nv)	
		}
	}
	close(cout)
}


func Reduce(cin chan []Value, acc Accumulator, e Expression, p Expression, cout chan []Value) {
	switch acc {
	case Sum:
		switch e.Type() {
		case Int: sumInt(cin, e, p, cout)
		case Float: sumFloat(cin, e, p, cout)
		}
	case Multiply:
		switch e.Type() {
		case Int: multInt(cin, e, p, cout)
		case Float: multFloat(cin, e, p, cout)
		}
	case Max:
		switch e.Type() {
		case Int: maxInt(cin, e, p, cout)
		case Float: maxFloat(cin, e, p, cout)
		}
	case Or:
		or(cin, e, p, cout)
	case And:
		and(cin, e, p, cout)
	case Union:
		unionSet(cin, e, p, cout)
	case BagUnion:
		unionBag(cin, e, p, cout)
	case Append:
		appendList(cin, e, p, cout)
	default: panic("invalid accumulator")
	}
}