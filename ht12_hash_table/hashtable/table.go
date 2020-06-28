package hashtable

// HashTable type
type HashTable struct {
	Data []*list
	Size int
}

// MakeTable func
func MakeTable(size int) *HashTable {
	return &HashTable{
		Data: make([]*list, size),
		Size: size,
	}
}

// Set func
func (t *HashTable) Set(key, val []byte) {
	i := t.makeHash(key)

	if t.Data[i] == nil {
		l := makeList()
		l.insert(key, val)
		t.Data[i] = l
		return
	}

	l := t.Data[i]
	l.insert(key, val)

}

// Get func
func (t *HashTable) Get(key []byte) []byte {
	i := t.makeHash(key)

	if t.Data[i] == nil {
		return nil
	}

	n, _ := t.Data[i].get(key)

	if n.isNil() {
		return nil
	}

	return n.value
}

// Delete func
func (t *HashTable) Delete(key []byte) bool {
	i := t.makeHash(key)

	if t.Data[i] == nil {
		return false
	}

	return t.Data[i].delete(key)

}
