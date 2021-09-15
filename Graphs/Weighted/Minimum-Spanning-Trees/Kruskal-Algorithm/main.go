// Minimum spanning tree using kruskal algorithm on weighted graph.
package main

import "fmt"

// Quick union object defined.
type quickUnionObject struct {
	sizes [10]int
	ids   [10]int
}

// This function connect two nodes in ids.
func (qu *quickUnionObject) union(a int, b int) {
	p := qu.root(a)
	q := qu.root(b)
	if p == q {
		return
	}
	sizeOfP := qu.sizes[p]
	sizeOfQ := qu.sizes[q]
	if sizeOfP > sizeOfQ {
		// P is a bigger tree.
		qu.ids[q] = p
		qu.sizes[p] += qu.sizes[q]
	} else {
		qu.ids[p] = q
		qu.sizes[q] += qu.sizes[p]
	}
}

// This function find root of given element.
func (qu quickUnionObject) root(a int) int {
	for {
		p := qu.ids[a]
		if p == a {
			return a
		}
		// Path compression.
		qu.ids[a] = qu.ids[p]
		a = p
	}
}

// Connected function checks if two nodes are connected (are in same connected components).
func (qu quickUnionObject) isConnected(a int, b int) bool {
	p := qu.root(a)
	q := qu.root(b)
	return p == q
}

// This function prepare union object values.
func (qu *quickUnionObject) prepare() {
	for i := 0; i < 10; i++ {
		qu.ids[i] = i
		qu.sizes[i] = 1
	}
}

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
	sortedEdges queueObject
}

// Edge object.
type edgeObject struct {
	vertexA int
	vertexB int
	weight  float64
	next    *edgeObject
}

// This function add edge in graph between given vertex.
func (graph *graphObject) addEdge(vertexA int, vertexB int, weight float64) {
	// Save edges in sorted order.
	graph.sortedEdges.enqueue(edgeObject{vertexA: vertexA, vertexB: vertexB, weight: weight})
}

// This function find minimum spanning tree using kruskal.
func (graph graphObject) minimumSpanningtree() {
	var mst = queueObject{}
	var sortedEdges = graph.sortedEdges
	// Create quick union object.
	var unionFind = quickUnionObject{}
	unionFind.prepare()
	for {
		if sortedEdges.isEmpty() {
			// No more edges in queue.
			mst.show()
			return
		}
		// Get edge from sorted edges.
		var edge = sortedEdges.dequeue()
		if !unionFind.isConnected(edge.vertexA, edge.vertexB) {
			// Save two vertices in union and add to mst.
			unionFind.union(edge.vertexA, edge.vertexB)
			mst.enqueue(edge)
		}
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
