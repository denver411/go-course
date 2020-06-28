package main

import (
	"cs-go-course5/dzhebov/ht8_binaryTree/bintree"
	"fmt"
)

// Реализуйте бинарное дерево, соответствующее следующему интерфейсу
// type BinTree interface {
// 	Min() int
// 	Max() int
// 	InOrder() []int
// 	Insert(key int)
// 	Exists(key int) bool
// 	Delete(key int) bool
// }

func main() {
	tree := bintree.MakeTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(15)
	tree.Insert(13)
	tree.Insert(17)

	min := tree.Min()
	max := tree.Max()
	arr := tree.InOrder()
	fmt.Printf("\nRoot: %v", tree.Root)
	fmt.Printf("\nCustom Node: %v", tree.Root.Left)
	fmt.Printf("\nInOrder tree: %v", arr)
	fmt.Printf("\nMin value: %v\nMax value: %v", min, max)
	fmt.Printf("\nExist %v: %v", 7, tree.Exists(7))
	fmt.Printf("\nExist %v: %v", 23, tree.Exists(23))

	tree.Delete(3)
	fmt.Printf("\nInOrder after deleting 3: %v", tree.InOrder())
}
