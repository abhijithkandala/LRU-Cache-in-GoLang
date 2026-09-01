package main

import "fmt"

type Node struct {
	key   string
	value int

	next *Node
	prev *Node
}

type LRUCache struct {
	capacity int
	cache    map[string]*Node

	head *Node
	tail *Node
}

func newLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func removeNode(node *Node) {

	node.prev.next = node.next
	node.next.prev = node.prev

}

func addToFront(lrucache *LRUCache, node *Node) {

	head := lrucache.head
	temp := head.next
	head.next = node
	node.prev = head
	node.next = temp
	temp.prev = node

}

func (c *LRUCache) Get(key string) (int, bool) {
	node, ok := c.cache[key]
	if !ok {
		return 0, false
	}

	removeNode(node)
	addToFront(c, c.cache[key])
	return c.cache[key].value, true
}

func main() {
	m := make([]int, 5)
	fmt.Println(m)
}
