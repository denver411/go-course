package hashtable

import (
	"reflect"
	"testing"
)

func Test_makeNode(t *testing.T) {
	type args struct {
		key   []byte
		value []byte
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{"correct", args{[]byte("key1"), []byte("value1")}, &node{
			key:   []byte("key1"),
			value: []byte("value1"),
			next:  nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeNode(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_isNil(t *testing.T) {
	tests := []struct {
		name string
		n    *node
		want bool
	}{
		{"is not nil", &node{
			key:   []byte("key1"),
			value: []byte("value1"),
			next:  nil,
		}, false},
		{"is nil", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.isNil(); got != tt.want {
				t.Errorf("node.isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_isEqualKey(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		n    *node
		args args
		want bool
	}{
		{"is equal", &node{
			key:   []byte("key1"),
			value: []byte("value1"),
			next:  nil,
		}, args{[]byte("key1")}, true},
		{"is not equal", &node{
			key:   []byte("key1"),
			value: []byte("value1"),
			next:  nil,
		}, args{[]byte("key2")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.isEqualKey(tt.args.key); got != tt.want {
				t.Errorf("node.isEqualKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_String(t *testing.T) {
	tests := []struct {
		name string
		n    *node
		want string
	}{
		{"correct", &node{
			key:   []byte("key1"),
			value: []byte("value1"),
			next:  nil,
		}, "data: k: key1, v: value1, next: nil"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeList(t *testing.T) {
	tests := []struct {
		name string
		want *list
	}{
		{"correct", &list{
			head: nil,
			size: 0,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_insert(t *testing.T) {
	type args struct {
		key   []byte
		value []byte
	}
	tests := []struct {
		name string
		l    *list
		args args
		want int
	}{
		{"correct", &list{
			head: nil,
			size: 0,
		}, args{[]byte("key1"), []byte("value1")}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.insert(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("list.insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_get(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name     string
		l        *list
		args     args
		wantCur  *node
		wantPrev *node
	}{
		{"correct", &list{
			head: &node{
				key:   []byte("key1"),
				value: []byte("value1"),
				next: &node{
					key:   []byte("key2"),
					value: []byte("value2"),
					next:  nil,
				},
			},
			size: 2,
		}, args{[]byte("key2")},
			&node{
				key:   []byte("key2"),
				value: []byte("value2"),
				next:  nil,
			},
			&node{
				key:   []byte("key1"),
				value: []byte("value1"),
				next: &node{
					key:   []byte("key2"),
					value: []byte("value2"),
					next:  nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCur, gotPrev := tt.l.get(tt.args.key)
			if !reflect.DeepEqual(gotCur, tt.wantCur) {
				t.Errorf("list.get() gotCur = %v, want %v", gotCur, tt.wantCur)
			}
			if !reflect.DeepEqual(gotPrev, tt.wantPrev) {
				t.Errorf("list.get() gotPrev = %v, want %v", gotPrev, tt.wantPrev)
			}
		})
	}
}

func Test_list_delete(t *testing.T) {
	l := &list{
		head: &node{
			key:   []byte("key1"),
			value: []byte("value1"),
			next: &node{
				key:   []byte("key2"),
				value: []byte("value2"),
				next:  nil,
			},
		},
		size: 2,
	}
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		l    *list
		args args
		want bool
	}{
		{"is deleted", l, args{[]byte("key2")}, true},
		{"is not deleted", l, args{[]byte("key3")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.delete(tt.args.key); got != tt.want {
				t.Errorf("list.delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
