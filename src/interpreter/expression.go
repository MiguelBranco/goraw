package interpreter

type ExpressionType int32

const (
	Bool ExpressionType = iota + 1
	Float
	Int
	String
	Record
)

type Expression interface {
	Execute(args []Value) Value
	Type() ExpressionType
}
