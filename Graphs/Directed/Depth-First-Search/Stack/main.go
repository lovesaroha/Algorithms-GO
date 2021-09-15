// Depth first search in directed graph (adjacency list) with stack.
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
	vertices [10]*nodeObject
}

// Node object.
type nodeObject struct {
	value int
	next  *nodeObject
}

// This function add node to given vertex.
func (graph *graphObject) addNode(vertex int, value int) {
	var newNode = nodeObject{value: value}
	if graph.vertices[vertex] == nil {
		// No node is added.
		graph.vertices[vertex] = &newNode
		return
	}
	newNode.next = graph.vertices[vertex]
	graph.vertices[vertex] = &newNode
}

// This function add directed path between given vertices.
func (graph *graphObject) addEdge(vertexA int, vertexB int) {
	graph.addNode(vertexA, vertexB)
}

// This function perform depth first search usign stack and return visited vertices and path.
func (graph graphObject) depthFirstSearch(source int) ([10]bool, [10]int) {
	var visited [10]bool
	var edgeTo [10]int
	// Create a stack and add source vertex.
	var stack = stackObject{}
	stack.push(source)
	for {
		if stack.isEmpty() {
			// No more vertices in stack.
			return visited, edgeTo
		}
		// Get vertex from stack and mark visited.
		var vertex = stack.pop()
		visited[vertex] = true
		// Find all adjacent vertices.
		var currentNode = graph.vertices[vertex]
		for {
			if currentNode == nil {
				// No more node.
				break
			}
			if !visited[currentNode.value] {
				// Current node is not visited add to stack.
				stack.push(currentNode.value)
				edgeTo[currentNode.value] = vertex
			}
			currentNode = currentNode.next
		}
	}
}

// This function find path between given vertices.
func (graph graphObject) pathBetween(vertexA int, vertexB int) {
	visited, edgeTo := graph.depthFirstSearch(vertexA)
	if !visited[vertexB] {
		// Vertex b is not connected to vertex a.
		fmt.Println("No directed path from ", vertexA, " to ", vertexB)
		return
	}

	// Show the path.
	var currentVertex = vertexB
	for {
		if currentVertex == vertexA {
			// Route end.
			return
		}
		fmt.Println("From ", edgeTo[currentVertex], "->", currentVertex)
		currentVertex = edgeTo[currentVertex]
	}
}

func main() {

	// Create a graph.
	g := graphObject{}
	g.addEdge(3, 4)
	g.addEdge(2, 3)
	g.addEdge(0, 2)
	g.addEdge(5, 7)
	g.addEdge(0, 4)
	g.addEdge(0, 3)
	g.addEdge(3, 5)
	g.addEdge(7, 4)
	g.pathBetween(0, 7)
}
