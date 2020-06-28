package btree

import (
	"strconv"
	"strings"
)

// Node type
type Node struct {
	keysCount int
	keys      []*int
	children  []*Node
	leaf      bool
	degree    int
	parent    *Node
}

// String
func (n *Node) String() string {
	var leaf string

	if n.isLeaf() {
		leaf = "isLeaf"
	} else {
		leaf = "isNotLeaf"
	}

	keys := []string{}
	for _, key := range n.keys {
		if key == nil {
			continue
		}
		keys = append(keys, strconv.Itoa(*key))
	}
	if len(keys) == 0 {
		keys = []string{"empty"}
	}

	children := []string{}
	for i, child := range n.children {
		if child == nil {
			continue
		}
		children = append(children, "c"+strconv.Itoa(i)+": "+child.String())
	}
	if len(children) == 0 {
		children = []string{"empty"}
	}

	return "keys: {" + strings.Join(keys, ",") + "}, children: [" + strings.Join(children, ", ") + "], " + leaf + ", count: " + strconv.Itoa(n.keysCount)
}

// BTree type
type BTree struct {
	Root   *Node
	Degree int
}

// MakeTree func
func MakeTree(degree int) *BTree {
	if degree < 2 {
		degree = 2
	}

	return &BTree{
		Root:   makeNode(degree, true),
		Degree: degree,
	}
}

// Min func
func (t *BTree) Min() int {
	if t.Root.keysCount == 0 {
		return 0
	}

	minLeaf := t.Root.findMinLeaf()

	return minLeaf.findMinKey()
}

// Max func
func (t *BTree) Max() int {
	if t.Root.keysCount == 0 {
		return 0
	}

	maxLeaf := t.Root.findMaxLeaf()

	return maxLeaf.findMaxKey()
}

// Insert func
func (t *BTree) Insert(key int) {
	root := t.Root

	if root.isFull() {
		new := makeNode(t.Degree, false)
		t.Root = new
		new.children[0] = root
		new.splitChildNode(0)
		new.insertToNotFullNode(key)
	} else {
		root.insertToNotFullNode(key)
	}
}

// InOrder func
func (t *BTree) InOrder() []int {
	// fmt.Printf("\nfull tree: %v\n", t.Root.String())
	return t.Root.inOrder()
}

// Exists func
func (t *BTree) Exists(key int) bool {
	node, _ := t.Root.findNode(key)

	return node != nil
}

// Delete func
func (t *BTree) Delete(k int) bool {
	if t.Root.keysCount == 0 || !t.Exists(k) {
		return false
	}

	// Call the remove function for root
	t.Root.remove(k)

	// If the root node has 0 keys, make its first child as the new root
	// if it has a child, otherwise set root as NULL
	if t.Root.keysCount == 0 {
		if t.Root.isLeaf() {
			t.Root = makeNode(t.Degree, true)
		} else {
			t.Root = t.Root.children[0]
		}
	}
	return true
}
