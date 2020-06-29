package main

import (
	"cs-go-course5/dzhebov/ht11_b_tree/btree"
	"fmt"
)

/* Реализуйте B-tree, соответствующее тому же интерфейсу, что и бинарное дерево поиска.
Допускается реализация в оперативной памяти без сброса на жесткий диск */

// type BinTree interface {
// 	Min() int
// 	Max() int
// 	InOrder() []int
// 	Insert(key int)
// 	Exists(key int) bool
// 	Delete(key int) bool
// }

func main() {
	tree := btree.MakeTree(3)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(17)
	tree.Insert(18)
	tree.Insert(19)
	tree.Insert(20)
	tree.Insert(21)
	tree.Insert(22)
	tree.Insert(23)
	tree.Insert(24)

	fmt.Printf("\ntree: %v", tree.InOrder())
	tree.Delete(24)
	fmt.Printf("\ntree after deleting: %v", tree.InOrder())
	min := tree.Min()
	max := tree.Max()
	fmt.Printf("\nMin value: %v\nMax value: %v", min, max)
	fmt.Printf("\nExist %v: %v", 111, tree.Exists(111))
	fmt.Printf("\nExist %v: %v", 7, tree.Exists(7))
}
