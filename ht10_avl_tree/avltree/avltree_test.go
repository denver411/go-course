package avltree

import (
	"reflect"
	"testing"
)

var nodeOne = makeNode(1)
var nodeThree = makeNode(3)
var nodeFive = makeNode(5)

func TestMakeTree(t *testing.T) {
	tests := []struct {
		name string
		want *AVLTree
	}{
		{"correct", &AVLTree{Root: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_Min(t *testing.T) {
	tree := MakeTree()
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)

	tests := []struct {
		name string
		tree *AVLTree
		want int
	}{
		{"correct", tree, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Min(); got != tt.want {
				t.Errorf("AVLTree.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_Max(t *testing.T) {
	tree := MakeTree()
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)

	tests := []struct {
		name string
		tree *AVLTree
		want int
	}{
		{"correct", tree, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Max(); got != tt.want {
				t.Errorf("AVLTree.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_Insert(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *AVLTree
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

func TestAVLTree_InOrder(t *testing.T) {
	tree := MakeTree()
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)

	tests := []struct {
		name string
		tree *AVLTree
		want []int
	}{
		{"correct", tree, []int{6, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AVLTree.InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_Exists(t *testing.T) {
	tree := MakeTree()
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)

	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *AVLTree
		args args
		want bool
	}{
		{"is exist", tree, args{7}, true},
		{"is not exist", tree, args{10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Exists(tt.args.key); got != tt.want {
				t.Errorf("AVLTree.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVLTree_Delete(t *testing.T) {
	tree := MakeTree()
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)

	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *AVLTree
		args args
		want bool
	}{
		{"is deleted", tree, args{6}, true},
		{"is not deleted", tree, args{10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Delete(tt.args.key); got != tt.want {
				t.Errorf("AVLTree.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
