// Depth first search in graph (adjacency list).
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
func (graph graphObject) depthFirstSearch(vertex int, visited *[10]bool, edgeTo *[10]int) {
	visited[vertex] = true
	var currentNode = graph.vertices[vertex]
	// All adjacent vertices.
	for {
		if currentNode == nil {
			// No adjacent node.
			return
		}
		if !visited[currentNode.value] {
			// Save path info with vertex value.
			edgeTo[currentNode.value] = vertex
			// Search if node is not visited.
			graph.depthFirstSearch(currentNode.value, visited, edgeTo)
		}
		currentNode = currentNode.next
	}
}

// This function show path between given vertices.
func (graph graphObject) pathBetween(vertexA int, vertexB int) {
	var visited [10]bool
	var edgeTo [10]int
	// Perform depth first search from vertex A.
	graph.depthFirstSearch(vertexA, &visited, &edgeTo)
	if !visited[vertexB] {
		// Vertex b is not connected to vertex a so no path.
		fmt.Println("No path between ", vertexA, vertexB)
		return
	}
	var currentVertex = vertexB
	for {
		if currentVertex == vertexA {
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
