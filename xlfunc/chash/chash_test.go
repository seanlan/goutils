package chash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCHash_Match(t *testing.T) {
	type fields struct {
		HashValues    HashValues
		HashNodes     map[uint32]*HashNode
		VirtualNumber int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *HashNode
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: fields{},
			args: args{
				name: "cc",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cHash := NewCHash()
			cHash.AddNode(HashNode{Name: "tab_1"})
			cHash.AddNode(HashNode{Name: "tab_2"})
			cHash.AddNode(HashNode{Name: "tab_3"})
			cHash.AddNode(HashNode{Name: "tab_4"})
			cHash.AddNode(HashNode{Name: "tab_5"})
			cHash.AddNode(HashNode{Name: "tab_6"})
			cHash.AddNode(HashNode{Name: "tab_7"})
			cHash.AddNode(HashNode{Name: "tab_8"})
			cHash.AddNode(HashNode{Name: "tab_9"})
			cHash.AddNode(HashNode{Name: "tab_10"})

			for i := 1; i < 100; i++ {
				fmt.Println(i, ":", cHash.Match(strconv.Itoa(i)).Name)
			}
		})
	}
}
