package cipher

import (
	"strings"
	"unicode"
)

type shift struct {
	distance int
}
type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 24 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
	var result strings.Builder
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shifted := ((char-'a'+rune(c.distance))%26 + 26) % 26
			result.WriteRune('a' + shifted)
		} else if char >= 'A' && char <= 'Z' {
			shifted := ((char-'A'+rune(c.distance))%26 + 26) % 26
			result.WriteRune('a' + shifted)
		}
	}
	return result.String()
}

func (c shift) Decode(input string) string {
	c.distance = -c.distance
	return c.Encode(input)
}

func NewVigenere(key string) Cipher {
	aCount := 0
	for _, r := range key {
		if !unicode.IsLetter(r) || unicode.IsUpper(r) {
			return nil
		}
		if r == 'a' {
			aCount++
		}
	}
	if aCount == len(key) {
		return nil
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	var result strings.Builder
	keyLen := len(v.key)
	v.key = strings.ToLower(v.key)

	for i, j := 0, 0; i < len(input); i++ {
		char := rune(input[i])
		if unicode.IsLetter(char) {
			keyChar := rune(v.key[j%keyLen])
			shift := keyChar - 'a'
			encodedChar := ShiftChar(unicode.ToLower(char), shift)
			result.WriteRune(encodedChar)
			j++
		}
	}

	return result.String()
}

func (v vigenere) Decode(input string) string {
	var result strings.Builder
	keyLen := len(v.key)
	v.key = strings.ToLower(v.key)

	for i, j := 0, 0; i < len(input); i++ {
		char := rune(input[i])
		if unicode.IsLetter(char) {
			keyChar := rune(v.key[j%keyLen])
			shift := keyChar - 'a'
			decodedChar := ShiftChar(char, -shift)
			result.WriteRune(decodedChar)
			j++
		}
	}

	return result.String()
}

func ShiftChar(char, shift rune) rune {
	if char >= 'a' && char <= 'z' {
		return 'a' + ((char - 'a' + shift + 26) % 26)
	} else if char >= 'A' && char <= 'Z' {
		return 'a' + ((char - 'A' + shift + 26) % 26)
	}
	return char
}
