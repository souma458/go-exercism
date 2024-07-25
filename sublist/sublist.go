package sublist

func Sublist(l1, l2 []int) Relation {
	switch {
	case len(l1) == len(l2) && areEqual(l1, l2):
		return RelationEqual
	case len(l1) > len(l2) && isSuperlist(l1, l2):
		return RelationSuperlist
	case isSublist(l1, l2):
		return RelationSublist
	default:
		return RelationUnequal
	}
}

func areEqual(l1, l2 []int) bool {
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func isSuperlist(l1, l2 []int) bool {
	return isSublist(l2, l1)
}

func isSublist(l1, l2 []int) bool {
	for i := 0; i < len(l2)-len(l1)+1; i++ {
		if areEqual(l1, l2[i:i+len(l1)]) {
			return true
		}
	}
	return false
}
