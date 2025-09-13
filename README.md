# Binary Tree Traversal in Go

A simple Go implementation demonstrating the four main binary tree traversal algorithms.

## Overview

This program implements a basic binary tree with methods for all standard traversal patterns:
- **Pre-order**: Root → Left → Right
- **In-order**: Left → Root → Right  
- **Post-order**: Left → Right → Root
- **Level-order**: Breadth-first traversal

## Usage

```bash
go run main.go
```

## Output

```
PREORDER: 1 2 4 7 5 3 6 8 9 
INORDER: 7 4 2 5 1 8 6 9 3 
POSTORDER: 7 4 5 2 8 9 6 3 1 
LEVEL ORDER: 1 2 3 4 5 6 7 8 9 
```

## Tree Structure

The example tree:
```
        1
      /   \
     2     3
   /  \   /
  4    5 6
 /      / \
7      8   9
```

## Methods

- `preOrder(visit func(int))` - Visits root, then left subtree, then right subtree
- `inOrder(visit func(int))` - Visits left subtree, then root, then right subtree  
- `postOrder(visit func(int))` - Visits left subtree, then right subtree, then root
- `levelOrder(visit func(int))` - Visits nodes level by level from top to bottom

Each method accepts a visitor function to process node values during traversal.
