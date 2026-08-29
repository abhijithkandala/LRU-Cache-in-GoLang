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

func main() {
	m := make([]int, 5)
	fmt.Println(m)
}
