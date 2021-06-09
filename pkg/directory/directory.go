package directory

import (
	"sync"
)

//FindCommonManager will travel down the tree in parallel, cache parents, then find a matching lowest common manager.
func FindCommonManager(root Manager, e1 Member, e2 Member) Manager {
	wg := sync.WaitGroup{}
	results := make([]map[int][]*Manager, 2)
	// Search for both sub-trees in concurrently
	wg.Add(1)
	go func() {
		// Only reads so no race-conditions
		// Edgecase:  One of the nodes is a root node
		if e1.GetID() == root.GetID() {
			results[0] = map[int][]*Manager{0: {&root}}
		} else {
			results[0] = findByIdDFS(&root, e1.GetID(), 0, map[int][]*Manager{})
		}

		wg.Done()
	}()
	wg.Add(1)
	go func() {
		if e1.GetID() == root.GetID() {
			results[1] = map[int][]*Manager{0: {&root}}
		} else {
			results[1] = findByIdDFS(&root, e2.GetID(), 0, map[int][]*Manager{})
		}
		wg.Done()
	}()
	wg.Wait()

	var lowestCommon Manager
	for depth, parents := range results[0] {
		for _, node := range parents {
			if contains(results[1][depth], *node) {
				lowestCommon = *node
			}
		}
	}
	return lowestCommon

}

//findByIdDFS Uses a depth first search algorithm to discover all the managers of a given node.
// This function is tail recursive.
func findByIdDFS(node *Manager, id string, currentDepth int, parents map[int][]*Manager) map[int][]*Manager {

	for _, child := range node.GetEmployees() {
		if child.GetID() == id {
			return parents
		}

		switch child := child.(type) {
		case *Manager:

			if len(child.GetEmployees()) < 1 {
				continue
			}
			if _, exists := parents[currentDepth+1]; !exists {
				parents[currentDepth] = []*Manager{child}
			} else {
				if !contains(parents[currentDepth+1], *child) {
					parents[currentDepth] = append(parents[currentDepth], child)
				}
			}
			return findByIdDFS(child, id, currentDepth+1, parents)
		default:
			// In case a child is a leaf node, do not explore them
			continue

		}

	}
	return nil
}

//contains is a util function to check if a particular Manager is contained in a slice of Managers
func contains(reports []*Manager, target Manager) bool {
	for _, manager := range reports {
		if (*manager).GetID() == target.GetID() {
			return true
		}
	}

	return false
}
