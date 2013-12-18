package operators

import (
	. "interpreter"
	. "interpreter/values"
)

func BindPath(path *Path, values []Value) CollectionValue {
	v := values[path.Arg]
	for _, n := range path.Children {
		rv, ok := v.(RecordValue)
		if !ok {
			panic("expected record value")
		}
		v = rv.GetValueByName(n)
	}
	cv, ok := v.(CollectionValue)
	if !ok {
		panic("expected collection value")
	}
	return cv
}

func Unnest(cin chan []Value, path *Path, p Expression, cout chan []Value) {
	for in := range cin {
		c := BindPath(path, in).NewCursor()
	    for !c.IsDone() {
			out := append(in, c.Next())
			switch v := p.Execute(out).(type) {
			case BoolValue:
				if v.Get() {
					cout <- out
				}
			default:
				panic("invalid type")
			}
	    }
	    c.Close()
	}
	close(cout)
}