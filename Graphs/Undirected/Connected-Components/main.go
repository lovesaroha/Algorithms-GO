// Connected components in graph (adjacency list).
package main

import "fmt"

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
func (graph graphObject) depthFirstSearch(source int, visited *[10]bool) {
	// Start depth first search from source.
	graph.search(source, visited)
}

// This function perform search.
func (graph graphObject) search(vertex int, visited *[10]bool) {
	visited[vertex] = true
	var currentNode = graph.vertices[vertex]
	// All adjacent vertices.
	for {
		if currentNode == nil {
			// No adjacent node.
			return
		}
		if !visited[currentNode.value] {
			// Search if node is not visited.
			graph.search(currentNode.value, visited)
		}
		currentNode = currentNode.next
	}
}

// This function find connected components in graph.
func (graph graphObject) connectedComponents() {
	var totalComponents int
	var visited [10]bool
	for i := 0; i < 10; i++ {
		if !visited[i] {
			// If vertex is not visited perform depth first search from this vertex.
			graph.depthFirstSearch(i, &visited)
			totalComponents++
		}
	}
	fmt.Println("Total components : ", totalComponents)
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
	g.connectedComponents()
}
