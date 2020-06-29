package avltree

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
		{"correct", args{&Node{
			P:      nil,
			Left:   nil,
			Right:  nil,
			Key:    4,
			Height: 1,
		}}, "4"},
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
	type args struct {
		x *Node
	}
	tests := []struct {
		name string
		args args
		want nodeAsString
	}{
		{"correct", args{&Node{
			P:      nil,
			Left:   nil,
			Right:  nil,
			Key:    4,
			Height: 1,
		}}, nodeAsString{
			key:    4,
			height: "height: 1",
			isRoot: true,
			left:   "left:nil",
			right:  "right:nil",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodeToString(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeToString(t *testing.T) {
	type args struct {
		x *Node
	}
	tests := []struct {
		name string
		args args
		want []nodeAsString
	}{
		{"correct", args{&Node{
			P:      nil,
			Left:   nil,
			Right:  nil,
			Key:    4,
			Height: 1,
		}}, []nodeAsString{
			nodeAsString{
				key:    4,
				height: "height: 1",
				isRoot: true,
				left:   "left:nil",
				right:  "right:nil",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeToString(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
