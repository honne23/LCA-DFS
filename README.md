# Lowest common ancestor using depth-first-search


Bureaucr.at is a typical hierarchical organisation. Claire, its CEO, has a hierarchy of employees reporting to her and each employee can have a list of other employees reporting to him/her. An employee with at least one report is called a Manager.

Your task is to implement a corporate directory for Bureaucr.at with an interface to find the closest common Manager (i.e. farthest from the CEO) between two employees. You may assume that all employees eventually report up to the CEO.

#### Use depth first search to trace a route to a leaf, then build the stack as you ascend back up the recursive calls
```go
//findByIdDFS Uses a depth first search algorithm to trace a route to a given node.
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
```

#### Finding the lowest common ancestor just searches for the first ID match between two stacks

#### To build:
```bash
go build -o optio main.go && chmod +x optio
```

#### Run:
```bash
./optio
```
#### Or:
```bash
go run main.go
```
