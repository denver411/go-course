package avltree

import (
	"reflect"
	"testing"
)

func Test_findMin(t *testing.T) {
	tree := &AVLTree{
		Root: &Node{
			P:      nil,
			Left:   nodeOne,
			Right:  nil,
			Key:    2,
			Height: 1,
		},
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct", args{tree.Root}, nodeOne},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMax(t *testing.T) {
	tree := &AVLTree{
		Root: &Node{
			P:      nil,
			Left:   nil,
			Right:  nodeFive,
			Key:    2,
			Height: 1,
		},
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct", args{tree.Root}, nodeFive},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMax(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inOrderTree(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(14)

	type args struct {
		x *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"correct", args{tree.Root}, []int{1, 3, 5, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inOrderTree(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inOrderTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNode(t *testing.T) {
	tree := MakeTree()
	tree.Root = nodeOne
	tree.Root.Right = nodeFive
	nodeFive.P = nodeOne
	tree.Root.Left = nodeThree
	nodeThree.P = nodeOne

	type args struct {
		tree *AVLTree
		key  int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct", args{tree, 5}, nodeFive},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNode(tt.args.tree, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leftRotate(t *testing.T) {
	type args struct {
		tree *AVLTree
		node *Node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			leftRotate(tt.args.tree, tt.args.node)
		})
	}
}

func Test_rightRotate(t *testing.T) {
	type args struct {
		tree *AVLTree
		node *Node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rightRotate(tt.args.tree, tt.args.node)
		})
	}
}

func Test_balance(t *testing.T) {
	type args struct {
		tree *AVLTree
		node *Node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance(tt.args.tree, tt.args.node)
		})
	}
}

func Test_transplant(t *testing.T) {
	type args struct {
		tree *AVLTree
		old  *Node
		new  *Node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transplant(tt.args.tree, tt.args.old, tt.args.new)
		})
	}
}

func Test_deleteKey(t *testing.T) {
	type args struct {
		tree *AVLTree
		node *Node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteKey(tt.args.tree, tt.args.node)
		})
	}
}
