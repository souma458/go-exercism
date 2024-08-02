package resistorcolortrio

import "strconv"

var colorValueMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Label(colors []string) string {
	ohms := strconv.Itoa(colorValueMap[colors[0]]) + strconv.Itoa(colorValueMap[colors[1]])
	zeros := colorValueMap[colors[2]]

	return convertUnit(ohms, zeros)
}

func convertUnit(ohms string, zeros int) string {
	unit := " ohms"
	for i := 0; i < zeros; i++ {
		ohms += "0"
	}

	num, _ := strconv.Atoi(ohms)
	switch {
	case num >= 1000000000:
		num /= 1000000000
		unit = " gigaohms"
	case num >= 1000000:
		num /= 1000000
		unit = " megaohms"
	case num >= 1000:
		num /= 1000
		unit = " kiloohms"
	}

	return strconv.Itoa(num) + unit
}
