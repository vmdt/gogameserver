package utils

func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range input {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, R any](input []T, mapper func(T) R) []R {
	var result []R
	for _, item := range input {
		result = append(result, mapper(item))
	}
	return result
}

func Reduce[T any, R any](input []T, initial R, reducer func(R, T) R) R {
	result := initial
	for _, item := range input {
		result = reducer(result, item)
	}
	return result
}
