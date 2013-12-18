package operators

import (
	"fmt"
	"strconv"
	. "interpreter"
	. "interpreter/values"
)

func pprint(v Value) string {
	switch v := v.(type) {
	case BoolValue:
		return strconv.FormatBool(v.Get())
	case FloatValue:
		return strconv.FormatFloat(v.Get(), 'b', -1, 64)
	case IntValue:
		return strconv.Itoa(v.Get())
	case StringValue:
		return "\"" + v.Get() + "\""
	case RecordValue:
		s := "("
		for i := 0; i < v.Size(); i++ {
			s += v.GetName(i) + " := " + pprint(v.GetValue(i))
			if i != v.Size() - 1 {
				s += ", "
			}
		}
		s += ")"
		return s
	case NullValue:
		return "null"
	case CollectionValue:
		return "[...]"
	default:
		panic("invalid value")
	}
}

func Print(cin chan []Value) {
	for in := range cin {
		if len(in) > 0 {
			fmt.Print(pprint(in[0]))
			for i := 1; i < len(in); i++ {
				fmt.Print("\t")
				fmt.Print(pprint(in[i]))
			}
		}
		fmt.Println()
	}
}
