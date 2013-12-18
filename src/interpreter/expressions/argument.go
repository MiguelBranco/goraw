package expressions

import . "interpreter"

type Argument struct {
	id int
}

func (e *Argument) Execute(args []Value) Value {
	return args[e.id]
}

func NewArgument(id int) *Argument {
	return &Argument{id}
}