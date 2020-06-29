package btree

import (
	"reflect"
	"testing"
)

func Test_makeNode(t *testing.T) {
	type args struct {
		degree int
		isLeaf bool
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct", args{5, true}, &Node{keysCount: 0, leaf: true, keys: make([]*int, 9), children: make([]*Node, 10), degree: 5, parent: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeNode(tt.args.degree, tt.args.isLeaf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_isLeaf(t *testing.T) {
	leaf := &Node{keysCount: 0, leaf: true, keys: nil, children: nil, degree: 2, parent: nil}
	notLeaf := &Node{keysCount: 0, leaf: false, keys: nil, children: nil, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want bool
	}{
		{"is leaf", leaf, true},
		{"is not leaf", notLeaf, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.isLeaf(); got != tt.want {
				t.Errorf("Node.isLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_isFull(t *testing.T) {
	full := &Node{keysCount: 3, leaf: true, keys: nil, children: nil, degree: 2, parent: nil}
	notFull := &Node{keysCount: 0, leaf: false, keys: nil, children: nil, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want bool
	}{
		{"is full", full, true},
		{"is not full", notFull, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.isFull(); got != tt.want {
				t.Errorf("Node.isFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_isEmpty(t *testing.T) {
	empty := &Node{keysCount: 1, leaf: true, keys: nil, children: nil, degree: 2, parent: nil}
	notEmpty := &Node{keysCount: 3, leaf: false, keys: nil, children: nil, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want bool
	}{
		{"is empty", empty, true},
		{"is not empty", notEmpty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.isEmpty(); got != tt.want {
				t.Errorf("Node.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_findMaxIdx(t *testing.T) {
	node := &Node{keysCount: 2, leaf: false, keys: []*int{&k1, &k2}, children: nil, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want int
	}{
		{"correct", node, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.findMaxIdx(); got != tt.want {
				t.Errorf("Node.findMaxIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_findMinKey(t *testing.T) {
	node := &Node{keysCount: 2, leaf: false, keys: []*int{&k1, &k2}, children: nil, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want int
	}{
		{"correct", node, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.findMinKey(); got != tt.want {
				t.Errorf("Node.findMinKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_findMaxKey(t *testing.T) {
	node := &Node{keysCount: 2, leaf: false, keys: []*int{&k1, &k2}, children: nil, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want int
	}{
		{"correct", node, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.findMaxKey(); got != tt.want {
				t.Errorf("Node.findMaxKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_findMaxChild(t *testing.T) {
	child1 := &Node{keysCount: 2, leaf: false, keys: []*int{&k1, &k2}, children: nil, degree: 2, parent: nil}
	child2 := &Node{keysCount: 2, leaf: false, keys: []*int{&k2, &k3}, children: nil, degree: 2, parent: nil}
	node := &Node{keysCount: 2, leaf: false, keys: nil, children: []*Node{child1, child2}, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want *Node
	}{
		{"correct", node, child2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.findMaxChild(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.findMaxChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_splitChildNode(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.splitChildNode(tt.args.idx)
		})
	}
}

func TestNode_insertToNotFullNode(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.insertToNotFullNode(tt.args.key)
		})
	}
}

func TestNode_findKey(t *testing.T) {
	node := &Node{keysCount: 3, leaf: true, keys: []*int{&k1, &k2, &k2}, children: nil, degree: 2, parent: nil}

	type args struct {
		k int
	}
	tests := []struct {
		name string
		n    *Node
		args args
		want int
	}{
		{"correct", node, args{2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.findKey(tt.args.k); got != tt.want {
				t.Errorf("Node.findKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_remove(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.remove(tt.args.k)
		})
	}
}

func TestNode_removeFromLeaf(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.removeFromLeaf(tt.args.i)
		})
	}
}

func TestNode_removeFromNotLeaf(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.removeFromNotLeaf(tt.args.i)
		})
	}
}

func TestNode_getPrev(t *testing.T) {
	child1 := &Node{keysCount: 1, leaf: true, keys: []*int{&k1}, children: nil, degree: 2, parent: nil}
	child2 := &Node{keysCount: 1, leaf: true, keys: []*int{&k3}, children: nil, degree: 2, parent: nil}
	node := &Node{keysCount: 1, leaf: false, keys: []*int{&k2}, children: []*Node{child1, child2}, degree: 2, parent: nil}

	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    *Node
		args args
		want int
	}{
		{"correct", node, args{0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.getPrev(tt.args.i); got != tt.want {
				t.Errorf("Node.getPrev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_getNext(t *testing.T) {
	child1 := &Node{keysCount: 1, leaf: true, keys: []*int{&k1}, children: nil, degree: 2, parent: nil}
	child2 := &Node{keysCount: 1, leaf: true, keys: []*int{&k3}, children: nil, degree: 2, parent: nil}
	node := &Node{keysCount: 1, leaf: false, keys: []*int{&k2}, children: []*Node{child1, child2}, degree: 2, parent: nil}

	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    *Node
		args args
		want int
	}{
		{"correct", node, args{0}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.getNext(tt.args.i); got != tt.want {
				t.Errorf("Node.getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_fill(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.fill(tt.args.i)
		})
	}
}

func TestNode_borrowFromPrev(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.borrowFromPrev(tt.args.i)
		})
	}
}

func TestNode_borrowFromNext(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.borrowFromNext(tt.args.i)
		})
	}
}

func TestNode_merge(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		n    *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.merge(tt.args.i)
		})
	}
}

func TestNode_findNode(t *testing.T) {
	child1 := &Node{keysCount: 1, leaf: true, keys: []*int{&k1}, children: nil, degree: 2, parent: nil}
	child2 := &Node{keysCount: 1, leaf: true, keys: []*int{&k3}, children: nil, degree: 2, parent: nil}
	node := &Node{keysCount: 1, leaf: false, keys: []*int{&k2}, children: []*Node{child1, child2}, degree: 2, parent: nil}

	type args struct {
		key int
	}
	tests := []struct {
		name  string
		n     *Node
		args  args
		want  *Node
		want1 int
	}{
		{"correct", node, args{3}, child2, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.n.findNode(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.findNode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Node.findNode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_findMinLeaf(t *testing.T) {
	child1 := &Node{keysCount: 1, leaf: true, keys: []*int{&k1}, children: make([]*Node, 3), degree: 2, parent: nil}
	child2 := &Node{keysCount: 1, leaf: true, keys: []*int{&k3}, children: make([]*Node, 3), degree: 2, parent: nil}
	node := &Node{keysCount: 1, leaf: false, keys: []*int{&k2}, children: []*Node{child1, child2}, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want *Node
	}{
		{"correct", node, child1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.findMinLeaf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.findMinLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_findMaxLeaf(t *testing.T) {
	child1 := &Node{keysCount: 1, leaf: true, keys: []*int{&k1}, children: make([]*Node, 3), degree: 2, parent: nil}
	child2 := &Node{keysCount: 1, leaf: true, keys: []*int{&k3}, children: make([]*Node, 3), degree: 2, parent: nil}
	node := &Node{keysCount: 1, leaf: false, keys: []*int{&k2}, children: []*Node{child1, child2}, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want *Node
	}{
		{"correct", node, child2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.findMaxLeaf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.findMaxLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_inOrder(t *testing.T) {
	child1 := &Node{keysCount: 1, leaf: true, keys: []*int{&k1}, children: make([]*Node, 3), degree: 2, parent: nil}
	child2 := &Node{keysCount: 1, leaf: true, keys: []*int{&k3}, children: make([]*Node, 3), degree: 2, parent: nil}
	node := &Node{keysCount: 1, leaf: false, keys: []*int{&k2}, children: []*Node{child1, child2}, degree: 2, parent: nil}

	tests := []struct {
		name string
		n    *Node
		want []int
	}{
		{"correct", node, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.inOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.inOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
