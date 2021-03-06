package rbtree

import (
	"reflect"
	"testing"
)

var k1 = 1
var k3 = 3
var k5 = 5
var nodeOne = makeNode(&k1)
var nodeThree = makeNode(&k3)
var nodeFive = makeNode(&k5)
var nodeNil = makeNil()

func TestMakeTree(t *testing.T) {
	tests := []struct {
		name string
		want *RBTree
	}{
		{"correct", &RBTree{Root: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRBTree_Min(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	tests := []struct {
		name string
		tree *RBTree
		want int
	}{
		{"correct", tree, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Min(); got != tt.want {
				t.Errorf("RBTree.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRBTree_Max(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	tests := []struct {
		name string
		tree *RBTree
		want int
	}{
		{"correct", tree, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Max(); got != tt.want {
				t.Errorf("RBTree.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRBTree_Insert(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *RBTree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tree.Insert(tt.args.key)
		})
	}
}

func TestRBTree_InOrder(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	tests := []struct {
		name string
		tree *RBTree
		want []int
	}{
		{"correct", tree, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RBTree.InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRBTree_Exists(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *RBTree
		args args
		want bool
	}{
		{"is exist", tree, args{1}, true},
		{"is not exist", tree, args{10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Exists(tt.args.key); got != tt.want {
				t.Errorf("RBTree.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRBTree_Delete(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *RBTree
		args args
		want bool
	}{
		{"is deleted", tree, args{1}, true},
		{"is not deleted", tree, args{10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Delete(tt.args.key); got != tt.want {
				t.Errorf("RBTree.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
