package main

func Map[T any, U any](input []T, fn func(T) U) []U {
	results := make([]U, len(input))

	for i, v := range input {
		results[i] = fn(v)
	}

	return results
}

func Filter[T any](input []T, fn func(T) bool) []T {
	results := []T{}

	for _, v := range input {
		if fn(v) {
			results = append(results, v)
		}
	}

	return results
}

func Reduce[T any, U any](input []T, initial U, fn func(U, T) U) U {
	result := initial

	for _, v := range input {
		result = fn(result, v)
	}

	return result
}

func Find[T any](input []T, fn func(T) bool) *T {
	for _, v := range input {
		if fn(v) {
			return &v
		}
	}

	return nil
}

func Some[T any](input []T, fn func(T) bool) bool {
	for _, v := range input {
		if fn(v) {
			return true
		}
	}

	return false
}

func Every[T any](input []T, fn func(T) bool) bool {
	for _, v := range input {
		if !fn(v) {
			return false
		}
	}

	return true
}
