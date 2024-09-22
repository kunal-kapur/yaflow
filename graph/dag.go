package graph

import (
	"errors"
	"fmt"
)

func (g *GraphExec) dfsCycle(v *Node, visited, recStack map[string]bool) (bool, error) {

	_, exists := g.NodeMap[v.Out]

	if !exists {

		return false, errors.New(("No such node exists," + v.Out))
	}

	visited[v.Out] = true
	recStack[v.Out] = true

	// Recur for all vertices adjacent to this vertex
	for _, pred := range v.Deps {
		// If the neighbor is not visited, recurse on it
		if !visited[pred.Out] {
			res, err := g.dfsCycle(pred, visited, recStack)
			if res {
				return true, err
			}
		} else if recStack[pred.Out] {
			// If the neighbor is in the recursion stack, then we found a cycle
			return true, errors.New("CYCLE DETECTED")
		}
	}

	// Remove the vertex from recursion stack
	recStack[v.Out] = false
	return false, nil
}

func (g *GraphExec) CheckGraph() bool {
	visited := make(map[string]bool)  // stack for everything seen
	recStack := make(map[string]bool) // stack for recursion
	for _, v := range g.Nodes {
		visited[v.Out] = false
		recStack[v.Out] = false
	}

	for _, node := range g.Nodes {
		if !visited[node.Out] {
			res, err := g.dfsCycle(node, visited, recStack)
			errors_list := []error{err}
			if res {
				errors_list = append(errors_list, errors.New("CYCLE DETECTED"))
			}
			fmt.Println(errors_list)
			for _, item := range errors_list {
				panic(item)
			}
		}
	}
	return true

}

func (manager *GraphExec) AddChild(element *Node) {
	_, exists := manager.NodeMap[element.Out]
	if exists {
		panic(errors.New("ELEMENT EXISTS"))
	}
	manager.Nodes = append(manager.Nodes, element)
	manager.NodeMap[element.Out] = element
}
