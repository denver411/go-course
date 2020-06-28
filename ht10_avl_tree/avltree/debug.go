package avltree

import (
	"strconv"
)

type nodeAsString struct {
	key    int
	height string
	isRoot bool
	left   string
	right  string
}

func getKey(x *Node) string {
	if x.isNil() {
		return "nil"
	}

	return strconv.Itoa(x.Key)
}

func nodeToString(x *Node) nodeAsString {
	if x.isNil() {
		return nodeAsString{key: 0, height: "", isRoot: false, left: "", right: ""}
	}

	return nodeAsString{key: x.Key, height: "height: " + strconv.Itoa(x.Height), isRoot: x.P.isNil(), left: "left:" + getKey(x.Left), right: "right:" + getKey(x.Right)}
}

func treeToString(x *Node) []nodeAsString {
	if x.isNil() {
		return []nodeAsString{}
	}

	left := treeToString(x.Left)
	right := treeToString(x.Right)

	return append(left, append([]nodeAsString{nodeToString(x)}, right...)...)
}
