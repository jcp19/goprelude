package option

import "github.com/jcp19/goprelude/utils"

type Option[T any] struct {
	value  T
	isSome bool
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func Some[T any](v T) Option[T] {
	return Option[T]{
		value:  v,
		isSome: true,
	}
}

func ValueOrElse[T any](o Option[T], fallback T) T {
	if o.isSome {
		return o.value
	}
	return fallback
}

func ValueOrDflt[T any](o Option[T]) T {
	if o.isSome {
		return o.value
	}
	var dflt T
	return dflt
}

func UnsafeGet[T any](o Option[T]) T {
	if o.isSome {
		return o.value
	}
	panic("Failed to get value from None.")
}

func IsSome[T any](o Option[T]) bool {
	return o.isSome
}

func IsNone[T any](o Option[T]) bool {
	return !o.isSome
}

func Map[T, S any](o Option[T], f func(T) S) Option[S] {
	if o.isSome {
		return Some(f(o.value))
	}
	return None[S]()
}

func Match[T, S any](o Option[T], caseNone func() S, caseSome func(T) S) S {
	if o.isSome {
		return caseSome(o.value)
	}
	return caseNone()
}

func Case[T any](o Option[T], caseNone func(), caseSome func(T)) {
	if o.isSome {
		caseSome(o.value)
	} else {
		caseNone()
	}
}

func Join[T any](o Option[Option[T]]) Option[T] {
	return Match(o, None, utils.Id)
}
