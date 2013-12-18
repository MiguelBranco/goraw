package interpreter

type CollectionValue interface {
	Value
	NewCursor() Cursor
}