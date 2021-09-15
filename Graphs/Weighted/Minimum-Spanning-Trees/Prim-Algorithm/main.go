// Minimum spanning tree using prim algorithm on weighted graph.
package main

import "fmt"

// Priority queue object with edge nodes.
type queueObject struct {
	first *edgeObject
}

// This function add edge to the queue based on weight.
func (queue *queueObject) enqueue(edge edgeObject) {
	var newEdge = edgeObject{vertexA: edge.vertexA, vertexB: edge.vertexB, weight: edge.weight}
	if queue.first == nil {
		// No object in queue.
		queue.first = &newEdge
		return
	}
	if newEdge.weight < queue.first.weight {
		// New edge is smaller than first.
		newEdge.next = queue.first
		queue.first = &newEdge
		return
	}
	var currentNode = queue.first
	for {
		// Check weights.
		if currentNode.next == nil {
			// Add new edge here.
			currentNode.next = &newEdge
			return
		}
		if newEdge.weight < currentNode.next.weight {
			// Add new node in between.
			newEdge.next = currentNode.next
			currentNode.next = &newEdge
			return
		}
		// Move to next node.
		currentNode = currentNode.next
	}
}

// This function remove the first edge (smallest) from queue.
func (queue *queueObject) dequeue() edgeObject {
	var edge = *queue.first
	queue.first = queue.first.next
	return edge
}

// This functon check if queue is empty.
func (queue queueObject) isEmpty() bool {
	return queue.first == nil
}

// This function shows values of queue.
func (queue queueObject) show() {
	if queue.first == nil {
		// Empty queue.
		return
	}
	var node = queue.first
	for {
		fmt.Println(node.vertexA, node.vertexB, node.weight)
		if node.next == nil {
			return
		}
		node = node.next
	}
}

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

// This function return other vertex in a edge.
func (edge edgeObject) other(vertex int) int {
	if vertex == edge.vertexA {
		return edge.vertexB
	}
	return edge.vertexA
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

// This function add edge in graph between given vertex.
func (graph *graphObject) addEdge(vertexA int, vertexB int, weight float64) {
	graph.addNode(vertexA, vertexB, weight)
	graph.addNode(vertexB, vertexA, weight)
}

// This function find minimum spanning tree using prim.
func (graph graphObject) minimumSpanningtree() {
	var visited [10]bool
	var sortedQueue = queueObject{}
	var mst = queueObject{}
	// Start from vertex 0.
	graph.visit(0, &visited, &sortedQueue)
	for {
		if sortedQueue.isEmpty() {
			// No more edges show mst.
			mst.show()
			return
		}
		// Get shortest edge from queue.
		var edge = sortedQueue.dequeue()
		if visited[edge.vertexA] && visited[edge.vertexB] {
			// Both are visited.
			continue
		}
		mst.enqueue(edge)
		if !visited[edge.vertexA] {
			graph.visit(edge.vertexA, &visited, &sortedQueue)
		}
		if !visited[edge.vertexB] {
			graph.visit(edge.vertexB, &visited, &sortedQueue)
		}
	}
}

// This function add adjacent edges in sorted queue.
func (graph graphObject) visit(vertex int, visited *[10]bool, sortedQueue *queueObject) {
	visited[vertex] = true
	var currentNode = graph.vertices[vertex]
	// For all adjacent edges.
	for {
		if currentNode == nil {
			// No more adjacent node.
			return
		}
		var otherVertex = currentNode.other(vertex)
		if !visited[otherVertex] {
			// If the adjacent vertex is not visited add to queue.
			sortedQueue.enqueue(*currentNode)
		}
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
	g.minimumSpanningtree()
}
