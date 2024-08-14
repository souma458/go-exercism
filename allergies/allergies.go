package allergies

var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	allergiesNames := []string{}
	for allergen := range allergens {
		if AllergicTo(allergies, allergen) {
			allergiesNames = append(allergiesNames, allergen)
		}
	}
	return allergiesNames
}

func AllergicTo(allergies uint, allergen string) bool {
	allergenValue := allergens[allergen]
	return allergies&allergenValue != 0
}
