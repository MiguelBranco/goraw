package operators

import (
	. "interpreter"
	. "interpreter/values"
)

func BindPath(path *Path, values []Value) CollectionValue {
	v := values[path.Arg]
	for _, n := range path.Children {
		rv := v.(RecordValue)
		v = rv.GetValueByName(n)
	}
	return v.(CollectionValue)
}

func Unnest(cin chan []Value, path *Path, p Expression, cout chan []Value) {
	for in := range cin {
		c := BindPath(path, in).NewCursor()
	    for !c.IsDone() {
			out := append(in, c.Next())
			v := p.Execute(out).(BoolValue)
			if v.Get() {
				cout <- out
			}
	    }
	    c.Close()
	}
	close(cout)
}