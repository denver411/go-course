package rbtree

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

	return append(left, append([]int{*x.Key}, right...)...)
}

func findNode(tree *RBTree, key *int) *Node {
	node := tree.Root

	for !node.isNil() {
		if *key == *node.Key {
			return node
		}
		if *key < *node.Key {
			node = node.Left
			continue
		}
		node = node.Right
	}

	return nil
}

func leftRotate(tree *RBTree, node *Node) {
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
}

func rightRotate(tree *RBTree, node *Node) {
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
}

func insertFixup(tree *RBTree, node *Node) {
	// if node == tree.Root {
	// 	tree.Root.makeBlack()
	// 	return
	// }

	for node.P.isRed() {
		if node.P == node.P.P.Left {
			uncle := node.P.P.Right

			if uncle.isRed() {
				node.P.makeBlack()
				uncle.makeBlack()
				node.P.P.makeRed()
				node = node.P.P
			} else {
				if node == node.P.Right {
					node = node.P
					leftRotate(tree, node)
				}
				node.P.makeBlack()
				node.P.P.makeRed()
				rightRotate(tree, node.P.P)
			}
		} else {
			uncle := node.P.P.Left

			if uncle.isRed() {
				node.P.makeBlack()
				uncle.makeBlack()
				node.P.P.makeRed()
				node = node.P.P
			} else {
				if node == node.P.Left {
					node = node.P
					rightRotate(tree, node)
				}
				node.P.makeBlack()
				node.P.P.makeRed()
				leftRotate(tree, node.P.P)
			}
		}
	}

	tree.Root.makeBlack()
}

func transplant(tree *RBTree, old *Node, new *Node) {
	switch {
	case old.P.isNil():
		tree.Root = new
	case old == old.P.Left:
		old.P.Left = new
	case old == old.P.Right:
		old.P.Right = new

	}

	new.P = old.P
}

func deleteFixup(tree *RBTree, node *Node) {
	for node != tree.Root && node.isBlack() {
		if node == node.P.Left {
			bro := node.P.Right

			if bro.isRed() {
				bro.makeBlack()
				node.P.makeRed()
				leftRotate(tree, node.P)
				bro = node.P.Right
			}

			if bro.Left.isBlack() && bro.Right.isBlack() {
				bro.makeRed()
				node = node.P
			} else {
				if bro.Right.isBlack() {
					bro.Left.makeBlack()
					bro.makeRed()
					rightRotate(tree, bro)
					bro = node.P.Right
				}

				bro.Color = node.P.Color
				node.P.makeBlack()
				bro.Right.makeBlack()
				leftRotate(tree, node.P)
				node = tree.Root
			}
		} else {
			bro := node.P.Left

			if bro.isRed() {
				bro.makeBlack()
				node.P.makeRed()
				rightRotate(tree, node.P)
				bro = node.P.Left
			}

			if bro.Left.isBlack() && bro.Right.isBlack() {
				bro.makeRed()
				node = node.P
			} else {
				if bro.Left.isBlack() {
					bro.Right.makeBlack()
					bro.makeRed()
					leftRotate(tree, bro)
					bro = node.P.Left
				}

				bro.Color = node.P.Color
				node.P.makeBlack()
				bro.Left.makeBlack()
				rightRotate(tree, node.P)
				node = tree.Root
			}
		}
	}

	node.makeBlack()
}

func deleteKey(tree *RBTree, node *Node) {
	new := node
	color := new.Color
	var next *Node

	switch {
	case node.Left.isNil() && !node.Right.isNil():
		next = node.Right
		transplant(tree, node, node.Right)
	case node.Right.isNil() && !node.Left.isNil():
		next = node.Left
		transplant(tree, node, node.Left)
	default:
		new = findMin(node.Right)
		color = new.Color
		next = new.Right
		if new.P == node {
			next.P = new
		} else {
			transplant(tree, new, new.Right)
			new.Right = node.Right
			new.Right.P = new
		}

		transplant(tree, node, new)
		new.Left = node.Left
		new.Left.P = new
		new.Color = node.Color
	}

	if color == black {
		deleteFixup(tree, next)
	}
}
