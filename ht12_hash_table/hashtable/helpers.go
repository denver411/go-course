package hashtable

import "strconv"

func (t *HashTable) makeHash(key []byte) int {
	var s byte

	for _, b := range key {
		s += b
	}

	return int(s) % t.Size
}

func (t *HashTable) String() string {
	if t == nil {
		return "nil"
	}

	s := ""
	for i, l := range t.Data {
		if l == nil {
			s += "\n" + strconv.Itoa(i) + ": nil"
			continue
		}
		s += "\n" + strconv.Itoa(i) + ": " + l.head.String()
	}

	return s

}
