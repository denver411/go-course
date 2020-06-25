package btree

func makeNode(degree int, isLeaf bool) *Node {
	return &Node{keysCount: 0, leaf: isLeaf, keys: make([]*int, 2*degree-1), children: make([]*Node, 2*degree), degree: degree, parent: nil}
}

func (n *Node) isLeaf() bool {
	return n.leaf
}

func (n *Node) isFull() bool {
	return n.keysCount == 2*n.degree-1
}

func (n *Node) isEmpty() bool {
	return n.keysCount == n.degree-1
}

func (n *Node) findMaxIdx() int {
	if n.keys == nil {
		return 0
	}

	max := 0

	for l := len(n.keys) - 1; l >= 0; l-- {
		if n.keys[l] != nil {
			max = l
			break
		}
	}
	return max
}

func (n *Node) findMinKey() int {
	if n.keys == nil || n.keys[0] == nil {
		return 0
	}
	return *n.keys[0]
}

func (n *Node) findMaxKey() int {
	max := n.findMaxIdx()

	return *n.keys[max]
}

func (n *Node) findMaxChild() *Node {
	var max *Node

	for l := len(n.children) - 1; l >= 0; l-- {
		if n.children[l] != nil {
			max = n.children[l]
			break
		}
	}
	return max
}

func (n *Node) splitChildNode(idx int) {
	splitted := n.children[idx]
	degree := n.degree
	new := makeNode(degree, splitted.isLeaf())
	new.keysCount = degree - 1

	for i := 0; i < degree-1; i++ {
		new.keys[i] = splitted.keys[i+degree]
	}

	if !splitted.isLeaf() {
		for i := 0; i < degree; i++ {
			new.children[i] = splitted.children[i+degree]
		}
		splitted.children = splitted.children[:degree]
	}

	for l := n.keysCount; l > idx; l-- {
		n.children[l] = n.children[l-1]
	}

	n.children[idx+1] = new

	for l := n.keysCount; l > idx; l-- {
		n.keys[l] = n.keys[l-1]
	}

	n.keys[idx] = splitted.keys[degree-1]
	splitted.keysCount = degree - 1
	splitted.keys = splitted.keys[:(degree - 1)]
	n.keysCount = n.keysCount + 1
	new.parent = n
	splitted.parent = n
}

func (n *Node) insertToNotFullNode(key int) {
	idx := n.keysCount - 1

	if n.isLeaf() {
		for idx >= 0 && key <= *n.keys[idx] {
			n.keys[idx+1] = n.keys[idx]
			idx = idx - 1
		}

		if n.keys == nil {
			n.keys = []*int{&key}
		} else {
			n.keys[idx+1] = &key
		}
		n.keysCount = n.keysCount + 1

	} else {
		for idx >= 0 && key <= *n.keys[idx] {
			idx = idx - 1
		}

		idx = idx + 1

		if n.children[idx].isFull() {
			n.splitChildNode(idx)
			if key > *n.keys[idx] {
				idx = idx + 1
			}
		}

		n.children[idx].insertToNotFullNode(key)
	}
}

func (n *Node) findKey(k int) int {
	i := 0

	for i < n.keysCount && *n.keys[i] < k {
		i = i + 1
	}

	return i
}

func (n *Node) remove(k int) {
	i := n.findKey(k)

	if i < n.keysCount && *n.keys[i] == k {
		if n.isLeaf() {
			n.removeFromLeaf(i)
		} else {
			n.removeFromNotLeaf(i)
		}
	} else {
		if n.isLeaf() {
			return
		}

		var flag bool
		if i == n.keysCount {
			flag = true
		} else {
			flag = false
		}

		if n.children[i].isEmpty() {
			n.fill(i)
		}

		if flag == true && i > n.keysCount {
			n.children[i-1].remove(k)
		} else {
			n.children[i].remove(k)
		}
	}
}

func (n *Node) removeFromLeaf(i int) {
	n.keys = append(n.keys[:i], n.keys[i+1:]...)
	n.keysCount = n.keysCount - 1
}

