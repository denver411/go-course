package bintree

func makeNode(key int) *Node {
	return &Node{
		P:     nil,
		Left:  nil,
		Right: nil,
		Key:   key,
	}
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func findMax(node *Node) *Node {
	for node.Right != nil {
		node = node.Right
	}

	return node
}

func inOrderTree(x *Node) []int {
	if x == nil {
		return []int{}
	}
	left := inOrderTree(x.Left)
	right := inOrderTree(x.Right)

	return append(left, append([]int{x.Key}, right...)...)

}

func findNode(root *Node, key int) *Node {
	node := root

	for node != nil {
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

func (node *Node) isNil() bool {
	return node == nil
}

func transplant(tree *BinaryTree, old *Node, new *Node) {
	switch {
	case old.P == nil:
		tree.Root = new
	case old == old.P.Left:
		old.P.Left = new
	case old == old.P.Right:
		old.P.Right = new

	}
	if new != nil {
		new.P = old.P
	}
}

func deleteKey(tree *BinaryTree, node *Node) {
	switch {
	case node.Left.isNil() && node.Right.isNil():
		transplant(tree, node, nil)
	case node.Left.isNil():
		transplant(tree, node, node.Right)
	case node.Right.isNil():
		transplant(tree, node, node.Left)
	default:
		new := findMin(node.Right)
		if new.P != node {
			transplant(tree, new, new.Right)
			new.Right = node.Right
			new.Right.P = new
		}
		transplant(tree, node, new)
		new.Left = node.Left
		new.Left.P = new
	}
}
