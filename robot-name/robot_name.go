package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var nameDb = map[string]struct{}{}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		rName := generateName()
		if _, ok := nameDb[rName]; ok {
			return "", errors.New("name already in use")
		}
		nameDb[rName] = struct{}{}
		r.name = rName
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

// generateName generates a random string of 2 uppercase letters and 3 digits
func generateName() string {
	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate two random capital letters (A-Z)
	letters := make([]byte, 2)
	for i := range letters {
		letters[i] = 'A' + byte(rand.Intn(26))
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate three random digits (0-9)
	numbers := rand.Intn(1000) // 0 <= n < 1000
	numberString := fmt.Sprintf("%03d", numbers)

	// Combine letters and numbers
	return string(letters) + numberString
}
