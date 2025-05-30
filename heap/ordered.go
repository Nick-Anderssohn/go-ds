package heap

import "cmp"

type Ordered interface {
	LessThan(other Ordered) bool
}

type OrderedPrimitive cmp.Ordered

type box[T cmp.Ordered] struct {
	val T
}

func (b *box[T]) LessThan(other Ordered) bool {
	otherAsBox, _ := other.(*box[T])

	return b.val < otherAsBox.val
}
