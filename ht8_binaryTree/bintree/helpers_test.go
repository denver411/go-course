package bintree

import (
	"reflect"
	"testing"
)

var node4 = &Node{P: nil, Left: nil, Right: nil, Key: 4}
var node5 = &Node{P: nil, Left: nil, Right: nil, Key: 5}
var node6 = &Node{P: nil, Left: nil, Right: nil, Key: 6}
var node7 = &Node{P: nil, Left: nil, Right: nil, Key: 7}
var node8 = &Node{P: nil, Left: nil, Right: nil, Key: 8}

func createSub(p *Node, l *Node, r *Node, k int) *Node {
	return &Node{P: p, Left: l, Right: r, Key: k}
}

func Test_makeNode(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct work", args{4}, node4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeNode(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct work", args{createSub(nil, node4, node6, 5)}, node4},
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
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct work", args{createSub(nil, node4, node6, 5)}, node6},
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
	type args struct {
		x *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"correct work", args{createSub(nil, node4, node6, 5)}, []int{4, 5, 6}},
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
	type args struct {
		root *Node
		key  int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct work", args{createSub(nil, node4, node6, 5), 4}, node4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_isNil(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		want bool
	}{
		{"is Node", node4, false},
		{"is Nil", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.isNil(); got != tt.want {
				t.Errorf("Node.isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
