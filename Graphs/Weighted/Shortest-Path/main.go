// Shortest path in directed acyclic graph using topological order.
package main

import "fmt"

// Node object.
type nodeObject struct {
	value int
	next  *nodeObject
}

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

// This function perform depth first search from given vertex.
func (graph graphObject) depthFirstSearch(vertex int, visited *[8]bool, order *stackObject) {
	visited[vertex] = true
	// For all adjacent vertices.
	var currentNode = graph.vertices[vertex]
	for {
		if currentNode == nil {
			// No more adjacent node.
			break
		}
		if !visited[currentNode.vertexB] {
			// If node is not visited.
			graph.depthFirstSearch(currentNode.vertexB, visited, order)
		}
		currentNode = currentNode.next
	}
	order.push(vertex)
}

// This function find the topological order.
func (graph graphObject) topologicalOrder() stackObject {
	var order stackObject
	var visited [8]bool
	// For every new vertex perform depth first search.
	for i := 0; i < 8; i++ {
		if !visited[i] {
			// Perform depth first search.
			graph.depthFirstSearch(i, &visited, &order)
		}
	}
	return order
}

// This function find shortest path using topological order.
func (graph graphObject) shortestPath(vertexA int, vertexB int) {
	var distanceTo [8]float64
	var edgeTo [8]int
	// Find topological order usind depth first search.
	order := graph.topologicalOrder()
	for {
		if order.isEmpty() {
			// No more vertices.
			break
		}
		// For every vertex in topological order update distance and edge.
		var vertex = order.pop()
		graph.visit(vertex, &distanceTo, &edgeTo)
	}
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

// This function visit given vertex.
func (graph graphObject) visit(vertex int, distanceTo *[8]float64, edgeTo *[8]int) {
	// Find all adjacent vertices.
	var currentNode = graph.vertices[vertex]
	for {
		if currentNode == nil {
			// No more adjacent vertex.
			return
		}
		if distanceTo[currentNode.vertexB] == 0 || distanceTo[currentNode.vertexB] > distanceTo[vertex]+currentNode.weight {
			// Change distance value.
			distanceTo[currentNode.vertexB] = distanceTo[vertex] + currentNode.weight
			edgeTo[currentNode.vertexB] = vertex
		}
		currentNode = currentNode.next
	}
}

func main() {
	// Create a weighted directed acyclic graph.
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
