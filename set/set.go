package set

type Set[T any] map[any]bool

func Put[T any](s Set[T], val T) Set[T] {
	if s == nil {
		s = Set[T]{}
	}

	s[val] = true

	return s
}

func Exists[T any](s Set[T], val T) bool {
	if s == nil {
		return false
	}

	_, exists := s[val]
	return exists
}

func Remove[T any](s Set[T], val T) {
	delete(s, val)
}
