package result

import (
	"fmt"
)

type Result[T any] struct {
	ok    bool
	value T
	err   error
}

func (r Result[T]) Value() T {
	if !r.ok {
		panic(r.err)
	}

	return r.value
}

func (r Result[T]) Err() error {
	if r.ok {
		return nil
	}

	return r.err
}

func (r Result[T]) Ok() bool {
	return r.ok
}

func Ok[T any](value T) Result[T] {
	return Result[T]{ok: true, value: value}
}

func Err[T any](err error) Result[T] {
	return Result[T]{ok: false, err: err}
}

func From[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}

	return Ok(value)
}

func Map[T, U any](f func(T) U, r Result[T]) Result[U] {
	if !r.ok {
		return Err[U](r.err)
	}

	return Ok(f(r.value))
}

func Filter[T any](predicate func(T) bool, r Result[T]) Result[T] {
	if !r.ok {
		return Err[T](r.err)
	}

	if !predicate(r.value) {
		return Err[T](fmt.Errorf("value does not match predicate"))
	}

	return Ok(r.value)
}
