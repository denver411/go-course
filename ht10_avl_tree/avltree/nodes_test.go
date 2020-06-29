package avltree

import (
	"reflect"
	"testing"
)

func Test_makeNode(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct", args{5}, &Node{
			P:      nil,
			Left:   nil,
			Right:  nil,
			Key:    5,
			Height: 1,
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeNode(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_getHeight(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		want int
	}{
		{"correct", &Node{
			P:      nil,
			Left:   nil,
			Right:  nil,
			Key:    5,
			Height: 1,
		}, 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.getHeight(); got != tt.want {
				t.Errorf("Node.getHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_setHeight(t *testing.T) {
	type args struct {
		h int
	}
	tests := []struct {
		name string
		node *Node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.node.setHeight(tt.args.h)
		})
	}
}

func TestNode_isNil(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		want bool
	}{
		{"is not nil", &Node{
			P:      nil,
			Left:   nil,
			Right:  nil,
			Key:    5,
			Height: 1,
		}, false,
		},
		{"is nil", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.isNil(); got != tt.want {
				t.Errorf("Node.isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_balanceFactor(t *testing.T) {
	tree := MakeTree()
	tree.Root = nodeOne
	tree.Root.Right = nodeFive
	nodeFive.P = nodeOne
	tree.Root.Left = nodeThree
	nodeThree.P = nodeOne

	tests := []struct {
		name string
		node *Node
		want int
	}{
		{"correct", tree.Root, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.balanceFactor(); got != tt.want {
				t.Errorf("Node.balanceFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_fixHeight(t *testing.T) {
	tests := []struct {
		name string
		node *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.node.fixHeight()
		})
	}
}
