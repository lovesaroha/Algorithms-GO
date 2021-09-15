// Graph using adjancency list.
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

// This function show adjacent vertices.
func (graph graphObject) adjacent(vertex int) {
	var currentNode = graph.vertices[vertex]
	for {
		if currentNode == nil {
			// No node found.
			return
		}
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
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
	g.adjacent(0)
}
