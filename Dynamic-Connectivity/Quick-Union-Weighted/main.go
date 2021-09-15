// Dynamic connectivity - quick union weighted.
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
func (qu quickUnionObject) connected(a int, b int) bool {
	p := qu.root(a)
	q := qu.root(b)
	return p == q
}

// This function create new quick union object and initialize values.
func newQuickUnion() quickUnionObject {
	qu := quickUnionObject{}
	for i := 0; i < len(qu.ids); i++ {
		qu.ids[i] = i
		qu.sizes[i] = 1
	}
	return qu
}

func main() {
	// Initialize quick union object.
	qu := newQuickUnion()
	// Perform some unions.
	qu.union(1, 3)
	qu.union(5, 6)
	qu.union(8, 9)
	qu.union(1, 6)

	// Check connections.
	fmt.Println("1 and 6 connected: ", qu.connected(1, 6))
	fmt.Println("9 and 6 connected: ", qu.connected(9, 6))
}
