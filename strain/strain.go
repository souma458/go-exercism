package strain

func Keep[T any](s []T, p func(T) bool) []T {
	return filter(s, p, true)
}

func Discard[T any](s []T, p func(T) bool) []T {
	return filter(s, p, false)
}

func filter[T any](s []T, p func(T) bool, b bool) []T {
	result := make([]T, 0)
	for _, v := range s {
		if p(v) == b {
			result = append(result, v)
		}
	}
	return result
}
