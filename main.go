package main

import (
	"fmt"
	"yaflow/graph"
)

func add(a int, b int) int {
	return a + b
}

func main() {

	g := graph.CreateGraphExec()

	node1 := graph.Node{Deps: []*graph.Node{}, Action: add, Out: "node1"}
	node2 := graph.Node{Deps: []*graph.Node{&node1}, Action: add, Out: "node2"}
	node3 := graph.Node{Deps: []*graph.Node{&node1}, Action: add, Out: "node3"}

	g.AddChild(&node1)
	g.AddChild(&node2)
	g.AddChild(&node3)
	fmt.Println("Done")
	// g.AddChild(&node3)

}
