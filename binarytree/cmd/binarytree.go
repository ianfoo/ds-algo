package main

import (
	"fmt"

	bt "github.com/ianfoo/ds-algo/binarytree"
)

func main() {
	tree := bt.NewTree()
	tree.Insert(7)
	tree.Insert(15)
	tree.Insert(17)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(29)
	tree.Insert(-3)

	fmt.Println("InOrder:  ", tree.InOrder())
	fmt.Println("PreOrder: ", tree.PreOrder())
	fmt.Println("PostOrder:", tree.PostOrder())
}
