package main

import "fmt"
import "time"

type Node struct {
	key      string
	value    int
	expireAt time.Time

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

func (lrucache *LRUCache) addToFront(node *Node) {

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
	if time.Now().After(node.expireAt) {
		delete(c.cache, key)
		removeNode(node)
		return 0, false
	}
	removeNode(node)
	c.addToFront(c.cache[key])
	return c.cache[key].value, true
}

func (c *LRUCache) Put(key string, value int, ttl time.Duration) bool {
	if c.cache[key] != nil {
		node := c.cache[key]
		node.expireAt = time.Now().Add(ttl)
		node.value = value
		removeNode(node)
		c.addToFront(node)
	} else {
		if len(c.cache) == c.capacity {
			delete(c.cache, c.tail.prev.key)
			removeNode(c.tail.prev)
		}
		node := &Node{key: key, value: value, expireAt: time.Now().Add(ttl)}
		c.cache[key] = node
		c.addToFront(node)
	}
	return true
}

func main() {
	m := make([]int, 5)
	fmt.Println(m)
}
