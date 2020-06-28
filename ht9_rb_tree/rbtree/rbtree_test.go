package rbtree

import (
	"reflect"
	"testing"
)

var k5 = 5
var nodeFive = makeNode(&k5)
var nodeNil = makeNil()

func TestMakeTree(t *testing.T) {
	tests := []struct {
		name string
		want *RBTree
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		tree *RBTree
		want int
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		tree *RBTree
		want int
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		tree *RBTree
		want []int
	}{
		// TODO: Add test cases.
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
	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *RBTree
		args args
		want bool
	}{
		// TODO: Add test cases.
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
	type args struct {
		key int
	}
	tests := []struct {
		name string
		tree *RBTree
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Delete(tt.args.key); got != tt.want {
				t.Errorf("RBTree.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
