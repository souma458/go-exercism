package cryptosquare

import (
	"math"
	"unicode"
)

func Encode(pt string) string {
	normalized := normalize(pt)

	columns := int(math.Ceil(math.Sqrt(float64(len(normalized)))))
	rows := int(math.Ceil(float64(len(normalized)) / float64(columns)))

	result := ""
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			index := i + columns*j
			if index < len(normalized) {
				result += string(normalized[index])
			} else {
				result += " "
			}
		}
		if i < columns-1 {
			result += " "
		}
	}

	return result
}

func normalize(pt string) string {
	normalized := ""
	for _, char := range pt {
		if char >= 'a' && char <= 'z' || char >= '1' && char <= '9' {
			normalized += string(char)
		}
		if char >= 'A' && char <= 'Z' {
			normalized += string(unicode.ToLower(char))
		}
	}
	return normalized
}
