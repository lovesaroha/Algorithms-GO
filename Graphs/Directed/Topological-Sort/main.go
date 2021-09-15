// Topological sort in directed acyclic graph.
package main

import "fmt"

// Stack object defined.
type stackObject struct {
	first *nodeObject
}

// Push function to add value in stack.
func (stack *stackObject) push(value int) {
	var newNode = nodeObject{next: stack.first, value: value}
	stack.first = &newNode
}

// This fucntion shows values of stack.
func (stack stackObject) show() {
	var node = stack.first
	if node == nil {
		// Empty stack.
		return
	}
	for {
		fmt.Println(node.value)
		if node.next == nil {
			return
		}
		node = node.next
	}
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

// This function perform depth first search from a given source vertex and change visited vertices and path.
func (graph graphObject) depthFirstSearch(source int, vertex int, visited *[10]bool, edgeTo *[10]int, order *stackObject, isCyclic *bool) {
	if *isCyclic {
		// Check if graph is cyclic.
		return
	}
	// Mark the vertex visited.
	visited[vertex] = true
	var currentNode = graph.vertices[vertex]
	// Find all adjacent vertices.
	for {
		if currentNode == nil {
			// No more adjacent node.
			break
		}
		if !visited[currentNode.value] {
			// Save the edge.
			edgeTo[currentNode.value] = vertex
			// Given adjacent node is not visited.
			graph.depthFirstSearch(source, currentNode.value, visited, edgeTo, order, isCyclic)
		} else if source == currentNode.value {
			// Adjacent node is source so graph is cyclic ( There is a directed path from source to source ).
			*isCyclic = true
			return
		}
		currentNode = currentNode.next
	}
	order.push(vertex)
}

// This function perform topological sort.
func (graph graphObject) topologicalSort() {
	var visited [10]bool
	var edgeTo [10]int
	var order = stackObject{}
	var isCyclic bool
	// For every unvisited vertices perform depth first search.
	for i := 0; i < 10; i++ {
		if visited[i] {
			// Vertex is already visited.
			continue
		}
		graph.depthFirstSearch(i, i, &visited, &edgeTo, &order, &isCyclic)
		if isCyclic {
			// Graph is cyclic.
			fmt.Println("Graph is cyclic")
			return
		}
	}
	order.show()
}

func main() {

	// Create a acyclic directed graph.
	g := graphObject{}
	g.addEdge(0, 2)
	g.addEdge(0, 5)
	g.addEdge(0, 1)
	g.addEdge(3, 2)
	g.addEdge(3, 6)
	g.addEdge(3, 4)
	g.addEdge(3, 5)
	g.addEdge(1, 4)
	g.addEdge(6, 0)
	g.addEdge(6, 4)
	g.topologicalSort()
}
