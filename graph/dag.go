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

func (g *GraphExec) dfs(node *Node, visited, tempMarked map[string]bool, result *[]string) error {
	if tempMarked[node.Out] {
		fmt.Println(node, node.Out)
		return fmt.Errorf("graph contains a cycle, cannot perform topological sort")
	}

	if !visited[node.Out] {
		tempMarked[node.Out] = true

		for _, neighbor := range node.Deps {
			if err := g.dfs(neighbor, visited, tempMarked, result); err != nil {
				return err
			}
		}
		visited[node.Out] = true
		tempMarked[node.Out] = false
		*result = append(*result, node.Out)
	}
	return nil
}

func (g *GraphExec) TopologicalSort() ([]string, error) {
	visited := make(map[string]bool)
	result := []string{}
	tempMarked := make(map[string]bool)

	for _, node := range g.Nodes {
		if !visited[node.Out] {
			if err := g.dfs(node, visited, tempMarked, &result); err != nil {
				return nil, err
			}
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result, nil
}

func (g *GraphExec) Execute(map[string]any) {

	ordering, e := g.TopologicalSort()
	if e != nil {
		panic(e)
	}
	// Need to reverse since did topological sort backwards
	for i, j := 0, len(ordering)-1; i < j; i, j = i+1, j-1 {
		ordering[i], ordering[j] = ordering[j], ordering[i]
	}
	fmt.Println(ordering)

}
