package hashtable

type node struct {
	key   []byte
	value []byte
	next  *node
}

func makeNode(key, value []byte) *node {
	return &node{
		key:   key,
		value: value,
		next:  nil,
	}
}

func (n *node) isNil() bool {
	return n == nil
}

func (n *node) isEqualKey(key []byte) bool {
	if len(n.key) != len(key) {
		return false
	}

	for i := range n.key {
		if n.key[i] != key[i] {
			return false
		}
	}

	return true
}

func (n *node) String() string {
	if n.isNil() {
		return "nil"
	}

	s := "k: " + string(n.key[:]) + ", v: " + string(n.value[:])

	return "data: " + s + ", next: " + n.next.String()

}

type list struct {
	head *node
	size int
}

func makeList() *list {
	return &list{
		size: 0,
		head: nil,
	}
}

func (l *list) insert(key, value []byte) int {
	new := makeNode(key, value)
	cur := l.head

	for !cur.isNil() && !cur.next.isNil() {
		cur = cur.next
	}

	if cur.isNil() {
		l.head = new
	} else {
		cur.next = new
	}

	l.size = l.size + 1

	return l.size
}

func (l *list) get(key []byte) (cur, prev *node) {
	if l.head.isNil() {
		return nil, nil
	}
	if l.head.isEqualKey(key) {
		return l.head, nil
	}

	prev = l.head

	for !prev.next.isNil() && !prev.next.isEqualKey(key) {
		prev = prev.next
	}

	if prev.next.isNil() {
		return nil, nil
	}

	return prev.next, prev
}

func (l *list) delete(key []byte) bool {
	cur, prev := l.get(key)

	if cur.isNil() {
		return false
	}

	if prev.isNil() {
		l.head = cur.next
	} else {
		prev.next = cur.next
	}

	l.size = l.size - 1

	return true
}
