// Ternary search tries.
package main

import "fmt"

// Queue object defined.
type queueObject struct {
	first *qnodeObject
	last  *qnodeObject
}

// Node object structure.
type qnodeObject struct {
	value int
	next  *qnodeObject
}

// Enqueue function add new value in queue.
func (queue *queueObject) enqueue(value int) {
	var newNode = qnodeObject{value: value}
	if queue.last != nil {
		// Last node next pointer now points to new node.
		queue.last.next = &newNode
	}
	// Last pointer points to new node.
	queue.last = &newNode
	if queue.first == nil {
		queue.first = &newNode
	}
}

// Dequeue function remove value from queue and return the value.
func (queue *queueObject) dequeue() int {
	var value int
	if queue.first == queue.last {
		// Queue have single value.
		queue.last = nil
	}
	value = queue.first.value
	// First points to second node.
	queue.first = queue.first.next
	return value
}

// Trie object.
type trieObject struct {
	root *nodeObject
}

// Node object.
type nodeObject struct {
	charCode int
	index    int
	left     *nodeObject
	right    *nodeObject
	medium   *nodeObject
}

// This function insert new value in trie.
func (trie *trieObject) insert(word string, index int) {
	trie.root = put(trie.root, word, index, 0)
}

// This function put value in trie.
func put(node *nodeObject, word string, index int, digit int) *nodeObject {
	var code = charCode(string(word[digit]))
	if node == nil {
		node = &nodeObject{charCode: code}
	}
	if code > node.charCode {
		// Go right.
		node.right = put(node.right, word, index, digit)
	} else if code < node.charCode {
		// Go left.
		node.left = put(node.left, word, index, digit)
	} else if digit < len(word)-1 {
		// Go medium.
		node.medium = put(node.medium, word, index, digit+1)
	} else {
		// Save index.
		node.index = index
	}
	return node
}

// This function return inxed of given word.
func (trie trieObject) get(word string) int {
	node := find(trie.root, word, 0)
	if node != nil {
		return node.index
	}
	return -1
}

// This function find given word index.
func find(node *nodeObject, word string, digit int) *nodeObject {
	var code = charCode(string(word[digit]))
	if node == nil {
		return nil
	}
	if code > node.charCode {
		// Go right.
		return find(node.right, word, digit)
	} else if code < node.charCode {
		// Go left.
		return find(node.left, word, digit)
	} else if digit < len(word)-1 {
		// Go medium.
		return find(node.medium, word, digit+1)
	} else {
		return node
	}
}

// This function shows all words in a trie.
func (trie trieObject) words() {

}

// Inorder traversal of trie.
func inOrder(node *nodeObject) {
	if node == nil {
		return
	}
	if node.index != 0 {

	}
	inOrder(node.left)

}

// This function return character code.
func charCode(char string) int {
	return int([]rune(char)[0])
}

func main() {
	// Create a trie.
	t := trieObject{}
	// Add some words in trie.
	t.insert("love", 1)
	t.insert("saroha", 2)
	t.insert("kush", 3)
	t.insert("lucky", 4)
	fmt.Println(t.get("lucky"))
}
