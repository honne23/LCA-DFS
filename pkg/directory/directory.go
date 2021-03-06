package directory

import (
	"log"
	"sync"
)

//FindCommonManager will travel down the tree in parallel, cache parents, then find a matching lowest common manager.
func FindCommonManager(root Manager, e1 Member, e2 Member) *Manager {

	results := make([]map[int][]Member, 2)
	// Search for both sub-trees in concurrently
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		baseCase := map[int][]Member{0: {&root}}
		if e1.GetID() == root.GetID() {
			results[0] = baseCase
		} else {
			results[0] = findByIdDFS(&root, e1.GetID(), 1, map[int][]Member{0: {&root}})
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		if e2.GetID() == root.GetID() {
			results[1] = map[int][]Member{0: {&root}}
		} else {
			results[1] = findByIdDFS(&root, e2.GetID(), 1, map[int][]Member{0: {&root}})
		}
		wg.Done()
	}()
	wg.Wait()
	minTree := -1
	if len(results[0]) < len(results[1]) {
		minTree = 0
	} else {
		minTree = 1
	}
	lowestCommon := func() *Manager {
		for depth := len(results[minTree]) - 1; depth > -1; depth-- {
			// Given two direct paths for two nodes, find the first node with the same id, this is the common ancestor
			if results[minTree][depth][0].GetID() == results[(minTree+1)%2][depth][0].GetID() {
				return (results[minTree][depth][0]).(*Manager)
			}
		}
		return nil
	}()
	log.Println(results)
	return lowestCommon

}

//findByIdDFS Uses a depth first search algorithm to trace a route to a given node.
// This function is tail recursive.
func findByIdDFS(node *Manager, id string, currentDepth int, parents map[int][]Member) map[int][]Member {

	if exists, _ := containsID(node.GetEmployees(), id); exists {
		// Once we have found a node, add it to the end of the queue
		parents[currentDepth] = []Member{node}
		return parents
	}
	for _, child := range node.GetEmployees() {
		switch child := child.(type) {
		case *Manager:
			parents = findByIdDFS(child, id, currentDepth+1, parents)
			if len(parents) > 1 {
				// Build the rest of the trace from the inside out
				parents[currentDepth] = []Member{node}
				return parents
			}

		default:
			continue
		}

	}
	return parents
}

func containsID(reports []Member, id string) (bool, *Member) {
	for _, member := range reports {
		if member.GetID() == id {
			return true, &member
		}
	}

	return false, nil
}
