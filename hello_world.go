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

func removeNode(key string, lrucache *LRUCache) {
	if lrucache.cache[key] != nil {
		node := lrucache.cache[key]
		delete(lrucache.cache, key)
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}

func addToFront(key string, value int, lrucache *LRUCache) {
	node := &Node{key: key, value: value}
	head := lrucache.head
	temp := head.next
	head.next = node
	node.prev = head
	node.next = temp
	temp.prev = node

}

func (c *LRUCache) Get(key string) (int, bool) {
	if c.cache[key] == nil {
		//cache miss
		return 0, false
	}
	value := c.cache[key].value
	removeNode(key, c)
	addToFront(key, value, c)
	return c.cache[key].value, true
}

func main() {
	m := make([]int, 5)
	fmt.Println(m)
}
