package interpreter

type Expression interface {
	Execute(args []Value) Value
}