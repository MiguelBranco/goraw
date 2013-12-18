package expressions

import . "interpreter"

type Argument struct {
	t ExpressionType
	id int
}

func (e *Argument) Execute(args []Value) Value {
	return args[e.id]
}

func (e *Argument) Type() ExpressionType {
	return e.t
}

func NewArgument(t ExpressionType, id int) *Argument {
	return &Argument{t, id}
}