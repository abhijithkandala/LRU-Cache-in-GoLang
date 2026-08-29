package main

import "fmt"

type Node struct {
	key   string
	value int

	next *Node
	prev *Node
}

func main() {
	m := make([]int, 5)
	fmt.Println(m)
}
