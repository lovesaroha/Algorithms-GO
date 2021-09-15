// Breadth first search in directed graph.
package main

import "fmt"

// Queue object defined.
type queueObject struct {
	first *nodeObject
	last  *nodeObject
}

// Enqueue function add new value in queue.
func (queue *queueObject) enqueue(value int) {
	var newNode = nodeObject{value: value}
	if queue.last != nil {
		// Last node next pointer now points to new node.
		queue.last.next = &newNode
	}
	// Now last pointer point to new node.
	queue.last = &newNode
	if queue.first == nil {
		queue.first = &newNode
	}
}

// This function check if queue is empty.
func (queue queueObject) isEmpty() bool {
	return queue.first == nil
}

// Dequeue function remove value from queue and return the value.
func (queue *queueObject) dequeue() int {
	var value int
	if queue.first == queue.last {
		// Queue have single value.
		queue.last = nil
	}
	value = queue.first.value
	queue.first = queue.first.next
	return value
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

// This function perform breadth first search.
func (graph graphObject) breadthFirstSearch(source int) ([10]bool, [10]int) {
	var visited [10]bool
	var edgeTo [10]int
	// Create a queue and add source vertex.
	var queue = queueObject{}
	queue.enqueue(source)
	for {
		if queue.isEmpty() {
			// No more vertex in queue.
			return visited, edgeTo
		}
		// Get vertex from queue.
		var vertex = queue.dequeue()
		var currentNode = graph.vertices[vertex]
		visited[vertex] = true
		// Find all adjacent vertices.
		for {
			if currentNode == nil {
				// No more adjacent vertices.
				break
			}
			if !visited[currentNode.value] {
				// If adjacent vertex not visited add to queue.
				queue.enqueue(currentNode.value)
				edgeTo[currentNode.value] = vertex
			}
			currentNode = currentNode.next
		}
	}
}

// This function find path between given vertices.
func (graph graphObject) pathBetween(vertexA int, vertexB int) {
	visited, edgeTo := graph.breadthFirstSearch(vertexA)
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
