package utils

func Id[T any](t T) T {
	return t
}

func Append[T any](s *[]T, elems ...T) {
	*s = append(*s, elems...)
}

func Unreachable() {
	panic("A suposedly unreachable location was reached.")
}
