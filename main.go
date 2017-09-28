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

func (n *node) inOrder(visit func(int)) {
	if (n == nil) {
		return 
	} 
	n.left.inOrder(visit)
	visit(n.value)
	n.right.inOrder(visit)
}

func (n *node) postOrder(visit func(int)) {
	if (n == nil) {
		return 
	}
	n.left.postOrder(visit)
	n.right.postOrder(visit)
	visit(n.value)
}