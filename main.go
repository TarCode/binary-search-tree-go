package main

import "fmt"

type node struct {
	value			int
	left, right		*node
}

func (n *node) preOrder(visit func(int)) {
	if (n == nil) {
		return 
	}
	visit(n.value)
	n.left.preOrder(visit)
	n.right.preOrder(visit)
}
