package interpreter

type Accumulator int32

const (
	Sum Accumulator = iota + 1
	Multiply
	Max
	Or
	And
	Union
	BagUnion
	Append
)