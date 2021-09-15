// R-Way tries.
package main

import "fmt"

var chars = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// Queue object defined.
type queueObject struct {
	first *qnodeObject
	last  *qnodeObject
}

// Node object structure.
type qnodeObject struct {
	value string
	next  *qnodeObject
}

// Enqueue function add new value in queue.
func (queue *queueObject) enqueue(value string) {
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

// This function shows values of queue.
func (queue queueObject) show() {
	if queue.first == nil {
		// Empty queue.
		return
	}
	var node = queue.first
	for {
		fmt.Println(node.value)
		if node.next == nil {
			return
		}
		node = node.next
	}
}

// Trie object.
type trieObject struct {
	root *nodeObject
}

// Node object.
type nodeObject struct {
	value int
	next  [26]*nodeObject
}

// This function insert value in a trie.
func (trie *trieObject) insert(word string, value int) {
	trie.root = put(trie.root, word, value, 0)
}

// This function put value in trie.
func put(node *nodeObject, word string, value int, digit int) *nodeObject {
	if node == nil {
		node = &nodeObject{}
	}
	if digit == len(word) {
		// End of word save value.
		node.value = value
		return node
	}
	var code = charCode(string(word[digit]))
	// Create a node for next digit.
	node.next[code] = put(node.next[code], word, value, digit+1)
	return node
}

// This function return value of given word in trie.
func (trie trieObject) get(word string) int {
	node := find(trie.root, word, 0)
	if node != nil {
		return node.value
	}
	return -1
}

// This function find given value.
func find(node *nodeObject, word string, digit int) *nodeObject {
	if node == nil {
		return nil
	}
	if len(word) == digit {
		// Last char found.
		return node
	}
	var code = charCode(string(word[digit]))
	return find(node.next[code], word, digit+1)
}

// This function remove word from trie.
func (trie *trieObject) remove(word string) {
	trie.root = delete(trie.root, word, 0)
}

// This function delete given word.
func delete(node *nodeObject, word string, digit int) *nodeObject {
	if node == nil {
		return nil
	}
	var isNil = true
	// Check all link.
	for i := 0; i < 26; i++ {
		if node.next[i] != nil {
			// More links.
			isNil = false
		}
	}
	if isNil {
		// No more links.
		return nil
	}
	if len(word) == digit {
		// Remove value.
		node.value = 0
		return node
	}
	var code = charCode(string(word[digit]))
	node.next[code] = delete(node.next[code], word, digit+1)
	return node
}

// This function show all words in trie.
func (trie trieObject) words() queueObject {
	queue := queueObject{}
	collect(trie.root, "", &queue)
	return queue
}

// This function collect given char.
func collect(node *nodeObject, prefix string, queue *queueObject) {
	if node == nil {
		return
	}
	if node.value > 0 {
		queue.enqueue(prefix)
	}
	for i := 0; i < 26; i++ {
		collect(node.next[i], prefix+chars[i], queue)
	}
}

// This function show all words matching prefix.
func (trie trieObject) wordsWithPrefix(prefix string) queueObject {
	queue := queueObject{}
	node := find(trie.root, prefix, 0)
	if node == nil {
		// No prefix.
		return queue
	}
	collect(node, prefix, &queue)
	return queue
}

// This function return character code.
func charCode(char string) int {
	// Example for a return 0.
	return int([]rune(char)[0]) - 97
}

func main() {
	// Create a trie.
	t := trieObject{}
	// Add some words in trie.
	t.insert("love", 1)
	t.insert("saroha", 2)
	t.insert("kush", 3)
	t.insert("lucky", 4)
	fmt.Println(t.get("kush"))
	t.remove("saroha")
	t.words().show()
	t.wordsWithPrefix("l").show()
}
