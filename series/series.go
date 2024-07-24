package series

func All(n int, s string) []string {
	result := make([]string, 0)
	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, UnsafeFirst(n, s[i:]))
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
