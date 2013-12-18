package interpreter

type Cursor interface {
	IsDone() bool
	Next() Value
	Close()
}