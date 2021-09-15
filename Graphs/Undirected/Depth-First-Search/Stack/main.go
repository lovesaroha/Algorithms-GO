// Depth first search in graph (adjacency list) with stack.
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

// Graph object structure.
type graphObject struct {
	vertices [10]*nodeObject
}

// Vertex node object.
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

// This function add edge in graph between given vertex.
func (graph *graphObject) addEdge(vertexA int, vertexB int) {
	graph.addNode(vertexA, vertexB)
	graph.addNode(vertexB, vertexA)
}

// This function perform depth first search from given source and return visited vertices and path.
func (graph graphObject) depthFirstSearch(source int) ([10]bool, [10]int) {
	var visited [10]bool
	var edgeTo [10]int
	var stack = stackObject{}
	stack.push(source)
	for {
		if stack.isEmpty() {
			// Stop loop if stack is empty.
			return visited, edgeTo
		}
		// Get vertex from stack mark it visited.
		var vertex = stack.pop()
		visited[vertex] = true
		// For all adjacent vertices.
		var currentNode = graph.vertices[vertex]
		for {
			if currentNode == nil {
				// No adjacent node.
				break
			}
			if !visited[currentNode.value] {
				// Save edge info and add node value in stack.
				edgeTo[currentNode.value] = vertex
				stack.push(currentNode.value)
			}
			currentNode = currentNode.next
		}
	}
}

// This function show path between given vertices.
func (graph graphObject) pathBetween(vertexA int, vertexB int) {
	// Perform depth first search from vertex A.
	visited, edgeTo := graph.depthFirstSearch(vertexA)
	if !visited[vertexB] {
		// Vertex b is not connected to vertex a so no path.
		fmt.Println("No path between ", vertexA, vertexB)
		return
	}
	var currentVertex = vertexB
	for {
		if edgeTo[currentVertex] == 0 {
			// End of path.
			return
		}
		fmt.Println(currentVertex, "-", edgeTo[currentVertex])
		currentVertex = edgeTo[currentVertex]
	}
}

func main() {

	// Create a graph.
	g := graphObject{}
	g.addEdge(3, 4)
	g.addEdge(2, 3)
	g.addEdge(0, 2)
	g.addEdge(7, 5)
	g.addEdge(0, 4)
	g.addEdge(0, 3)
	g.addEdge(3, 5)
	g.addEdge(7, 4)
	g.pathBetween(7, 0)
}
