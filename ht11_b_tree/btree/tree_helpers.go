package btree

func pointsToInts(a []*int) []int {
	var res []int

	for _, key := range a {
		if key == nil {
			break
		}
		res = append(res, *key)
	}

	return res
}
