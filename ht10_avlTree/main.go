package main

import (
	"cs-go-course5/dzhebov/ht10_avlTree/avltree"
	"fmt"
)

// Реализуйте AVL дерево, соответствующее следующему интерфейсу
// type Avltree interface {
// 	Min() int
// 	Max() int
// 	InOrder() []int
// 	Insert(key int)
// 	Exists(key int) bool
// 	Delete(key int) bool
// }

func main() {
	tree := avltree.MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
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
	tree.Delete(10)
	tree.Delete(17)
	fmt.Printf("\ntree after deleting: %v", tree.InOrder())
	min := tree.Min()
	max := tree.Max()
	fmt.Printf("\nMin value: %v\nMax value: %v", min, max)
	fmt.Printf("\nExist %v: %v", 11, tree.Exists(11))
	fmt.Printf("\nExist %v: %v", 7, tree.Exists(7))
}
