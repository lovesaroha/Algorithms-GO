// Weighted graph using adjancency list.
package main

import "fmt"

// Graph object structure.
type graphObject struct {
	vertices [10]*edgeObject
}

// Edge object.
type edgeObject struct {
	vertexA int
	vertexB int
	weight  float64
	next    *edgeObject
}

// This function add node to given vertex.
func (graph *graphObject) addNode(vertexA int, vertexB int, weight float64) {
	var newEdge = edgeObject{vertexA: vertexA, vertexB: vertexB, weight: weight}
	if graph.vertices[vertexA] == nil {
		// No node is added.
		graph.vertices[vertexA] = &newEdge
		return
	}
	newEdge.next = graph.vertices[vertexA]
	graph.vertices[vertexA] = &newEdge
}

// This function add edge in graph between given vertex.
func (graph *graphObject) addEdge(vertexA int, vertexB int, weight float64) {
	graph.addNode(vertexA, vertexB, weight)
	graph.addNode(vertexB, vertexA, weight)
}

// This function show adjacent vertices.
func (graph graphObject) adjacent(vertex int) {
	var currentNode = graph.vertices[vertex]
	for {
		if currentNode == nil {
			// No node found.
			return
		}
		fmt.Println(currentNode.vertexA, currentNode.vertexB, currentNode.weight)
		currentNode = currentNode.next
	}
}

func main() {

	// Create a weighted graph.
	g := graphObject{}
	g.addEdge(3, 4, 0.12)
	g.addEdge(2, 3, 1.09)
	g.addEdge(0, 2, 0.45)
	g.addEdge(7, 5, 2.30)
	g.addEdge(0, 4, 0.11)
	g.addEdge(4, 7, 0.15)
	g.addEdge(5, 6, 0.89)
	g.addEdge(5, 8, 0.90)
	g.addEdge(6, 9, 0.66)
	g.adjacent(0)
}
