package directory

import (
	"sync"
)

/**
* Travel down the tree, remember parents, find matching subtree for second employee
 */

func FindCommonManger(root Manager, e1 Member, e2 Member) Manager {
	wg := sync.WaitGroup{}
	results := make([]map[int][]*Manager, 2)
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

// Root node = ceo
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
			continue

		}

	}
	return nil
}

func contains(reports []*Manager, target Manager) bool {
	for _, manager := range reports {
		if (*manager).GetID() == target.GetID() {
			return true
		}
	}

	return false
}
