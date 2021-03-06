package rbtree

import (
	"reflect"
	"testing"
)

func Test_makeNil(t *testing.T) {
	tests := []struct {
		name string
		want *Node
	}{
		{"correct", &Node{
			P:     nil,
			Left:  nil,
			Right: nil,
			Key:   nil,
			Color: black,
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeNil(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeNode(t *testing.T) {
	type args struct {
		key *int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"correct", args{&k5}, &Node{
			P:     nil,
			Left:  makeNil(),
			Right: makeNil(),
			Key:   &k5,
			Color: red,
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

func TestNode_makeRed(t *testing.T) {
	tests := []struct {
		name string
		node *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.node.makeBlack()
		})
	}
}

func TestNode_makeBlack(t *testing.T) {
	tests := []struct {
		name string
		node *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.node.makeBlack()
		})
	}
}

func TestNode_isNil(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		want bool
	}{
		{"is nil", nodeNil, true},
		{"is not nil", nodeOne, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.isNil(); got != tt.want {
				t.Errorf("Node.isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_isRed(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		want bool
	}{
		{"is red", nodeOne, true},
		{"is not red", nodeNil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.isRed(); got != tt.want {
				t.Errorf("Node.isRed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_isBlack(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		want bool
	}{
		{"is black", nodeNil, true},
		{"is not black", nodeOne, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.isBlack(); got != tt.want {
				t.Errorf("Node.isBlack() = %v, want %v", got, tt.want)
			}
		})
	}
}
