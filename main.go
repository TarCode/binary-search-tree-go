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

func (n *node) levelOrder(visit func(int)) {
	if (n == nil) {
		return 
	}
	for queue := []*node{n}; ; {
		n = queue[0]
		visit(n.value)
		copy(queue, queue[1:])
		queue = queue[:len(queue)-1]

		if (n.left != nil) {
			queue = append(queue, n.right)
		}

		if (n.right != nil) {
			queue = append(queue, n.right)
		}

		if len(queue) == 0 {
			return 
		}
	}
}

func main() {
	
	tree := &node { 1, 
		&node { 2, 
			&node { 4, 
				&node { 7, 
					nil, nil }, nil}, 
					&node { 5, nil, nil } },
			&node { 3,
				 &node { 6,
					&node{ 8, nil, nil },
					&node { 9, nil, nil } },
			nil}}


	fmt.Print("PREORDER: ")
	tree.preOrder(visitor)
}

func visitor(value int) {
	fmt.Print(value, " ")
}