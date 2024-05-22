package errors

import (
	"fmt"
)

type Result[T any] struct {
	ok    bool
	value T
	err   error
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

func Map[T, U any](f func(T) U, result Result[T]) Result[U] {
	if !result.ok {
		return Err[U](result.err)
	}

	return Ok(f(result.value))
}

func Filter[T any](predicate func(T) bool, result Result[T]) Result[T] {
	if !result.ok {
		return Err[T](result.err)
	}

	if !predicate(result.value) {
		return Err[T](fmt.Errorf("division by zero"))
	}

	return Ok[T](result.value)
}

func divide(numerator, denominator int) Result[int] {
	if denominator == 0 {
		return Err[int](fmt.Errorf("division by zero"))
	}

	return Ok[int](numerator / denominator)
}

func main() {
	result := divide(10, 2)

	switch result {
	case Ok[int](result.value):
		fmt.Println("Resultado da divisão:", result.value)

	case Err[int](result.err):
		fmt.Println("Erro:", result.err)
	}
}
