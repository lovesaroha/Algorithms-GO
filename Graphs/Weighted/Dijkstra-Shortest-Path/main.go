// Dijkstra algorithm to find shortes path in weighted directed graph.
package main

import "fmt"

// Graph object.
type graphObject struct {
	vertices [8]*edgeObject
}

// Edge object.
type edgeObject struct {
	vertexA int
	vertexB int
	weight  float64
	next    *edgeObject
}

// This function add node to given vertex in sorted order.
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

// This function add directed path between given vertices.
func (graph *graphObject) addEdge(vertexA int, vertexB int, weight float64) {
	graph.addNode(vertexA, vertexB, weight)
}

// This function perform dijkstra algorithm from given source to all vertices.
func (graph graphObject) shortestPathFrom(source int) ([8]float64, [8]int) {
	var distanceTo [8]float64
	var edgeTo [8]int
	var marked [8]bool
	var currentVertex = source
	for {
		if marked[currentVertex] {
			// No more vertex.
			return distanceTo, edgeTo
		}
		marked[currentVertex] = true
		// Find all adjacent vertices.
		var currentNode = graph.vertices[currentVertex]
		for {
			if currentNode == nil {
				// No more adjacent node.
				break
			}
			if distanceTo[currentNode.vertexB] == 0 || distanceTo[currentNode.vertexB] > distanceTo[currentVertex]+currentNode.weight {
				// Update distance.
				distanceTo[currentNode.vertexB] = distanceTo[currentVertex] + currentNode.weight
				edgeTo[currentNode.vertexB] = currentVertex
			}
			currentNode = currentNode.next
		}
		currentVertex = findNextCloseVertex(distanceTo, marked)
	}
}

// This function return next minimum distance vertex.
func findNextCloseVertex(distanceTo [8]float64, marked [8]bool) int {
	var vertex int
	for i := 0; i < 8; i++ {
		if marked[i] {
			// Vertex is already marked.
			continue
		}
		if (distanceTo[i] < distanceTo[vertex] || distanceTo[vertex] == 0) && distanceTo[i] != 0 {
			// Minimum.
			vertex = i
		}
	}
	return vertex
}

// This function shows shortest path.
func (graph graphObject) shortestPath(vertexA int, vertexB int) {
	// Perform dijkstra to calculate shortest path from vertex a to all vertices.
	distanceTo, edgeTo := graph.shortestPathFrom(vertexA)
	// Show path.
	var currentVertex = vertexB
	for {
		if currentVertex == vertexA {
			// Path found.
			fmt.Println("Total distance ", distanceTo[vertexB])
			return
		}
		fmt.Println("From ", edgeTo[currentVertex], " to ", currentVertex)
		currentVertex = edgeTo[currentVertex]
	}
}

func main() {
	// Create a weighted directed graph.
	g := graphObject{}
	g.addEdge(0, 1, 5.0)
	g.addEdge(0, 7, 8.0)
	g.addEdge(0, 4, 9.0)
	g.addEdge(1, 7, 4.0)
	g.addEdge(1, 2, 12.0)
	g.addEdge(1, 3, 15.0)
	g.addEdge(7, 2, 7.0)
	g.addEdge(7, 5, 6.0)
	g.addEdge(5, 2, 1.0)
	g.addEdge(5, 6, 13.0)
	g.addEdge(4, 7, 5.0)
	g.addEdge(4, 5, 4.0)
	g.addEdge(4, 6, 20.0)
	g.addEdge(2, 3, 3.0)
	g.addEdge(2, 6, 11.0)
	g.addEdge(3, 6, 9.0)
	g.shortestPath(0, 6)
}
