package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	tests := []struct {
		name         string
		channels     []<-chan interface{}
		expectClosed bool
	}{
		{
			name: "All channels close",
			channels: []<-chan interface{}{
				sig(100 * time.Millisecond),
				sig(200 * time.Millisecond),
				sig(300 * time.Millisecond),
			},
			expectClosed: true,
		},
		{
			name: "Immediate close",
			channels: []<-chan interface{}{
				immediateCloseChan(),
			},
			expectClosed: true,
		},
		{
			name:         "No channels",
			channels:     []<-chan interface{}{},
			expectClosed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orChan := or(tt.channels...)
			select {
			case <-orChan:
				if !tt.expectClosed {
					t.Errorf("orChan closed unexpectedly")
				}
			case <-time.After(500 * time.Millisecond):
				if tt.expectClosed {
					t.Errorf("orChan did not close as expected")
				}
			}
		})
	}
}

func immediateCloseChan() <-chan interface{} {
	c := make(chan interface{})
	close(c)
	return c
}
