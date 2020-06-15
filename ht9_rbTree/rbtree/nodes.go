package rbtree

type color string

const red color = "red"
const black color = "black"

// Node type
type Node struct {
	P     *Node
	Left  *Node
	Right *Node
	Key   *int
	Color color
}

func makeNil() *Node {
	return &Node{
		P:     nil,
		Left:  nil,
		Right: nil,
		Key:   nil,
		Color: black,
	}
}

func makeNode(key *int) *Node {
	return &Node{
		P:     nil,
		Left:  makeNil(),
		Right: makeNil(),
		Key:   key,
		Color: red,
	}
}

func (node *Node) makeRed() {
	node.Color = red
}

func (node *Node) makeBlack() {
	node.Color = black
}

func (node *Node) isNil() bool {
	return node == nil || node.Key == nil
}

func (node *Node) isRed() bool {
	return !node.isNil() && node.Color == red
}

func (node *Node) isBlack() bool {
	return node.Color == black
}
