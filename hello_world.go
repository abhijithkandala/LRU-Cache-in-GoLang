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

func main() {
	m := make([]int, 5)
	fmt.Println(m)
}
