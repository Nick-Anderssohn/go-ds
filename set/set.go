package set

type Set[T comparable] map[T]bool

func Put[T comparable](s Set[T], val T) Set[T] {
	if s == nil {
		s = Set[T]{}
	}

	s[val] = true

	return s
}

func Exists[T comparable](s Set[T], val T) bool {
	if s == nil {
		return false
	}

	_, exists := s[val]
	return exists
}

func Remove[T comparable](s Set[T], val T) {
	delete(s, val)
}

func FromSlice[T comparable](slice []T) (s Set[T]) {
	for _, val := range slice {
		Put(s, val)
	}

	return
}
