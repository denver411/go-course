package rbtree

// RBTree type
type RBTree struct {
	Root *Node
}

// MakeTree func
func MakeTree() *RBTree {
	return &RBTree{
		Root: nil,
	}
}

// Min func
func (tree *RBTree) Min() int {
	if tree.Root.isNil() {
		return 0
	}
	node := findMin(tree.Root)

	return *node.Key
}

// Max func
func (tree *RBTree) Max() int {
	if tree.Root.isNil() {
		return 0
	}
	node := findMax(tree.Root)

	return *node.Key
}

// Insert func
func (tree *RBTree) Insert(key int) {
	parent := makeNil()
	node := tree.Root
	new := makeNode(&key)

	for !node.isNil() {
		parent = node
		if *new.Key < *node.Key {
			node = node.Left
		} else {
			node = node.Right
		}

		new.P = parent
	}

	switch {
	case parent.isNil():
		tree.Root = new
	case *new.Key < *parent.Key:
		parent.Left = new
	default:
		parent.Right = new
	}

	insertFixup(tree, new)
}

// InOrder func
func (tree *RBTree) InOrder() []int {
	return inOrderTree(tree.Root)
}

// Exists func
func (tree *RBTree) Exists(key int) bool {
	node := findNode(tree, &key)

	return !node.isNil()
}

// Delete func
func (tree *RBTree) Delete(key int) bool {
	node := findNode(tree, &key)

	if !node.isNil() {
		deleteKey(tree, node)
		return true
	}

	return false
}
