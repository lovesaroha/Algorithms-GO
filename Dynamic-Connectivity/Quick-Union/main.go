// Dynamic connectivity - quick union.
package main

import "fmt"

// Quick union object defined.
type quickUnionObject struct {
	ids [10]int
}

// This function find root of given element.
func (qu quickUnionObject) root(a int) int {
	for {
		p := qu.ids[a]
		if p == a {
			return a
		}
		a = qu.ids[a]
	}
}

// This function connect two nodes in ids.
func (qu *quickUnionObject) union(a int, b int) {
	p := qu.root(a)
	q := qu.root(b)
	if p == q {
		return
	}
	qu.ids[p] = q
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
