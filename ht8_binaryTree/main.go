// Реализуйте бинарное дерево, соответствующее следующему интерфейсу
// type BinTree interface {
// 	Min() int
// 	Max() int
// 	InOrder() []int
// 	Insert(key int)
// 	Exists(key int) bool
// 	Delete(key int) bool
// }
package main

import (
	"cs-go-course5/dzhebov/ht8_binaryTree/btree"
	"fmt"
)

func main() {
	fmt.Println("binary tree")

	tree := btree.NewTree(11)
	tree.Insert(10)
	min := tree.Min()
	fmt.Println(*tree.R)
	fmt.Printf("Min tree value: %v", min)
}
