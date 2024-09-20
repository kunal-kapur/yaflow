package graph

type Node struct {
	Deps []Node
	Action []interface{} 
	Out string
}

