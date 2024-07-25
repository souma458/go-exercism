package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, i := range s {
		initial = fn(initial, i)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0)
	for _, i := range s {
		if fn(i) {
			result = append(result, i)
		}
	}
	return result
}

func (s IntList) Length() int {
	length := 0
	for range s {
		length++
	}
	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, 0)
	for _, i := range s {
		result = append(result, fn(i))
	}
	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, 0)
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = s.Append(list)
	}
	return s
}
