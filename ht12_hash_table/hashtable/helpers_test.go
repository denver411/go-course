package hashtable

import "testing"

func TestHashTable_makeHash(t *testing.T) {
	tab := &HashTable{
		Data: nil,
		Size: 10,
	}

	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		tr   *HashTable
		args args
		want int
	}{
		{"correct", tab, args{[]byte("key1")}, 2},
		{"correct", tab, args{[]byte("key2")}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.makeHash(tt.args.key); got != tt.want {
				t.Errorf("HashTable.makeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_String(t *testing.T) {
	tab := &HashTable{
		Data: nil,
		Size: 10,
	}

	tests := []struct {
		name string
		tr   *HashTable
		want string
	}{
		{"correct", tab, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("HashTable.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
