package stringset

// Define the Set type here.
type Set []string

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, v := range l {
		if !set.Has(v) {
			set = append(set, v)
		}
	}
	return Set(set)
}

func (s Set) String() string {
	result := "{"
	for i, v := range s {
		result += "\"" + v + "\""
		if i < len(s)-1 {
			result += ", "
		}
	}
	return result + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		*s = append(*s, elem)
	}
}

func Subset(s1, s2 Set) bool {
	switch {
	case s1.IsEmpty():
		return true
	case len(s1) > len(s2):
		return false
	default:
		for _, v := range s1 {
			if !s2.Has(v) {
				return false
			}
		}
		return true
	}
}

func Disjoint(s1, s2 Set) bool {
	switch {
	case s1.IsEmpty() || s2.IsEmpty():
		return true
	default:
		for _, v := range s1 {
			if s2.Has(v) {
				return false
			}
		}
		return true
	}
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		if !s2.Has(v) {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	intersection := New()

	for _, v := range s1 {
		if s2.Has(v) {
			intersection.Add(v)
		}
	}

	return intersection
}

func Difference(s1, s2 Set) Set {
	difference := New()

	for _, v := range s1 {
		if !s2.Has(v) {
			difference.Add(v)
		}
	}

	return difference
}

func Union(s1, s2 Set) Set {
	union := New()

	for _, v := range s1 {
		union.Add(v)
	}
	for _, v := range s2 {
		union.Add(v)
	}

	return union
}
