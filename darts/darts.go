package darts

import "math"

var innerCircleArea = circleArea(1)
var middleCircleArea = circleArea(5)
var outerCircleArea = circleArea(10)

func circleArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func Score(x, y float64) int {
	var circleArea = circleArea(distanceFromCenter(x, y))
	switch {
	case circleArea <= innerCircleArea:
		return 10
	case circleArea <= middleCircleArea:
		return 5
	case circleArea <= outerCircleArea:
		return 1
	}
	return 0
}

func distanceFromCenter(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
