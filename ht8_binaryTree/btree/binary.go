package btree

// BinTree ...
type BinTree interface {
	Min() int
	Max() int
	InOrder() []int
	Insert(key int)
	Exists(key int) bool
	Delete(key int) bool
}

// Node ...
type Node struct {
	P     *Node
	Left  *Node
	Right *Node
	Key   int
}

//BinaryTree ...
type BinaryTree struct {
	R *Node
}

func makeNode(key int) *Node {
	return &Node{
		P:     nil,
		Left:  nil,
		Right: nil,
		Key:   key,
	}
}

func NewTree(key int) *BinaryTree {
	return &BinaryTree{
		R: makeNode(key),
	}
}

// Min ...
func (tree *BinaryTree) Min() int {
	if tree == nil || tree.R == nil {
		return 0
	}

	node := tree.R
	for node.Left != nil {
		node = node.Left
	}

	return node.Key
}

// Max ...
func (tree *BinaryTree) Max() int {
	if tree == nil || tree.R == nil {
		return 0
	}

	node := tree.R
	for node.Right != nil {
		node = node.Right
	}

	return node.Key
}

// Insert ...
func (tree *BinaryTree) Insert(key int) {
	node := tree.R
	new := makeNode(key)
	if node == nil {
		tree.R = new
	}

	var parent Node

	for node != nil {
		parent = *node
		if new.Key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
		new.P = &parent
	}

}

// InOrder ...
// func (tree *BinaryTree) InOrder() []int {
// 	// var node = tree.r
// 	return []int{1, 2}

// }
