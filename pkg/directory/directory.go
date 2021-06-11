package directory

import (
	"log"
	"reflect"
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
			results[0] = findByIdBFS(&root, e1.GetID(), 1, map[int][]Member{0: {&root}})
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		if e2.GetID() == root.GetID() {
			results[1] = map[int][]Member{0: {&root}}
		} else {
			results[1] = findByIdBFS(&root, e2.GetID(), 1, map[int][]Member{0: {&root}})
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
			for _, member := range results[minTree][depth] {
				switch member := member.(type) {
				case *Manager:
					equal := reflect.DeepEqual(member.GetEmployees(), results[(minTree+1)%2][depth+1])
					if equal {
						return member
					}
				default:
					continue
				}
			}
		}
		return nil
	}()

	return lowestCommon

}

//findByIdBFS Uses a breadth first search algorithm to discover all the managers of a given node.
// This function is tail recursive.
func findByIdBFS(node *Manager, id string, currentDepth int, parents map[int][]Member) map[int][]Member {
	if _, exists := parents[currentDepth]; !exists {
		parents[currentDepth] = node.GetEmployees()
	} else {
		parents[currentDepth] = append(parents[currentDepth], node.GetEmployees()...)
	}
	if containsID(node.GetEmployees(), id) {
		for depth := 0; depth < len(parents); depth++ {
			for _, child := range parents[depth] {
				log.Printf("%d | %s", depth, child.GetID())
			}
		}
		return parents
	}

	for _, child := range node.GetEmployees() {
		switch child := child.(type) {
		case *Manager:
			findByIdBFS(child, id, currentDepth+1, parents)
		default:
			continue
		}

	}
	return parents
}

func containsID(reports []Member, id string) bool {
	for _, member := range reports {
		if member.GetID() == id {
			return true
		}
	}

	return false
}
