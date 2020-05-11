package main

import (
	"sync"
	"testing"
)

func TestCounts_Add(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		data    map[string]int
	}
	type args struct {
		name string
		n    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counts{
				RWMutex: tt.fields.RWMutex,
				data:    tt.fields.data,
			}
		})
	}
}

func TestCounts_Register(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		data    map[string]int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counts{
				RWMutex: tt.fields.RWMutex,
				data:    tt.fields.data,
			}
		})
	}
}

func TestCounts_Val(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		data    map[string]int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counts{
				RWMutex: tt.fields.RWMutex,
				data:    tt.fields.data,
			}
			if got := c.Val(tt.args.name); got != tt.want {
				t.Errorf("Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getName(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getName(tt.args.idx); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}