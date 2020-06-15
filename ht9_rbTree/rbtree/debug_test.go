package rbtree

import (
	"reflect"
	"testing"
)

func Test_getKey(t *testing.T) {
	type args struct {
		x *Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"correct work", args{nodeFive}, "5"},
		{"correct work nil", args{nodeNil}, "nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKey(tt.args.x); got != tt.want {
				t.Errorf("getKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeToString(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)

	type args struct {
		tree *RBTree
		x    *Node
	}
	tests := []struct {
		name string
		args args
		want nodeAsString
	}{
		{"correct work", args{tree, nodeFive}, nodeAsString{key: 5, color: red, isRoot: false, left: "left:nil", right: "right:nil"}},
		{"correct work root", args{tree, tree.Root}, nodeAsString{key: 1, color: black, isRoot: true, left: "left:nil", right: "right:nil"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodeToString(tt.args.tree, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeToString(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)

	type args struct {
		tree *RBTree
		x    *Node
	}
	tests := []struct {
		name string
		args args
		want []nodeAsString
	}{
		{"correct work root", args{tree, tree.Root}, []nodeAsString{nodeAsString{key: 1, color: black, isRoot: true, left: "left:nil", right: "right:nil"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeToString(tt.args.tree, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
