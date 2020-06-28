package avltree

func findMin(node *Node) *Node {
	for !node.isNil() && !node.Left.isNil() {
		node = node.Left
	}

	return node
}

func findMax(node *Node) *Node {
	for !node.isNil() && !node.Right.isNil() {
		node = node.Right
	}

	return node
}

func inOrderTree(x *Node) []int {
	if x.isNil() {
		return []int{}
	}
	left := inOrderTree(x.Left)
	right := inOrderTree(x.Right)

	return append(left, append([]int{x.Key}, right...)...)
}

func findNode(tree *AVLTree, key int) *Node {
	node := tree.Root

	for !node.isNil() {
		if key == node.Key {
			return node
		}
		if key < node.Key {
			node = node.Left
			continue
		}
		node = node.Right
	}

	return nil
}

func leftRotate(tree *AVLTree, node *Node) {
	if node.Right.isNil() || !tree.Root.P.isNil() {
		return
	}
	rotated := node.Right
	node.Right = rotated.Left

	if !rotated.Left.isNil() {
		rotated.Left.P = node
	}
	rotated.P = node.P

	if node.P.isNil() {
		tree.Root = rotated
	} else if node == node.P.Left {
		node.P.Left = rotated
	} else {
		node.P.Right = rotated
	}
	rotated.Left = node
	node.P = rotated
	rotated.fixHeight()
	node.fixHeight()
}

func rightRotate(tree *AVLTree, node *Node) {
	if node.Left.isNil() || !tree.Root.P.isNil() {
		return
	}
	rotated := node.Left
	node.Left = rotated.Right

	if !rotated.Right.isNil() {
		rotated.Right.P = node
	}
	rotated.P = node.P

	if node.P.isNil() {
		tree.Root = rotated
	} else if node == node.P.Right {
		node.P.Right = rotated
	} else {
		node.P.Left = rotated
	}
	rotated.Right = node
	node.P = rotated
	rotated.fixHeight()
	node.fixHeight()
}

func balance(tree *AVLTree, node *Node) {
	for !node.isNil() {
		node.fixHeight()

		if node.balanceFactor() > 1 {
			if node.Right.balanceFactor() < 0 {
				rightRotate(tree, node.Right)
			}
			leftRotate(tree, node)
		}

		if node.balanceFactor() < -1 {
			if node.Left.balanceFactor() > 0 {
				leftRotate(tree, node.Left)
			}
			rightRotate(tree, node)
		}

		node = node.P
	}
}

func transplant(tree *AVLTree, old *Node, new *Node) {
	switch {
	case old.P.isNil():
		tree.Root = new
	case old == old.P.Left:
		old.P.Left = new
	case old == old.P.Right:
		old.P.Right = new
	}
}

func deleteKey(tree *AVLTree, node *Node) {
	var next *Node

	if node.Right.isNil() && node.Left.isNil() {
		transplant(tree, node, nil)
		balance(tree, node.P)
		return
	}

	if node.Right.isNil() && !node.Left.isNil() {
		transplant(tree, node, node.Left)
		balance(tree, node.Left)
		return
	}

	next = findMin(node.Right)

	if next == next.P.Left {
		next.P.Left = nil
	}
	if next == next.P.Right {
		next.P.Right = nil
	}

	transplant(tree, node, next)
	next.Left = node.Left
	next.Right = node.Right

	if !next.Left.isNil() {
		next.Left.P = next
	}
	if !next.Right.isNil() {
		next.Right.P = next
	}
	next.P = node.P
	balance(tree, next.Right)

}
