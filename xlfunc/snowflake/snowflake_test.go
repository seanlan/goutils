package snowflake

import (
	"sync"
	"testing"
	"time"
)

func TestNode_Generate(t *testing.T) {
	type fields struct {
		mu        sync.Mutex
		epoch     time.Time
		time      int64
		node      int64
		step      int64
		nodeMax   int64
		nodeMask  int64
		stepMask  int64
		timeShift uint8
		nodeShift uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   ID
	}{
		// TODO: Add test cases.
		{},
		{},
		{},
		{},
	}
	n, _ := NewNode(1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Generate(): %s", n.Generate().String())
		})
	}
}
