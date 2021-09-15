// Breadth first search in graph (adjacency list).
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

// Graph object structure.
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

// This function add edge in graph between given vertex.
func (graph *graphObject) addEdge(vertexA int, vertexB int) {
	graph.addNode(vertexA, vertexB)
	graph.addNode(vertexB, vertexA)
}

// This function perform breadth first search from given source and return visited vertices and path.
func (graph graphObject) breadthFirstSearch(source int) ([10]bool, [10]int) {
	var visited [10]bool
	var edgeTo [10]int
	queue := queueObject{}
	queue.enqueue(source)
	visited[source] = true
	for {
		if queue.isEmpty() {
			// Close loop if queue is empty.
			return visited, edgeTo
		}
		// Get vertex from queue.
		var vertex = queue.dequeue()
		// For all adjacent vertices.
		var currentNode = graph.vertices[vertex]
		for {
			if currentNode == nil {
				// No node.
				break
			}
			if !visited[currentNode.value] {
				// Save edge info and add node value in queue.
				edgeTo[currentNode.value] = vertex
				queue.enqueue(currentNode.value)
				visited[currentNode.value] = true
			}
			currentNode = currentNode.next
		}
	}
}

// This function show path between given vertices.
func (graph graphObject) pathBetween(vertexA int, vertexB int) {
	// Perform breadth first search from vertex A.
	visited, edgeTo := graph.breadthFirstSearch(vertexA)
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
