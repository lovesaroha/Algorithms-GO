// Max flow using ford fulkerson in flow network.
package main

import "fmt"

// Queue object defined.
type queueObject struct {
	first *nodeObject
	last  *nodeObject
}

// Node object.
type nodeObject struct {
	value int
	next  *nodeObject
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
	vertices [10]*edgeObject
}

// Edge object.
type edgeObject struct {
	vertexA  int
	vertexB  int
	capacity int
	flow     int
	next     *edgeObject
}

// This function return residual capacity from given vertex to other.
func (edge edgeObject) residualCapacityTo(vertex int) int {
	if vertex == edge.vertexB {
		// Forward from a to b.
		return edge.capacity - edge.flow
	}
	// Backward edge from b to a.
	return edge.flow
}

// This function add residual flow to from given vertex.
func (edge *edgeObject) addResidualFlowTo(vertex int, flow int) {
	if vertex == edge.vertexB {
		// Forward from a to b.
		edge.flow += flow
		return
	}
	// Backward edge from b to a.
	edge.flow -= flow
}

// This function add node to given vertex.
func (graph *graphObject) addNode(vertexA int, vertexB int, capacity int) {
	var newEdge = edgeObject{vertexA: vertexA, vertexB: vertexB, capacity: capacity}
	if graph.vertices[vertexA] == nil {
		// No node is added.
		graph.vertices[vertexA] = &newEdge
		return
	}
	newEdge.next = graph.vertices[vertexA]
	graph.vertices[vertexA] = &newEdge
}

// This function add edge in graph between given vertex.
func (graph *graphObject) addEdge(vertexA int, vertexB int, capacity int) {
	graph.addNode(vertexA, vertexB, capacity)
}

// This function perform ford fulkerson to find max flow.
func (graph graphObject) Maxflow(source int, target int) {
	var maxFlow int
	for {
		path, edgeTo := graph.hasAugmentedPath(source, target)
		if !path {
			// No more path from source to target.
			fmt.Println("Maximum flow : ", maxFlow)
			return
		}
		var minimumCapacity int
		// Find bottleneck or minimum capacity.
		var currentVertex = target
		for {
			if currentVertex == source {
				// Source vertex reached.
				break
			}
			if minimumCapacity == 0 || minimumCapacity > edgeTo[currentVertex].residualCapacityTo(currentVertex) {
				// Update minimum capacity.
				minimumCapacity = edgeTo[currentVertex].residualCapacityTo(currentVertex)
			}
			currentVertex = edgeTo[currentVertex].vertexA
		}
		// Add flow to the network.
		currentVertex = target
		for {
			if currentVertex == source {
				// Source vertex reached.
				break
			}
			edgeTo[currentVertex].addResidualFlowTo(currentVertex, minimumCapacity)
			currentVertex = edgeTo[currentVertex].vertexA
		}
		maxFlow += minimumCapacity
	}
}

// This function checks augmented path.
func (graph graphObject) hasAugmentedPath(source int, target int) (bool, [8]*edgeObject) {
	var visited [8]bool
	var queue = queueObject{}
	var edgeTo [8]*edgeObject
	queue.enqueue(source)
	visited[source] = true
	// Perform breadth first search.
	for {
		if queue.isEmpty() {
			break
		}
		// Get vertex from queue.
		var vertex = queue.dequeue()
		// For all adjacent vertices.
		var currentNode = graph.vertices[vertex]
		for {
			if currentNode == nil {
				// No more adjacent nodes.
				break
			}
			if !visited[currentNode.vertexB] && currentNode.residualCapacityTo(currentNode.vertexB) > 0 {
				// Add to queue if not visited and has residual capacity.
				queue.enqueue(currentNode.vertexB)
				visited[currentNode.vertexB] = true
				edgeTo[currentNode.vertexB] = currentNode
			}
			currentNode = currentNode.next
		}
	}
	// True if there is a path from source to target with capacity.
	return visited[target], edgeTo
}

func main() {

	// Create a flow network.
	g := graphObject{}
	g.addEdge(0, 1, 10)
	g.addEdge(0, 2, 5)
	g.addEdge(0, 3, 15)
	g.addEdge(1, 4, 9)
	g.addEdge(1, 2, 4)
	g.addEdge(1, 5, 15)
	g.addEdge(2, 5, 8)
	g.addEdge(2, 3, 4)
	g.addEdge(3, 6, 16)
	g.addEdge(4, 5, 15)
	g.addEdge(4, 7, 10)
	g.addEdge(5, 7, 10)
	g.addEdge(5, 6, 15)
	g.addEdge(6, 2, 6)
	g.addEdge(6, 7, 10)
	g.Maxflow(0, 7)
}
