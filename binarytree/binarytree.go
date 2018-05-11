package binarytree

import "strconv"

// Node is a single element in a tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Tree is a binary tree.
type Tree struct {
	*Node
}

// NewTree creates a new Tree.
func NewTree() *Tree {
	return &Tree{}
}

// Insert a value into a Tree.
func (t *Tree) Insert(v int) {
	if t.Node == nil {
		t.Node = newNode(v)
		return
	}
	t.Node = t.Node.insert(v)
}

func newNode(v int) *Node {
	return &Node{
		Value: v,
	}
}

func (n *Node) insert(v int) *Node {
	if n == nil {
		return newNode(v)
	}
	if v < n.Value {
		n.Left = n.Left.insert(v)
		return n
	}
	n.Right = n.Right.insert(v)
	return n
}

// InOrder returns a string of the binary tree in order, where the current
// node's value is appended between the left and right values.
func (n *Node) InOrder() string {
	s := n.inOrder()
	return s[:len(s)-1]
}

func (n *Node) inOrder() string {
	var s string
	if n.Left != nil {
		s += n.Left.inOrder()
	}
	s += strconv.Itoa(n.Value) + " "
	if n.Right != nil {
		s += n.Right.inOrder()
	}
	return s
}

// PreOrder returns a string of the binary tree as traversed in "preorder,"
// where the current node's value is appended first.
func (n *Node) PreOrder() string {
	s := n.preOrder()
	return s[:len(s)-1]
}

func (n *Node) preOrder() string {
	s := strconv.Itoa(n.Value) + " "
	if n.Left != nil {
		s += n.Left.preOrder()
	}
	if n.Right != nil {
		s += n.Right.preOrder()
	}
	return s
}

// PostOrder returns a string of the binary tree as traversed in "postorder,"
// where the current node's value is appended last.
func (n *Node) PostOrder() string {
	s := n.postOrder()
	return s[:len(s)-1]
}

func (n *Node) postOrder() string {
	var s string
	if n.Left != nil {
		s += n.Left.postOrder()
	}
	if n.Right != nil {
		s += n.Right.postOrder()
	}
	s += strconv.Itoa(n.Value) + " "
	return s
}
