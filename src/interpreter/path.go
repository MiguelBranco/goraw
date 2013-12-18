package interpreter

type Path struct {
	Arg int
	Children []string
}

func NewPath(arg int, children []string) *Path {
	return &Path{arg, children}
}