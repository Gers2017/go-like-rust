package option

type Option[T any] struct {
	Value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{Value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{Value: nil}
}

func (o *Option[T]) IsSome() bool {
	return o.Value != nil
}

func (o *Option[T]) IsNone() bool {
	return o.Value == nil
}

func (o *Option[T]) AsTuple() (*T, bool) {
	return o.Value, o.IsSome()
}

func (o *Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("Can't unwrap a nil value")
	}

	return *o.Value
}

func (o *Option[T]) UnwrapOr(fallback T) T {
	if o.IsNone() {
		return fallback
	}

	return *o.Value
}

func (o *Option[T]) UnwrapOrElse(f func() T) T {
	if o.IsNone() {
		return f()
	}

	return *o.Value
}

func Map[T any, U any](o *Option[T], f func(value T) U) Option[U] {
	if o.IsNone() {
		return None[U]()
	}

	value := f(o.Unwrap())
	return Some(value)
}

func MapOr[T any, U any](o *Option[T], fallback U, f func(value T) U) Option[U] {
	if o.IsNone() {
		return Some(fallback)
	}

	value := f(o.Unwrap())
	return Some(value)
}
