package message

import (
	"fmt"
	"testing"
	// "time"
)

func TestMessage(t *testing.T) {
	m := NewMessage("Hello")
	if m == nil {
		t.Errorf("Failed to create a new message")
	}

	a := make(map[int]string)
	a[2] = "Two"
	fmt.Println(a)
}
