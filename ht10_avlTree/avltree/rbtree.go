package avltree

// AVLTree type
type AVLTree struct {
	Root *Node
}

// MakeTree func
func MakeTree() *AVLTree {
	return &AVLTree{
		Root: nil,
	}
}

// Min func
func (tree *AVLTree) Min() int {
	if tree.Root.isNil() {
		return 0
	}
	node := findMin(tree.Root)

	return node.Key
}

// Max func
func (tree *AVLTree) Max() int {
	if tree.Root.isNil() {
		return 0
	}
	node := findMax(tree.Root)

	return node.Key
}

// Insert func
func (tree *AVLTree) Insert(key int) {
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
	case parent.isNil():
		tree.Root = new
	case new.Key < parent.Key:
		parent.Left = new
	default:
		parent.Right = new
	}

	balance(tree, parent)
}

// InOrder func
func (tree *AVLTree) InOrder() []int {
	return inOrderTree(tree.Root)
}

// Exists func
func (tree *AVLTree) Exists(key int) bool {
	node := findNode(tree, key)

	return !node.isNil()
}

// Delete func
func (tree *AVLTree) Delete(key int) bool {
	node := findNode(tree, key)

	if !node.isNil() {
		deleteKey(tree, node)
		return true
	}

	return false
}
