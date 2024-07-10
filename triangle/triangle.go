package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if a == 0 || b == 0 || c == 0 || !areValidSidesOfTriangle(a, b, c) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

func areValidSidesOfTriangle(a, b, c float64) bool {
	return a+b >= c && b+c >= a && a+c >= b
}
