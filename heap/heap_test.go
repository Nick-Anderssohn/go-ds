package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	h, _ := CreateHeap(HeapTypeMin, []int{69, 42, 420, -69, 1000})

	val := h.Pop()
	if val != -69 {
		t.Fatalf("expected %v got %v", -69, val)
	}
	// 42, 69, 420, 1000

	val = h.Pop()
	if val != 42 {
		t.Fatalf("expected %v got %v", 42, val)
	}
	// 69, 420, 1000

	h.Push(-420)
	// -420, 69, 420, 1000

	val = h.Pop()
	if val != -420 {
		t.Fatalf("expected %v got %v", -420, val)
	}
	// 69, 420, 1000

	h.Push(75)
	h.Push(420)
	h.Push(9999)
	// 69, 75, 420, 420, 1000, 9999

	val = h.Pop()
	if val != 69 {
		t.Fatalf("expected %v got %v", 69, val)
	}

	val = h.Pop()
	if val != 75 {
		t.Fatalf("expected %v got %v", 75, val)
	}

	val = h.Pop()
	if val != 420 {
		t.Fatalf("expected %v got %v", 420, val)
	}

	val = h.Pop()
	if val != 420 {
		t.Fatalf("expected %v got %v", 420, val)
	}

	val = h.Pop()
	if val != 1000 {
		t.Fatalf("expected %v got %v", 1000, val)
	}

	val = h.Pop()
	if val != 9999 {
		t.Fatalf("expected %v got %v", 9999, val)
	}
}
