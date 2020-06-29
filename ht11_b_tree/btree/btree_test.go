package btree

import (
	"reflect"
	"testing"
)

var (
	k1 int = 1
	k2 int = 2
	k3 int = 3
)

var node *Node = &Node{keysCount: 3, leaf: false, keys: []*int{&k1, &k2, &k3}, children: []*Node{}, degree: 2, parent: nil}

func TestNode_String(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want string
	}{
		{"correct", node, "keys: {1,2,3}, children: [empty], isNotLeaf, count: 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeTree(t *testing.T) {
	tree := &BTree{
		Root:   &Node{keysCount: 0, leaf: true, keys: make([]*int, 3), children: make([]*Node, 4), degree: 2, parent: nil},
		Degree: 2,
	}
	type args struct {
		degree int
	}
	tests := []struct {
		name string
		args args
		want *BTree
	}{
		{"correct", args{2}, tree},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTree(tt.args.degree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Min(t *testing.T) {
	tree := MakeTree(2)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(8)

	tests := []struct {
		name string
		tr   *BTree
		want int
	}{
		{"correct", tree, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Min(); got != tt.want {
				t.Errorf("BTree.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Max(t *testing.T) {
	tree := MakeTree(2)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(8)

	tests := []struct {
		name string
		tr   *BTree
		want int
	}{
		{"correct", tree, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Max(); got != tt.want {
				t.Errorf("BTree.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Insert(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		tr   *BTree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Insert(tt.args.key)
		})
	}
}

func TestBTree_InOrder(t *testing.T) {
	tree := MakeTree(2)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(8)

	tests := []struct {
		name string
		tr   *BTree
		want []int
	}{
		{"correct", tree, []int{2, 6, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BTree.InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Exists(t *testing.T) {
	tree := MakeTree(2)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(8)

	type args struct {
		key int
	}
	tests := []struct {
		name string
		tr   *BTree
		args args
		want bool
	}{
		{"is exist", tree, args{6}, true},
		{"is not exist", tree, args{7}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Exists(tt.args.key); got != tt.want {
				t.Errorf("BTree.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Delete(t *testing.T) {
	tree := MakeTree(2)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(8)

	type args struct {
		k int
	}
	tests := []struct {
		name string
		tr   *BTree
		args args
		want bool
	}{
		{"is deleted", tree, args{6}, true},
		{"is not deleted", tree, args{6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Delete(tt.args.k); got != tt.want {
				t.Errorf("BTree.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
