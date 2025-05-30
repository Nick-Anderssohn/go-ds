package set

type Set[T any] map[any]bool

func (s Set[T]) Put(val T) {
	s[val] = true
}

func (s Set[T]) Exists(val T) bool {
	_, exists := s[val]
	return exists
}

func (s Set[T]) Remove(val T) {
	delete(s, val)
}
