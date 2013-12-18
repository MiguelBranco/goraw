package interpreter

type Plugin interface {
	Init()
	Fini()
	GetCollection() CollectionValue
}