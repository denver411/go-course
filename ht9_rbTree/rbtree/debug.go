package rbtree

import (
	"strconv"
)

type nodeAsString struct {
	key    int
	color  color
	isRoot bool
	left   string
	right  string
}

func getKey(x *Node) string {
	if x.isNil() {
		return "nil"
	}

	return strconv.Itoa(*x.Key)
}

func nodeToString(tree *RBTree, x *Node) nodeAsString {
	return nodeAsString{key: *x.Key, color: x.Color, isRoot: x == tree.Root, left: "left:" + getKey(x.Left), right: "right:" + getKey(x.Right)}
}

func treeToString(tree *RBTree, x *Node) []nodeAsString {
	if x.isNil() {
		return []nodeAsString{}
	}

	left := treeToString(tree, x.Left)
	right := treeToString(tree, x.Right)

	return append(left, append([]nodeAsString{nodeToString(tree, x)}, right...)...)
}
