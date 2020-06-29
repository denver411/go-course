package hashtable

import (
	"reflect"
	"testing"
)

func TestMakeTable(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *HashTable
	}{
		{"correct", args{10}, &HashTable{
			Data: make([]*list, 10),
			Size: 10,
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTable(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Set(t *testing.T) {
	type args struct {
		key []byte
		val []byte
	}
	tests := []struct {
		name string
		tr   *HashTable
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Set(tt.args.key, tt.args.val)
		})
	}
}

func TestHashTable_Get(t *testing.T) {
	table := &HashTable{
		Data: make([]*list, 10),
		Size: 10,
	}
	table.Set([]byte("key1"), []byte("value1"))

	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		tr   *HashTable
		args args
		want []byte
	}{
		{"is exist", table, args{[]byte("key1")}, []byte("value1")},
		{"is not exist", table, args{[]byte("key2")}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashTable.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Delete(t *testing.T) {
	table := &HashTable{
		Data: make([]*list, 10),
		Size: 10,
	}
	table.Set([]byte("key1"), []byte("value1"))
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		tr   *HashTable
		args args
		want bool
	}{
		{"is deleted", table, args{[]byte("key1")}, true},
		{"is not teleted", table, args{[]byte("key2")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Delete(tt.args.key); got != tt.want {
				t.Errorf("HashTable.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
