package graph

type Node struct {
	Deps   []*Node
	Action interface{}
	Out    string
}

type GraphExec struct {
	Nodes   []*Node
	NodeMap map[string]*Node
}

func CreateGraphExec() *GraphExec {
	return &GraphExec{
		Nodes:   []*Node{},
		NodeMap: map[string]*Node{},
	}
}


