package main

import (
	"cs-go-course5/dzhebov/ht9_rbTree/rbtree"
	"fmt"
)

// Реализуйте красно-черное дерево, соответствующее следующему интерфейсу
// type RBTree interface {
// 	Min() int
// 	Max() int
// 	InOrder() []int
// 	Insert(key int)
// 	Exists(key int) bool
// 	Delete(key int) bool
// }

func main() {
	tree := rbtree.MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(17)

	fmt.Printf("\ntree: %v", tree.InOrder())
	tree.Delete(11)
	fmt.Printf("\ntree after deleting: %v", tree.InOrder())
	min := tree.Min()
	max := tree.Max()
	fmt.Printf("\nMin value: %v\nMax value: %v", min, max)
	fmt.Printf("\nExist %v: %v", 11, tree.Exists(11))
	fmt.Printf("\nExist %v: %v", 7, tree.Exists(7))
}