func (n *Node) removeFromNotLeaf(i int) {
	k := n.keys[i]

	if !n.children[i].isEmpty() {
		prev := n.getPrev(i)
		n.keys[i] = &prev
		n.children[i].remove(prev)

	} else if !n.children[i+1].isEmpty() {
		next := n.getNext(i)
		n.keys[i] = &next
		n.children[i+1].remove(next)

	} else {
		n.merge(i)
		n.children[i].remove(*k)
	}
}

func (n *Node) getPrev(i int) int {
	current := n.children[i]

	for !current.isLeaf() {
		current = current.children[current.keysCount]
	}

	return *current.keys[current.keysCount-1]
}

func (n *Node) getNext(i int) int {

	current := n.children[i+1]
	for !current.isLeaf() {
		current = current.children[0]
	}

	return *current.keys[0]
}

func (n *Node) fill(i int) {
	if i > 0 && !n.children[i-1].isEmpty() {
		n.borrowFromPrev(i)

	} else if i < n.keysCount && !n.children[i+1].isEmpty() {
		n.borrowFromNext(i)

	} else {
		if i < n.keysCount {
			n.merge(i)
		} else {
			n.merge(i - 1)
		}
	}
}

func (n *Node) borrowFromPrev(i int) {

	prev := n.children[i]
	next := n.children[i-1]

	for i := prev.keysCount - 1; i >= 0; i-- {
		prev.keys[i+1] = prev.keys[i]
	}

	if !prev.isLeaf() {
		for i := prev.keysCount; i >= 0; i-- {
			prev.children[i+1] = prev.children[i]
		}
	}

	prev.keys[0] = n.keys[i-1]

	if !prev.isLeaf() {
		prev.children[0] = next.children[next.keysCount]
	}

	n.keys[i-1] = next.keys[next.keysCount-1]

	prev.keysCount = prev.keysCount + 1
	next.keysCount = next.keysCount - 1

}

func (n *Node) borrowFromNext(i int) {

	prev := n.children[i]
	next := n.children[i+1]

	prev.keys = append(prev.keys, n.keys[i])

	if !prev.isLeaf() {
		prev.children = append(prev.children, next.children[0])
	}

	n.keys[i] = next.keys[0]

	next.keys = append(next.keys[:i], next.keys[i+1:]...)

	if !next.isLeaf() {
		next.children = append(next.children[:i], next.children[i+1:]...)
	}

	prev.keysCount = prev.keysCount + 1
	next.keysCount = next.keysCount - 1
}

func (n *Node) merge(i int) {
	prev := n.children[i]
	next := n.children[i+1]

	prev.keys = append(prev.keys, n.keys[i])
	prev.keys = append(prev.keys, next.keys...)

	if !prev.isLeaf() {
		prev.children = append(prev.children, next.children...)
	}

	n.keys = append(n.keys[:i], n.keys[i+1:]...)

	n.children = append(n.children[:i+1], n.children[i+2:]...)

	prev.keysCount = prev.keysCount + next.keysCount + 1
	n.keysCount = n.keysCount - 1

	next.parent = nil
}

func (n *Node) findNode(key int) (*Node, int) {
	idx := 0

	for idx < n.keysCount && key > *n.keys[idx] {
		idx = idx + 1
	}

	if idx < n.keysCount && key == *n.keys[idx] {
		return n, idx
	} else if n.isLeaf() {
		return nil, -1
	} else {
		return n.children[idx].findNode(key)
	}

}

func (n *Node) findMinLeaf() *Node {
	for n.children[0] != nil {
		n = n.children[0]
	}

	return n
}

func (n *Node) findMaxLeaf() *Node {
	for n.children[0] != nil {
		n = n.findMaxChild()
	}

	return n
}

func (n *Node) inOrder() []int {
	if n.keysCount == 0 {
		return nil
	}
	if n.isLeaf() {
		return pointsToInts(n.keys)
	}
	a := n.children[0].inOrder()

	for i := 0; i < n.keysCount; i++ {
		a = append(a, append([]int{*n.keys[i]}, n.children[i+1].inOrder()...)...)
	}

	return a
}
