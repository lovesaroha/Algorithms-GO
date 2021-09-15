// Dynamic connectivity - quick find.
package main

import "fmt"

// Quick find object defined.
type quickFindObject struct {
	ids [10]int
}

// Union function connect two nodes in ids.
func (qf *quickFindObject) union(a int, b int) {
	p := qf.ids[a]
	q := qf.ids[b]
	// Example union(4,3) will result in ids[i] = ids[3] where ids[i] == ids[4].
	for i := 0; i < len(qf.ids); i++ {
		if qf.ids[i] == p {
			qf.ids[i] = q
		}
	}
}

// Connected function checks if two nodes are connected (are in same connected components).
func (qf quickFindObject) connected(a int, b int) bool {
	return qf.ids[a] == qf.ids[b]
}

// This function create new quick find object and initialize values.
func newQuickFind() quickFindObject {
	qf := quickFindObject{}
	for i := 0; i < len(qf.ids); i++ {
		qf.ids[i] = i
	}
	return qf
}

func main() {
	// Initialize quick find object.
	qf := newQuickFind()
	// Perform some unions.
	qf.union(1, 3)
	qf.union(5, 6)
	qf.union(8, 9)
	qf.union(3, 6)

	// Check connections.
	fmt.Println("1 and 6 connected: ", qf.connected(1, 6))
	fmt.Println("9 and 6 connected: ", qf.connected(9, 6))
}
