// Depth first search in directed graph.
package main

import "fmt"

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

// This function perform depth first search from a given source vertex and return visited vertices and path.
func (graph graphObject) depthFirstSearch(vertex int, visited *[10]bool, edgeTo *[10]int) {
	// Mark the vertex visited.
	visited[vertex] = true
	var currentNode = graph.vertices[vertex]
	// Find all adjacent vertices.
	for {
		if currentNode == nil {
			// No more adjacent node.
			return
		}
		if !visited[currentNode.value] {
			// Save the edge.
			edgeTo[currentNode.value] = vertex
			// Given adjacent node is not visited.
			graph.depthFirstSearch(currentNode.value, visited, edgeTo)
		}
		currentNode = currentNode.next
	}
}

// This function find path between given vertices.
func (graph graphObject) pathBetween(vertexA int, vertexB int) {
	var visited [10]bool
	var edgeTo [10]int
	graph.depthFirstSearch(vertexA, &visited, &edgeTo)
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
