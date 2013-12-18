package interpreter

type Value interface {
	Equal(that Value) bool
}