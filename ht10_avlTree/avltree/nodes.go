package avltree

// Node type
type Node struct {
	P      *Node
	Left   *Node
	Right  *Node
	Key    int
	Height int
}

func makeNode(key int) *Node {
	return &Node{
		P:      nil,
		Left:   nil,
		Right:  nil,
		Key:    key,
		Height: 1,
	}
}

func (node *Node) getHeight() int {
	if node.isNil() {
		return 0
	}

	return node.Height
}

func (node *Node) setHeight(h int) {
	if !node.isNil() {
		node.Height = h
	}
}

func (node *Node) isNil() bool {
	return node == nil
}

func (node *Node) balanceFactor() int {
	if node.isNil() {
		return 0
	}

	return node.Right.getHeight() - node.Left.getHeight()
}

func (node *Node) fixHeight() {
	if node.isNil() {
		return
	}

	var height int

	if node.Left.getHeight() > node.Right.getHeight() {
		height = node.Left.getHeight()
	} else {
		height = node.Right.getHeight()
	}

	node.setHeight(height + 1)
}
