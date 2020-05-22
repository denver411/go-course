package bintree

import (
	"reflect"
	"testing"
)

func TestMakeTree(t *testing.T) {
	tests := []struct {
		name string
		want *BinaryTree
	}{
		{"correct work", &BinaryTree{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Min(t *testing.T) {
	tests := []struct {
		name string
		tree *BinaryTree
		want int
	}{
		{"correct work", &BinaryTree{Root: createSub(nil, node4, node6, 5)}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Min(); got != tt.want {
				t.Errorf("BinaryTree.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Max(t *testing.T) {
	tests := []struct {
		name string
		tree *BinaryTree
		want int
	}{
		{"correct work", &BinaryTree{Root: createSub(nil, node4, node6, 5)}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Max(); got != tt.want {
				t.Errorf("BinaryTree.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_InOrder(t *testing.T) {
	tests := []struct {
		name string
		tree *BinaryTree
		want []int
	}{
		{"correct work", &BinaryTree{Root: createSub(nil, node4, node6, 5)}, []int{4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTree.InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Exists(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *BinaryTree
		args args
		want bool
	}{
		{"is exist", &BinaryTree{Root: createSub(nil, node4, node6, 5)}, args{4}, true},
		{"is not exist", &BinaryTree{Root: createSub(nil, node4, node6, 5)}, args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Exists(tt.args.key); got != tt.want {
				t.Errorf("BinaryTree.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Delete(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *BinaryTree
		args args
		want bool
	}{
		{"is deleted", &BinaryTree{Root: createSub(nil, node4, node6, 5)}, args{6}, true},
		{"is  not deleted", &BinaryTree{Root: createSub(nil, node4, node6, 5)}, args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Delete(tt.args.key); got != tt.want {
				t.Errorf("BinaryTree.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
