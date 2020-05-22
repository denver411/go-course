package bintree

// Node type
type Node struct {
	P     *Node
	Left  *Node
	Right *Node
	Key   int
}

// BinaryTree type
type BinaryTree struct {
	Root *Node
}

// MakeTree func
func MakeTree() *BinaryTree {
	return &BinaryTree{
		Root: nil,
	}
}

// Min func
func (tree *BinaryTree) Min() int {
	if tree.Root == nil {
		return 0
	}
	node := findMin(tree.Root)

	return node.Key
}

// Max func
func (tree *BinaryTree) Max() int {
	if tree.Root == nil {
		return 0
	}
	node := findMax(tree.Root)

	return node.Key
}

// Insert func
func (tree *BinaryTree) Insert(key int) {
	var parent *Node
	node := tree.Root
	new := makeNode(key)

	for !node.isNil() {
		parent = node
		if new.Key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
		new.P = parent
	}

	switch {
	case parent == nil:
		tree.Root = new
	case new.Key < parent.Key:
		parent.Left = new
	case new.Key >= parent.Key:
		parent.Right = new
	}
}

// InOrder func
func (tree *BinaryTree) InOrder() []int {
	return inOrderTree(tree.Root)
}

// Exists func
func (tree *BinaryTree) Exists(key int) bool {
	node := findNode(tree.Root, key)

	return node != nil
}

// Delete func
func (tree *BinaryTree) Delete(key int) bool {
	node := findNode(tree.Root, key)

	if !node.isNil() {
		deleteKey(tree, node)
		return true
	}

	return false
}
