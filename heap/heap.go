package heap

import (
	stdheap "container/heap"
	"fmt"
	"strconv"

	"github.com/Nick-Anderssohn/go-ds/heap/ordered"
)

type HeapType int8

const (
	HeapTypeMin HeapType = iota
	HeapTypeMax
)

// heap implements the standard library's heap.Interface
type heapWorker[T ordered.Type] struct {
	heapType HeapType
	values   []T
}

type heap[T ordered.Type] struct {
	worker heapWorker[T]
}

// Heap is a heap...duh.
type Heap[T ordered.Type] interface {
	Push(val T)
	Pop() T
	Len() int
}

// CreateHeap creates a heap. The only case where this will return
// an error is if you pass in an invalid value for heapType. So if
// you are directly passing in HeapTypeMin or HeapTypeMax, then it
// is safe to ignore the error.
func CreateHeap[T ordered.Type](heapType HeapType, initialValuesUnordered []T) (Heap[T], error) {
	if heapType > HeapTypeMax {
		return nil, fmt.Errorf("invalid value for heapType: %v", heapType)
	}

	h := &heap[T]{
		worker: heapWorker[T]{
			heapType: heapType,
			values:   initialValuesUnordered,
		},
	}

	stdheap.Init(&h.worker)

	return h, nil
}

func (h *heapWorker[T]) Push(val any) {
	h.values = append(h.values, val.(T))
}

func (h *heapWorker[T]) Pop() any {
	val := h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]
	return val
}

func (h *heapWorker[T]) Len() int {
	return len(h.values)
}

func (h *heapWorker[T]) Less(i, j int) bool {
	iVal := h.values[i]
	jVal := h.values[j]

	switch h.heapType {
	case HeapTypeMin:
		return iVal.LessThan(jVal)
	case HeapTypeMax:
		// The double check to handle the case where the values are equal
		return jVal.LessThan(iVal)
	}

	// should never happen since initialization validates heapType
	panic("invalid heap type: " + strconv.Itoa(int(h.heapType)))
}

func (h *heapWorker[T]) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *heap[T]) Push(val T) {
	stdheap.Push(&h.worker, val)
}

func (h *heap[T]) Pop() T {
	return stdheap.Pop(&h.worker).(T)
}

func (h *heap[T]) Len() int {
	return len(h.worker.values)
}
