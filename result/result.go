package result

import "go-like-rust/option"

type Result[T any] struct {
	Value *T
	Error *error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{Value: &value, Error: nil}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Value: nil, Error: &err}
}

func (r *Result[T]) IsOk() bool {
	return r.Value != nil
}

func (r *Result[T]) IsErr() bool {
	return r.Error != nil
}

func (r *Result[T]) OkOption() option.Option[T] {
	return option.Option[T]{Value: r.Value}
}

func (r *Result[T]) ErrOption() option.Option[error] {
	return option.Option[error]{Value: r.Error}
}

func (r *Result[T]) Unwrap() T {
	if r.IsErr() {
		panic("Can't unwrap a nil value")
	}

	return *r.Value
}

func (r *Result[T]) UnwrapOr(fallback T) T {
	if r.IsErr() {
		return fallback
	}

	return *r.Value
}

func (r *Result[T]) UnwrapOrElse(f func() T) T {
	if r.IsErr() {
		return f()
	}

	return *r.Value
}

func (r *Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic("Can't unwrap err a nil value")
	}

	return *r.Error
}

func (r *Result[T]) And(other *Result[T]) Result[T] {
	var err *Result[T]
	if r.IsErr() {
		err = r
	} else if other.IsErr() {
		err = other
	}

	if err != nil {
		return *err
	}

	return *other
}

func MapResult[T any, U any](r *Result[T], f func(T) U) Result[U] {
	if r.IsErr() {
		return Err[U](*r.Error)
	}

	value := f(r.Unwrap())
	return Ok(value)
}

func MapResultOr[T any, U any](r *Result[T], fallback U, f func(T) U) Result[U] {
	if r.IsErr() {
		return Ok(fallback)
	}

	value := f(r.Unwrap())
	return Ok(value)
}
