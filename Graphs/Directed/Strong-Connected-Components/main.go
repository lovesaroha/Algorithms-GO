// Strong connected components in directed graph using kosaraju-sharir algorithm.
package main

import "fmt"

// Stack object defined.
type stackObject struct {
	first *nodeObject
}

// Push function to add value in stack.
func (stack *stackObject) push(value int) {
	var newNode = nodeObject{next: stack.first, value: value}
	stack.first = &newNode
}

// Pop function remove value from stack.
func (stack *stackObject) pop() int {
	var value int
	if stack.first == nil {
		// Empty stack.
		return 0
	}
	value = stack.first.value
	stack.first = stack.first.next
	return value
}

// This function checks if stack is empty.
func (stack stackObject) isEmpty() bool {
	return stack.first == nil
}

// Graph object.
type graphObject struct {
	vertices [13]*nodeObject
}

// Node object.
type nodeObject struct {
	value int
	next  *nodeObject
}

// This function add node to given vertex.
func (graph *graphObject) addNode(vertex int, value int) {
	var currentNode = graph.vertices[vertex]
	if currentNode == nil {
		// No node is added.
		graph.vertices[vertex] = &nodeObject{value: value}
		return
	}
	for {
		if currentNode.next == nil {
			// Add new node.
			currentNode.next = &nodeObject{value: value}
			return
		}
		currentNode = currentNode.next
	}
}

// This function add directed path between given vertices.
func (graph *graphObject) addEdge(vertexA int, vertexB int) {
	graph.addNode(vertexA, vertexB)
}

// This function reverse the given graph.
func (graph graphObject) reverse() graphObject {
	graphReverse := graphObject{}
	for i := 0; i < 13; i++ {
		// For every vertex reverse path direction.
		var currentNode = graph.vertices[i]
		for {
			if currentNode == nil {
				// No more adjacent vertices.
				break
			}
			graphReverse.addEdge(currentNode.value, i)
			currentNode = currentNode.next
		}
	}
	return graphReverse
}

// This function perform depth first search from a given source vertex and change visited vertices and path.
func (graph graphObject) depthFirstSearch(vertex int, visited *[13]bool, edgeTo *[13]int, order *stackObject) {
	// Mark the vertex visited.
	visited[vertex] = true
	var currentNode = graph.vertices[vertex]
	// Find all adjacent vertices.
	for {
		if currentNode == nil {
			// No more adjacent node.
			break
		}
		if !visited[currentNode.value] {
			// Save the edge.
			edgeTo[currentNode.value] = vertex
			// Given adjacent node is not visited.
			graph.depthFirstSearch(currentNode.value, visited, edgeTo, order)
		}
		currentNode = currentNode.next
	}
	order.push(vertex)
}

// This function perform topological sort.
func (graph graphObject) topologicalSort() stackObject {
	var visited [13]bool
	var edgeTo [13]int
	var order = stackObject{}
	// For every unvisited vertices perform depth first search.
	for i := 0; i < 13; i++ {
		if !visited[i] {
			// If vertex is not visited.
			graph.depthFirstSearch(i, &visited, &edgeTo, &order)
		}
	}
	return order
}

// This function return strong connected components in graph using kosaraju-sharir algorithm.
func (graph graphObject) strongConnectedComponents() {
	var totalComponents int
	// First reverse the given graph.
	graphReverse := graph.reverse()
	// Find topological order of reverse graph.
	order := graphReverse.topologicalSort()
	// Now perform depth first search on given order.
	var visited [13]bool
	var edgeTo [13]int
	for {
		if order.isEmpty() {
			// No more value in stack.
			break
		}
		// Get value from stack.
		var vertex = order.pop()
		if !visited[vertex] {
			// Perform depth first search if vertex is not visited yet.
			graph.depthFirstSearch(vertex, &visited, &edgeTo, &stackObject{})
			totalComponents++
		}
	}
	fmt.Println("Total strong connected components : ", totalComponents)
}

func main() {
	// Create a graph.
	g := graphObject{}
	g.addEdge(0, 1)
	g.addEdge(0, 5)
	g.addEdge(5, 4)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 2)
	g.addEdge(3, 5)
	g.addEdge(4, 3)
	g.addEdge(4, 2)
	g.addEdge(6, 0)
	g.addEdge(6, 8)
	g.addEdge(6, 9)
	g.addEdge(8, 6)
	g.addEdge(7, 6)
	g.addEdge(7, 9)
	g.addEdge(9, 11)
	g.addEdge(9, 10)
	g.addEdge(11, 4)
	g.addEdge(11, 12)
	g.addEdge(12, 9)
	g.addEdge(10, 12)
	g.strongConnectedComponents()
}
